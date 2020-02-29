package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"errors"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func Verify(hash []byte, public_key_hex database.PublicKey, signature_hex string) (result bool, err error) {
	// decode public key from hex
	public_key_bytes, err := hex.DecodeString(string(public_key_hex))
	if err != nil {
		return false, err
	}

	// x509 parse public key
	public_key, err := x509.ParsePKIXPublicKey(public_key_bytes)
	if err != nil {
		return false, err
	}

	// check that parse key is ecdsa.PublicKey
	switch public_key := public_key.(type) {
	case *ecdsa.PublicKey:
		signature_bytes, err := hex.DecodeString(string(signature_hex))
		if err != nil {
			return false, err
		}
		// unmarhsal signature structure to extract signature from
		signature := new(Signature)
		_, err = asn1.Unmarshal(signature_bytes, signature)
		if err != nil {
			return false, err
		}

		// verify signature
		return ecdsa.Verify(public_key, hash, signature.R, signature.S), nil

	default:
		// only ECDSA public keys are supported
		return false, errors.New("only ECDSA public keys supported")
	}
}

// Verify verifies a previously generated signature for byte array hash using hex-encoded public key
func VerifyRaw(hash []byte, public_key_hex database.PublicKey, signature_hex string) (result bool, err error) {
	// decode public key from hex
	public_key_bytes, err := hex.DecodeString(string(public_key_hex))
	if err != nil {
		logger.Println("verify_message.go", "VerifyRaw()", err.Error())
		return false, err
	}

	point := new(ECCPoint)
	_, err = asn1.Unmarshal(public_key_bytes, point)
	if err != nil {
		logger.Println("verify_message.go", "VerifyRaw()", err.Error())
		return false, err
	}

	public_key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     StringToBigInt(point.X),
		Y:     StringToBigInt(point.Y),
	}

	signature_bytes, err := hex.DecodeString(string(signature_hex))
	if err != nil {
		logger.Println("verify_message.go", "VerifyRaw()", err.Error())
		return false, err
	}

	// unmarhsal signature structure to extract signature from
	signature := new(Signature)
	_, err = asn1.Unmarshal(signature_bytes, signature)
	if err != nil {
		logger.Println("verify_message.go", "VerifyRaw()", err.Error())
		return false, err
	}

	// verify signature
	return ecdsa.Verify(public_key, hash, signature.R, signature.S), nil

}

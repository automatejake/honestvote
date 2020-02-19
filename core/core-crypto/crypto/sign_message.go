package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

// signature is a structure for storing signature obtained from ecdsa.Sign
type Signature struct {
	R,
	S *big.Int
}

// Sign calculates a signature for a byte array hash using hex-encoded private key
// It is supposed that a hash is calculated for an original message to sign
// Signature is a hex-encoded JSON
func Sign(hash []byte, private_key_hex string) (signature_hex string, err error) {
	// decode private key from hex
	private_key_bytes, err := hex.DecodeString(private_key_hex)
	if err != nil {
		return "", err
	}

	// x509 parse private key
	private_key, err := x509.ParseECPrivateKey(private_key_bytes)
	if err != nil {
		return "", err
	}

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, private_key, hash)
	if err != nil {
		return "", err
	}

	// prepare a signature structure to marshal into json
	signature := &Signature{
		R: r,
		S: s,
	}

	// marshal to json
	signature_asn1, err := asn1.Marshal(*signature)
	if err != nil {
		return "", err
	}

	// encode to hex
	signature_hex = hex.EncodeToString(signature_asn1)
	return signature_hex, nil
}

type ECCPoint struct {
	X string
	Y string
}

func StringToBigInt(s string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {

	}
	return n
}

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
		return false, err
	}

	point := new(ECCPoint)
	_, err = asn1.Unmarshal(public_key_bytes, point)
	if err != nil {
		return false, err
	}

	public_key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     StringToBigInt(point.X),
		Y:     StringToBigInt(point.Y),
	}

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

}

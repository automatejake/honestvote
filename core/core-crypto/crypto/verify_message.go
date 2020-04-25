package crypto

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/hex"

	"github.com/jneubaum/honestvote/tests/logger"
)

// Verify verifies a previously generated signature for byte array hash using hex-encoded public key
func Verify(hash []byte, public_key_hex string, signature_hex string) (result bool, err error) {
	// decode public key from hex
	public_key_bytes, err := hex.DecodeString(string(public_key_hex))
	if err != nil {
		logger.Println("verify_message.go", "Verify()", err)
		return false, err
	}

	public_key := DecompressPoint(public_key_bytes)

	signature_bytes, err := hex.DecodeString(string(signature_hex))
	if err != nil {
		logger.Println("verify_message.go", "Verify()", err)
		return false, err
	}

	// unmarhsal signature structure to extract signature from
	signature := new(Signature)
	_, err = asn1.Unmarshal(signature_bytes, signature)
	if err != nil {
		logger.Println("verify_message.go", "Verify()", err)
		return false, err
	}

	// verify signature
	return ecdsa.Verify(public_key, hash, signature.R, signature.S), nil

}

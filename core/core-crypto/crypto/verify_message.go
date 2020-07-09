package crypto

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/hex"

	"github.com/jneubaum/honestvote/tests/logger"
)

var filename string = "verify_message"

// Verify verifies a previously generated signature for byte array hash using hex-encoded public key
func Verify(hash []byte, public_key_hex string, signature_hex string) (result bool, err error) {
	// decode public key from hex
	public_key_bytes, err := hex.DecodeString(string(public_key_hex))
	if err != nil || public_key_bytes == nil {
		logger.Println(filename, "Verify()", err)
		return false, err
	}
	logger.Println(filename, "Verify()", "Decoded string")

	public_key, err := DecompressPoint(public_key_bytes)
	if err != nil {
		logger.Println(filename, "Verify()", err)
	}

	logger.Println(filename, "Verify()", "Decompressed point to public key: ", public_key)

	signature_bytes, err := hex.DecodeString(string(signature_hex))
	if err != nil {
		logger.Println("verify_message.go", "Verify()", err)
		return false, err
	}
	logger.Println(filename, "Verify()", "Decoded signature from hex string: ", signature_bytes)

	// unmarhsal signature structure to extract signature from
	signature := new(Signature)
	_, err = asn1.Unmarshal(signature_bytes, signature)
	if err != nil {
		logger.Println(filename, "Verify()", err)
		return false, err
	}
	logger.Println(filename, "Verify()", "DER encoded signature: ", signature)

	// verify signature
	return ecdsa.Verify(public_key, hash, signature.R, signature.S), nil

}

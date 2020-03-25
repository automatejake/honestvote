package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"math/big"
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

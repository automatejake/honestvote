package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/asn1"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/jneubaum/honestvote/tests/logger"
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
	priv, shouldWork := new(big.Int).SetString(private_key_hex, 16)
	if !shouldWork {
		logger.Println("sign_message", "Sign", "Signature invalid")
		return "", errors.New("Private key not valid.")
	}

	private_key := new(ecdsa.PrivateKey)
	private_key.PublicKey.Curve = p256
	private_key.D = priv

	// private_key.PublicKey.X, private_key.PublicKey.Y = p256.ScalarBaseMult(priv.Bytes())

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, private_key, hash)
	if err != nil {
		logger.Println("sign_message.go", "Sign()", err)
		return "", err
	}

	// prepare a signature structure to marshal into asn1 der encoding
	signature := &Signature{
		R: r,
		S: s,
	}

	// marshal to asn1 der encoding
	signature_asn1, err := asn1.Marshal(*signature)
	if err != nil {
		logger.Println("sign_message.go", "Sign()", err)
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
		logger.Println("sign_message.go", "StringToBigInt()", "Set string not ok")
	}
	return n
}

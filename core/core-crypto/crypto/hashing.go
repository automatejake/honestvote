package crypto

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CalculateHash(encodedMessage []byte) string {
	hash32 := sha256.Sum256(encodedMessage)

	hash := make([]byte, 0, 32)
	copy(hash32[:], hash[:])

	return base64.URLEncoding.EncodeToString(hash)
}

func SignBlock(header []byte, privKey string) (string, error) {
	return Sign(header, privKey)
}

func SignTransaction(hash string, privKey string) (string, error) {
	signature, err := Sign([]byte(hash), privKey)
	if err != nil {
		logger.Println("hashing.go", "SignTransaction()", err)
		return "", err
	}
	return signature, nil
}

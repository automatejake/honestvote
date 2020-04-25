package crypto

import (
	"crypto/sha256"
	"fmt"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CalculateHash(encodedMessage []byte) string {
	hash32 := sha256.Sum256(encodedMessage)

	fmt.Println("this gonna be a fucking pain: ", hash32)

	return string(hash32[:])
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

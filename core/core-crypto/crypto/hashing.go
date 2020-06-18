package crypto

import (
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CalculateHash(encodedMessage []byte) []byte {
	hash32 := sha256.Sum256(encodedMessage)
	return hash32[:]
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

func HashTransaction(transaction interface{}) string {
	encoded, err := asn1.Marshal(transaction)
	if err != nil {
		logger.Println("hashing.go", "HashTransaction()", err)
	}

	calcHash := CalculateHash(encoded)
	hexTransaction := hex.EncodeToString(calcHash)

	return hexTransaction
}

package crypto

import (
	"crypto/sha256"
	"encoding/base64"
)

func CalculateHash([]byte) string {
	hash := sha256.New()
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func SignBlock(header []byte, privKey string) (string, error) {
	return Sign(header, privKey)
}

func SignTransaction(hash string, privKey string) (string, error) {
	signature, err := Sign([]byte(hash), privKey)
	if err != nil {
		return "", err
	}
	return signature, nil
}

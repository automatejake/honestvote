package crypto

import (
	"crypto/sha256"
	"encoding/base64"
)

func CalculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func SignBlock(header string, privKey string) (string, error) {
	return Sign([]byte(header), privKey)
}

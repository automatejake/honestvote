//generates public and private keys
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"

	"github.com/jneubaum/honestvote/tests/logger"
)

var prefix string = "33"

// GenerateKeyPair generates a private/public key pair,
// keys are returned as hex-encoded strings
func GenerateKeyPair() (private_key_hex, public_key_hex string) {
	// generate keys
	private_key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) //P256 returns a Curve
	if err != nil {
		logger.Println("key_generation.go", "GenerateKeyPair()", err.Error())
	}

	// marshal private key
	private_key_bytes, err := x509.MarshalECPrivateKey(private_key)
	if err != nil {
		logger.Println("key_generation.go", "GenerateKeyPair()", err.Error())
	}

	//get the bitLength for priv key then check if it's 256:
	bitLen := private_key.Curve.Params().BitSize //wont work with pub key : different data type
	LengthIsValid(bitLen)                        //send bitLen in order to check if len is correct for priv key, returns true

	// marshal public key
	public_key_bytes, err := x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	if err != nil {
		logger.Println("key_generation.go", "GenerateKeyPair()", err.Error())
	}

	// hex encode and return result
	private_key_hex = hex.EncodeToString(private_key_bytes)
	public_key_hex = hex.EncodeToString(public_key_bytes)

	return private_key_hex, public_key_hex
}

func LengthIsValid(x int) bool { // checks if the len of priv key is 256 as it should be

	//fmt.Println(x)
	if x == 256 {
		return true
	} else {
		return false
	}

}

//generates public and private keys
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"math/big"

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

func DecompressPoint(compressed_bytes []byte) *ecdsa.PublicKey {
	// Split the sign byte from the rest
	sign_byte := uint(compressed_bytes[0])
	x_bytes := compressed_bytes[1:]
	// Convert to big Int.
	x := new(big.Int).SetBytes(x_bytes)
	// We use 3 a couple of times
	three := big.NewInt(3)
	// and we need the curve params for P256
	c := elliptic.P256().Params()
	// The equation is y^2 = x^3 - 3x + b
	// First, x^3, mod P
	x_cubed := new(big.Int).Exp(x, three, c.P)
	// Next, 3x, mod P
	three_X := new(big.Int).Mul(x, three)
	three_X.Mod(three_X, c.P)
	// x^3 - 3x ...
	y_squared := new(big.Int).Sub(x_cubed, three_X)
	// ... + b mod P
	y_squared.Add(y_squared, c.B)
	y_squared.Mod(y_squared, c.P)
	// Now we need to find the square root mod P.
	// This is where Go's big int library redeems itself.
	y := new(big.Int).ModSqrt(y_squared, c.P)
	if y == nil {
		// If this happens then you're dealing with an invalid point.
		// Panic, return an error, whatever you want here.
	}
	// Finally, check if you have the correct root by comparing
	// the low bit with the low bit of the sign byte. If it’s not
	// the same you want -y mod P instead of y.
	if y.Bit(0) != sign_byte&1 {
		y.Neg(y)
		y.Mod(y, c.P)
	}
	// Now your y coordinate is in y, for all your ScalarMult needs.
	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x, Y: y,
	}
	return publicKey
}

func LengthIsValid(x int) bool { // checks if the len of priv key is 256 as it should be

	if x == 256 {
		return true
	} else {
		return false
	}

}

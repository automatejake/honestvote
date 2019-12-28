//generates public and private keys
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/jneubaum/honestvote/tests/logger"
)

var prefix string = "33"

func GenerateKey() (*ecdsa.PrivateKey, ecdsa.PublicKey) {

	//priv key gen start
	pubkeyCurve := elliptic.P256() //P256 returns a Curve
	privKey := new(ecdsa.PrivateKey)
	privKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	if err != nil {
		logger.Println("key_generation.go", "KeyGen()", err.Error())
	}
	//end priv key gen

	//pub key gen start
	var pubkey ecdsa.PublicKey
	pubkey = privKey.PublicKey
	//end pub key gen

	//get the bitLength for priv key then check if it's 256:
	var bitLen int
	bitLen = privKey.Curve.Params().BitSize //wont work with pub key : different data type
	LengthIsValid(bitLen)                   //send bitLen in order to check if len is correct for priv key, returns true

	//CREATING ADDRESS FROM PUBKEY
	var add1 string
	add1 = prefix + pubkey.Y.String() //String address
	//TURN ADDRESS INTO A BIGINT:
	//https://stackoverflow.com/questions/46783352/string-to-big-int-in-go
	//https://golang.org/pkg/math/big/#example_Int_SetString   - see SetString header
	n := new(big.Int)
	n, ok := n.SetString(pubkey.Y.String(), 10)
	if !ok {
		fmt.Println("SetString: error")
		//return
		os.Exit(1)
	}
	fmt.Println(add1) //temp: here to remover error
	//n is the second portion of the pub key with a prefix

	return privKey, pubkey

}
func LengthIsValid(x int) bool { // checks if the len of priv key is 256 as it should be

	//fmt.Println(x)
	if x == 256 {
		return true
	} else {
		return false
	}

}

func PrefixIsValid(z *big.Int) bool { //takes in address as a bigInt, returns true or false depending on if it has the correct prefix
	var t string
	t = z.String()
	return strings.HasPrefix(t, prefix)
}

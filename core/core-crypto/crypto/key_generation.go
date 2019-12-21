//generates public and private keys
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"math/big"
	"os"
)

func genKey() (*ecdsa.PrivateKey, ecdsa.PublicKey) {

	//priv key gen start

	pubkeyCurve := elliptic.P256() //P256 returns a Curve

	privKey := new(ecdsa.PrivateKey)

	privKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

	//At this point: privatekey is complete and err is null

	if err != nil { //exits if err contains a value
		//return //"cannot use nil as type ecdsa.PublicKey in return argumentgo"
		os.Exit(1)
	}

	//end priv key gen

	//pub key gen start
	var pubkey ecdsa.PublicKey

	pubkey = privKey.PublicKey

	return privKey, pubkey

	//end pub key gen

}
func isValid(*ecdsa.PrivateKey, ecdsa.PublicKey) {

	var hash1 hash.Hash //hash value: hash1=<nil>

	hash1 = md5.New() // the md5 fn :producing a (128-bit) hash value

	fmt.Println(hash1) //temp
	r := big.NewInt(0)
	fmt.Println(r) //temp
}

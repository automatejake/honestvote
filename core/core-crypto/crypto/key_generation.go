//generates public and private keys
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
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
	//end pub key gen

	//get the bitLength for priv key then check if it's 256:
	var bitLen int
	bitLen = privKey.Curve.Params().BitSize //wont work with pub key : different data type
	//fmt.Println(bitLen)
	lenIsValid(bitLen) //send bitLen in order to check if len is correct for priv key, returns true
	//fmt.Println(lenIsValid(bitLen)) //returns bool val: true

	//testing sign1
	var msg string
	msg = "message"
	//fmt.Println(msg)
	sign1(msg, privKey, pubkey)

	return privKey, pubkey

}
func lenIsValid(x int) bool { // checks if the len of priv key is 256 as it should be

	//fmt.Println(x)
	if x == 256 {
		return true
	} else {
		return false
	}

}
func sign1(msg string, q *ecdsa.PrivateKey, w ecdsa.PublicKey) bool {

	var hash1 hash.Hash //hash value: hash1=<nil>

	hash1 = md5.New() // the md5 fn :producing a (128-bit) hash value

	fmt.Println(hash1) //temp

	s := new(big.Int) //values used in ECDSA
	r := new(big.Int)

	//fmt.Println(r) //temp
	//fmt.Println(s) //temp

	io.WriteString(hash1, msg) //the message is now hashed
	sum1 := hash1.Sum(nil)[:]

	r, s, nilVal := ecdsa.Sign(rand.Reader, q, sum1)
	//fmt.Println(nilVal)
	if nilVal != nil {
		fmt.Println(nilVal)
		os.Exit(1)
	}
	verify := ecdsa.Verify(&w, sum1, r, s)
	return verify

}

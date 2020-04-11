package main

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

var plaintext string = "this should work, please work.  I wonder why this does not work.  Kernel Sanders is king of all chicken and I wonder how chicken would taste if it were not for the great KFC Lord"

func main() {
	private_key, public_key := crypto.GenerateKeyPair()
	hash := []byte(plaintext)
	signature, err := crypto.Sign(hash, private_key)
	fmt.Println(signature)
	if err != nil {
		fmt.Println("Signature bad", err)
		return
	}

	itShouldWork, err := crypto.Verify(hash, public_key, signature)
	if !itShouldWork {
		fmt.Println("not verified sir", err)
		return
	}
}

package consensus

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsBlockValid(prevBlock database.Block, block database.Block) bool {
	if prevBlock.Hash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	}
	//add code to validate other stuff

	//add code to verify each transaction

	return true
}

func VerifySignature(transaction interface{}) bool {
	var header string

	switch t := transaction.(type) {
	case *database.Vote:
		header = string(t.Sender) + t.Election + t.Signature +
			t.Type

		for k, v := range t.Receiver {
			header = header + k + v
		}

		correct, err := crypto.Verify([]byte(header), t.Sender, t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		}
	case *database.Election:
		header = t.ElectionName + t.EmailDomain + t.Institution +
			string(t.Sender) + t.Signature + t.Type + t.Start +
			t.End

		correct, err := crypto.Verify([]byte(header), t.Sender, t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		}
	case *database.Registration:
		header = t.Election + t.Receiver + string(t.Sender) + t.Signature +
			t.Type

		correct, err := crypto.Verify([]byte(header), t.Sender, t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		}
	}

	return false
}

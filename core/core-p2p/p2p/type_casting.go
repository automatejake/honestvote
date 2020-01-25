package p2p

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func VerifySignature(transaction interface{}) bool {
	var header string

	switch t := transaction.(type) {
	case *database.Vote:
		header = string(t.Sender)

		for k, v := range t.Receiver {
			header = header + k + v
		}

		correct, err := crypto.Verify([]byte(header), string(t.Sender), t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		}
	case *database.Election:
		header = t.ElectionName + t.Start + t.End

		correct, err := crypto.Verify([]byte(header), string(t.Sender), t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		}
	}

	return false
}

func TransactionType(transaction interface{}) string {

	switch t := transaction.(type) {
	case database.Vote:
		return "Vote"
	case database.Registration:
		return "Registration"
	case database.Election:
		return "Election"
	default:
		fmt.Println(t)
	}

	return ""
}

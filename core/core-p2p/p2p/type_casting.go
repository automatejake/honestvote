package p2p

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func CreateSignature(transaction interface{}, privKey string) string {
	var header string

	if t, ok := transaction.(database.Vote); ok {
		header = string(t.Sender)

		for k, v := range t.Receiver {
			header = header + k + v
		}

		sig, err := crypto.Sign([]byte(header), privKey)

		if err == nil {
			return sig
		}
	} else if t, ok := transaction.(database.Election); ok {
		header = t.ElectionName + t.Start + t.End

		sig, err := crypto.Sign([]byte(header), privKey)

		if err == nil {
			return sig
		}
	}

	return "There was an error"
}

func VerifySignature(transaction interface{}) bool {
	var header string

	if t, ok := transaction.(*database.Vote); ok {
		header = string(t.Sender)

		for k, v := range t.Receiver {
			header = header + k + v
		}

		correct, err := crypto.Verify([]byte(header), string(t.Sender), t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		} else {
			fmt.Println("Error!")
		}
	} else if t, ok := transaction.(*database.Election); ok {
		header = t.ElectionName + t.Start + t.End

		correct, err := crypto.Verify([]byte(header), string(t.Sender), t.Signature)

		if err == nil {
			fmt.Println("Signature is ", correct)
			return correct
		} else {
			fmt.Println("Error!")
		}
	}

	return false
}

func TransactionType(transaction interface{}) string {

	if _, ok := transaction.(database.Vote); ok {
		return "Vote"
	} else if _, ok := transaction.(database.Registration); ok {
		return "Registration"
	} else if _, ok := transaction.(database.Election); ok {
		return "Election"
	}

	return ""
}

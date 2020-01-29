package p2p

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func CreateSignature(transaction interface{}, privKey string) string {
	var header string

	switch t := transaction.(type) {
	case database.Vote:
		header = string(t.Sender) + t.Election + t.Signature +
			t.Type

		for k, v := range t.Receiver {
			header = header + k + v
		}

		sig, err := crypto.Sign([]byte(header), privKey)

		if err == nil {
			return sig
		}
	case database.Election:
		header = t.ElectionName + t.EmailDomain + t.Institution +
			string(t.Sender) + t.Signature + t.Type + t.Start +
			t.End

		sig, err := crypto.Sign([]byte(header), privKey)

		if err == nil {
			return sig
		}
	case database.Registration:
		header = t.Election + string(t.Receiver) + string(t.Sender) + t.Signature +
			t.Type

		sig, err := crypto.Sign([]byte(header), privKey)

		if err == nil {
			return sig
		}
	}

	return "There was an error"
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

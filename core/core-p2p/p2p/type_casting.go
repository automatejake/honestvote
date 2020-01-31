package p2p

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

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

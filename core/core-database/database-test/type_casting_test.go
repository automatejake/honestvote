package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
	//"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestTransactionType(t *testing.T) {
	var rType string

	rType = database.TransactionType(database.Vote{})
	if rType == "Vote" {
		t.Log("Vote was given the correct type: ", rType)
	} else {
		t.Log("Vote was NOT given the correct type: ", rType)
	}

	rType = database.TransactionType(database.Registration{})
	if rType == "Registration" {
		t.Log("Registration was given the correct type: ", rType)
	} else {
		t.Log("Registration was NOT given the correct type: ", rType)
	}

	rType = database.TransactionType(database.Election{})
	if rType == "Election" {
		t.Log("Election was given the correct type: ", rType)
	} else {
		t.Log("Election was NOT given the correct type: ", rType)
	}

}

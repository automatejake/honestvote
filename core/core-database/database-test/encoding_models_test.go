package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestBlockEncode(t *testing.T) {
	block := new(database.Block)

	encoded, err := block.Encode()

	if err != nil {
		t.Error("There shouldn't be an error when encoding a block: ", err)
	}

	t.Log("The block was encoded successfully: ", encoded)
}

func TestElectionEncode(t *testing.T) {
	node := new(database.Node)

	encoded, err := node.Encode()

	if err != nil {
		t.Error("There shouldn't be an error when encoding a block: ", err)
	}

	t.Log("The block was encoded successfully: ", encoded)
}

func TestRegistrationEncode(t *testing.T) {
	registration := new(database.Registration)

	encoded, err := registration.Encode()

	if err != nil {
		t.Error("There shouldn't be an error when encoding a block: ", err)
	}

	t.Log("The block was encoded successfully: ", encoded)
}

func TestVoteEncode(t *testing.T) {
	vote := new(database.Vote)

	encoded, err := vote.Encode()

	if err != nil {
		t.Error("There shouldn't be an error when encoding a block: ", err)
	}

	t.Log("The block was encoded successfully: ", encoded)
}

func TestNodeEncode(t *testing.T) {
	node := new(database.Node)

	encoded, err := node.Encode()

	if err != nil {
		t.Error("There shouldn't be an error when encoding a block: ", err)
	}

	t.Log("The block was encoded successfully: ", encoded)
}

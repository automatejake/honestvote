package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestCorrespondingRegistration(t *testing.T) { //see test for AcceptConnectionMessage(p2p)
	database.MongoDB = database.MongoConnect("localhost")

	registration := database.CorrespondingRegistration(database.Vote{})

	t.Log("Here's the corresponding registration to the vote given: ", registration)
}

// func TestContainsRegistration(t *testing.T) {
// 	database.MongoDB = database.MongoConnect("localhost")
// }

// func TestContainsVote(t *testing.T) {

// }

func TestMarkDishonestNode(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	err := database.MarkDishonestNode(database.Node{})

	if err != nil {
		t.Error("There shouldn't be an error when marking a dishonest node. Error: ", err)
		return
	}

	t.Log("There was no error when marking the dishonest node.")
}

// func TestRetTrue(t *testing.T) {

// }

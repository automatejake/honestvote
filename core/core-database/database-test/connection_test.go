package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestMongoConnect(t *testing.T) {
	connection := database.MongoConnect("localhost")

	if connection == nil {
		t.Error("Connection should not be nil.")
		return
	}

	// t.Log("Connection is active: ", connection)
}

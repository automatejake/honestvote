package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestGetElections(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	elections, err := database.GetElections()

	if err != nil {
		t.Error("There shouldn't be an error when grabbing elections from database. Error: ", err)
		return
	}
	if elections == nil {
		t.Log("No elections were returned when grabbing from the databasae, this could be an error.")
	}

	t.Log("There was no error when grabbing elections from the database: ", elections)
}
func TestGetElection(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	election, err := database.GetElection(" ") //TODO: This shouldn't be an empty string

	if err != nil {
		t.Error("There shouldn't be an error when grabbing elections from database. Error: ", err)
		return
	}

	t.Log("There was no error when grabbing the election from the database: ", election)
}
func TestGetVotes(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	votes, err := database.GetVotes(" ") //TODO: This shouldn't be an empty string

	if err != nil {
		t.Error("There shouldn't be an error when grabbing votes from database. Error: ", err)
		return
	}
	if votes == nil {
		t.Log("No votes were returned when grabbing from the databasae, this could be an error.")
	}

	t.Log("There was no error when grabbing the votes from the database: ", votes)
}
func TestGetPermissions(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	permissions, err := database.GetPermissions(" ") //TODO: This shouldn't be an empty string

	if err != nil {
		t.Error("There shouldn't be an error when grabbing permissions from database. Error: ", err)
		return
	}
	if permissions == nil {
		t.Log("No permissions were returned when grabbing from the databasae, this could be an error.")
	}

	t.Log("There was no error when grabbing the permissions from the database: ", permissions)
}

func TestGetEndPoint(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	endpoint, err := database.GetEndpoint()

	if err != nil {
		t.Error("There shouldn't be an error when grabbing the endpoint from database. Error: ", err)
		return
	}
	if endpoint == "" {
		t.Log("No endpoint was returned when grabbing from the databasae, this could be an error.")
	}

	t.Log("There was no error when grabbing the endpoint from the database: ", endpoint)
}

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

	election, err := database.GetElection("") //TODO: This shouldn't be an empty string

	if err != nil {
		t.Error("There shouldn't be an error when grabbing elections from database. Error: ", err)
		return
	}

	t.Log("There was no error when grabbing the election from the database: ", election)
}
func TestGetVotes(t *testing.T) {
	//	database.GetVotes()

}
func TestGetPermissions(t *testing.T) {
	//	database.GetPermissions()

}

func GetEndPoint(t *testing.T) {
	//	database.GetEndpoint()

}

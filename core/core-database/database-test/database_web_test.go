package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestGetElections(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	elections, err := database.GetElections()

	if err != nil {
		t.Error("There was an error when grabbing elections from database. Error: ", err)
		return
	}
	if elections == nil {
		t.Log("No elections were returned when grabbing from the databasae, this could be an error.")
	}

	t.Log("There was no error when grabbing elections from the database: ", elections)
}
func TestGetElection(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")
	database.CollectionPrefix = "a_"

	election, err := database.GetElection(
		"30440220391cca210fde6ed8b1b7315da6b9e12b1c3b314d76b9c5ebc2a4ec5a1854397302200200e0448984b94a9c95cc1112449f6e4f611ec2092ae30376da674c8762f5f9") //TODO: This shouldn't be an empty string

	if err != nil {
		t.Error("There shouldn't be an error when grabbing elections from database. Error: ", err)
		return
	}

	t.Log("There was no error when grabbing the election from the database: ", election)
}
func TestGetVotes(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")
	database.CollectionPrefix = "a_"

	votes, err := database.GetVotes(
		"30440220391cca210fde6ed8b1b7315da6b9e12b1c3b314d76b9c5ebc2a4ec5a1854397302200200e0448984b94a9c95cc1112449f6e4f611ec2092ae30376da674c8762f5f9") //TODO: This shouldn't be an empty string

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
	database.CollectionPrefix = "a_"

	permissions, err := database.GetPermissions(
		"30819e134d3130333232323837363238303134313435323230373237353435313335353934353139313231313731333136363030313434313830323734343038363138373438363734353233343433353332134d3231363537383431343132303231373936373339323136333034303030313930343935363530383737313139383235303432353931353234333632323035393333363039303435343534373233") //TODO: This shouldn't be an empty string

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
	database.CollectionPrefix = "a_"

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

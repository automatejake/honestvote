package database

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func GetElections() ([]Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "elections")
	var elections []Election
	var election Election

	query := bson.M{}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err)
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&election)
		if err != nil {
			logger.Println("database_web", "GetElection", err)
		}

		elections = append(elections, election)
	}

	result.Close(context.TODO())
	return elections, nil
}

func GetElection(election_signature string) (Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "elections")
	var election Election

	query := bson.M{"signature": election_signature}

	result := collection.FindOne(context.TODO(), query)

	err := result.Decode(&election)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
		return election, err
	}

	return election, nil
}

func GetVotes(electionId string) ([]Vote, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "votes")

	var votes []Vote
	var vote Vote

	query := bson.M{"electionId": electionId}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&vote)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		votes = append(votes, vote)
	}

	result.Close(context.TODO())
	return votes, nil
}

func GetPermissions(public_key string) ([]string, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "registrations")
	var registration Registration
	var elections []string

	query := bson.M{"receiver": public_key}
	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&registration)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		election := registration.Election
		elections = append(elections, election)
	}

	result.Close(context.TODO())
	return elections, nil
}

func GetEndpoint() (string, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "node_list")

	var nodes []Node
	var node Node

	query := bson.M{}
	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetEndpoint", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&node)
		if err != nil {
			logger.Println("database_web", "GetEndpoint", err.Error())
		}
		nodes = append(nodes, node)
	}

	randNode := rand.Intn(len(nodes))
	port := strconv.Itoa(nodes[randNode].Port)
	endpoint := nodes[randNode].IPAddress + ":" + port

	return endpoint, nil
}

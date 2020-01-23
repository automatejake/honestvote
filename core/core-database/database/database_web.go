package database

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetElections() ([]Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var elections []Election
	var election Election

	query := bson.M{"transaction.type": "Election"}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		annoying_mongo_form := block.Transaction.(primitive.D)
		mapstructure.Decode(annoying_mongo_form.Map(), &election)

		elections = append(elections, election)

	}

	result.Close(context.TODO())
	return elections, nil
}

func GetElection(election_signature string) (Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var election Election

	query := bson.M{"transaction.type": "Election", "transaction.signature": election_signature}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

	}

	annoying_mongo_form := block.Transaction.(primitive.D)

	mapstructure.Decode(annoying_mongo_form.Map(), &election)
	result.Close(context.TODO())

	return election, nil
}

func GetVotes(election_signature string) ([]Vote, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block

	var votes []Vote
	var vote Vote

	query := bson.M{"transaction.type": "Vote", "transaction.election": election_signature}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}
		annoying_mongo_form := block.Transaction.(primitive.D)
		mapstructure.Decode(annoying_mongo_form.Map(), &vote)

		votes = append(votes, vote)

	}

	result.Close(context.TODO())
	return votes, nil
}

func GetPermissions(public_key string) ([]string, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration
	var elections []string

	query := bson.M{"transaction.type": "Registration", "transaction.sender": public_key}
	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		annoying_mongo_form := block.Transaction.(primitive.D)
		mapstructure.Decode(annoying_mongo_form.Map(), &registration)

		elections = append(elections, registration.Election)

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
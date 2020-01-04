package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateBlockchain(client *mongo.Client, block Block) bool {
	//Make the block a document and add it to local database
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	document := Block{
		Index:       block.Index,
		Timestamp:   block.Timestamp,
		Transaction: block.Transaction,
		Hash:        block.Hash,
		PrevHash:    block.PrevHash,
		Signatures:  block.Signatures,
	}

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return false
	}

	return true
}

func UpdateElections(client *mongo.Client, election Election) bool {
	//Make the block a document and add it to local database
	collection := client.Database("honestvote").Collection(CollectionPrefix + "elections")

	document := Election{
		Name:             election.Name,
		RegisteredVoters: election.RegisteredVoters,
		Start:            election.Start,
		End:              election.End,
		Positions:        election.Positions,
	}

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return false
	}

	return true
}

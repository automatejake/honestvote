package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
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
		Valid:       true,
	}

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return false
	}

	return true
}

func GrabPort(client *mongo.Client, id string) int {
	var result Node
	collection := client.Database("honestvote").Collection(CollectionPrefix + "connections")
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)

	if err == nil {
		return result.Port
	}

	return 0
}

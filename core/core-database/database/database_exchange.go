package database

import (
	"context"
	"fmt"

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
	}

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return false
	}

	return true
}

func LastIndex(client *mongo.Client) int64 {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	index, err := collection.CountDocuments(context.TODO(), bson.M{})

	if err == nil {
		GrabDocuments(client, 1)

		return index
	}

	return 0
}

func GrabDocuments(client *mongo.Client, old_index int64) {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	result, err := collection.Find(context.TODO(), bson.M{"index": bson.M{"$gt": old_index}})

	if err != nil {
		return
	}

	for result.Next(context.TODO()) {
		var document bson.M
		err = result.Decode(&document)
		fmt.Println(document)
	}
}

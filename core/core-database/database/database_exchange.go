package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateBlockchain(client *mongo.Client, block Block) {
	//Make the block a document and add it to local database
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	document := Block{
		Index:       block.Index,
		Timestamp:   block.Timestamp,
		Transaction: block.Transaction,
		Hash:        block.Hash,
		PrevHash:    block.PrevHash,
		Validator:   block.Validator,
	}

	result, err := collection.InsertOne(context.TODO(), document)

	if err == nil {
		fmt.Println(result)
	}
}

package database

import (
	"context"
	"fmt"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CheckVote(client *mongo.Client) {

	var object bson.M

	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	result, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println("Database Error! ", err)
	}

	for result.Next(context.TODO()) {
		result.Decode(&object)

		transcation := object["transaction"]

		if t, ok := transcation.(primitive.M); ok {
			if t["type"] == "Vote" && t["sender"] == "0xcheese" {
				fmt.Println(t)
			}
		}
	}
}

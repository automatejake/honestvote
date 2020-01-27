package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateBlockchain(client *mongo.Client, block Block) bool {
	//Make the block a document and add it to local database
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), block)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return false
	}

	return true
}

package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBlock(block Block) error {
	//Make the block a document and add it to local database
	collection := MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), block)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err)
		return err
	}

	return nil
}

func LastIndex(client *mongo.Client) int64 {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	index, err := collection.CountDocuments(context.TODO(), bson.M{})

	if err == nil {
		return index
	}

	return 0
}

func UpdateMongo(client *mongo.Client, data Block) error {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), data)

	return err
}

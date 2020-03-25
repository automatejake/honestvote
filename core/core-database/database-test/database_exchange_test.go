package main

import (
	"context"
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAddBlock(t *testing.T) {
	//gets the blockchain length, add a block, then makes sure the current blockchain length is one higher

	database.MongoDB = database.MongoConnect("localhost")

	//collection := database.MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")
	collection := database.MongoDB.Database("honestvote").Collection("blockchain")
	index, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		var block1 database.Block
		database.AddBlock(block1)
	}
	index2, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		if index2 != index+1 {
			t.Error()
		}

	}

}
func TestLastIndex(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	//1) test that the latest index is the latest index

	//collection := database.MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")
	collection := database.MongoDB.Database("honestvote").Collection("blockchain")
	index, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		got := database.LastIndex(database.MongoDB)
		if got != index {
			t.Error()
		}
	}

	//2) add a block, make sure the latest index is correct

	var block1 database.Block
	database.AddBlock(block1)
	index2, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		if index2 != index+1 {
			t.Error()
		}
	}

}
func TestUpdateMongo(t *testing.T) {
	//database.UpdateMongo()

}

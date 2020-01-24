package database

import (
	"context"
	"encoding/json"
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

//Find the document you're looking for
func FindDocument(client *mongo.Client, collection string, info interface{}, dType string) interface{} {
	var document bson.M

	search := bson.D{info.(primitive.E)}

	c := client.Database("honestvote").Collection(collection)

	err := c.FindOne(context.TODO(), search).Decode(&document)

	if err != nil {
		fmt.Println("Document not found!")
		return nil
	}

	j, err := json.Marshal(document)
	if err == nil {
		if dType == "Vote" {
			vote := new(Vote)
			json.Unmarshal(j, &vote)
			fmt.Println(vote)
		}
	}

	return nil
}

func CheckVote(client *mongo.Client, search bson.D) {

	var object bson.M

	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	result := collection.FindOne(context.TODO(), search)

	result.Decode(&object)

	transcation := object["transaction"]

	if t, ok := transcation.(primitive.M); ok {
		fmt.Println(t["type"])
	}
}

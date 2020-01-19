package database

import (
	"context"
	"fmt"

	"github.com/jneubaum/honestvote/tests/logger"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetElection(election_signature string) Election {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var election Election

	query := bson.M{"transaction.signature": election_signature}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("web_requests", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("web_requests", "GetElection", err.Error())
		}

	}

	switch t := block.Transaction.(type) {
	default:
		fmt.Printf("%T\n", t)
	}
	a := block.Transaction.(primitive.D)

	mapstructure.Decode(a.Map(), &election)
	result.Close(context.TODO())

	return election
}

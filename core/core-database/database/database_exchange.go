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

func AddTransaction(transaction interface{}, tranType string) error {
	//Make the block a document and add it to local database
	collection := MongoDB.Database("honestvote").Collection(CollectionPrefix + tranType)

	_, err := collection.InsertOne(context.TODO(), transaction)

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

func UpdateBlockMongo(client *mongo.Client, data Block) error {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockMongo()", err)
	}

	return err
}

func UpdateElectionMongo(client *mongo.Client, election Election) error {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "elections")

	_, err := collection.InsertOne(context.TODO(), election)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateElectionMongo()", err)
	}

	return err
}

func UpdateRegistrationMongo(client *mongo.Client, registration Registration) error {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "registrations")

	_, err := collection.InsertOne(context.TODO(), registration)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateRegistrationMongo()", err)
	}

	return err
}

func UpdateVoteMongo(client *mongo.Client, vote Vote) error {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "votes")

	_, err := collection.InsertOne(context.TODO(), vote)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateVoteMongo()", err)
	}

	return err
}

func GrabElectionsInBlock(index int) ([]Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "elections")

	var election Election
	var elections []Election

	query := bson.M{"blockIndex": index}
	result, err := collection.Find(context.TODO(), query)

	if err != nil {
		logger.Println("database_exchange.go", "GrabElectionsInBlock()", err)
	}

	for result.Next(context.TODO()) {
		err := result.Decode(&election)
		if err == nil {
			elections = append(elections, election)
		}
	}

	return elections, nil
}

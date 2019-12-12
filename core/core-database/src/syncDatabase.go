package coredb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Get all the mongoDB data to send over to a full node or peer node that asked for it
func GatherMongoData(client *mongo.Client, filter bson.M) []Candidate {
	//mongodump on peer, send created file to remote node.

	//mongo restore on remote node
	var Candidates []Candidate
	collection := client.Database("test_database").Collection("test_collection")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var candidate Candidate
		err = cur.Decode(&candidate)
		if err != nil {
			log.Fatal(err)
		}

		Candidates = append(Candidates, candidate)
	}

	return Candidates
}

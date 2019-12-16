package database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

//Send the data to the full/peer node
func MoveDocuments(peers []Peer) {
	MongoDB = MongoConnect()
	MongoData := GatherMongoData(MongoDB, bson.M{})
	buffer := new(bytes.Buffer)
	tmpArray := MongoData
	js := json.NewEncoder(buffer)
	err := js.Encode(tmpArray)
	if err == nil {
		for _, socket := range peers {
			fmt.Println("Sending documents.")
			socket.Socket.Write(append([]byte("recieve data "), buffer.Bytes()...))
		}
	}
}

func UpdateMongo(client *mongo.Client, data []Candidate) {
	collection := client.Database("new_database").Collection("new_collection")

	var ui []interface{}
	for _, candidate := range data {
		ui = append(ui, candidate)
	}

	collection.InsertMany(context.TODO(), ui)
}

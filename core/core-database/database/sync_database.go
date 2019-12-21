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
func GatherMongoData(client *mongo.Client, filter bson.M, database_name string, collection_name string) []Candidate {
	var Candidates []Candidate
	collection := client.Database(database_name).Collection(collection_name)

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
func MoveDocuments(nodes []TempNode, database_name string, collection_name string) {

	MongoData := GatherMongoData(MongoDB, bson.M{}, database_name, collection_name)
	buffer := new(bytes.Buffer)
	tmpArray := MongoData
	js := json.NewEncoder(buffer)
	err := js.Encode(tmpArray)
	if err == nil {
		for _, socket := range nodes {
			fmt.Println("Sending documents.")
			socket.Socket.Write(append([]byte("recieve data "), buffer.Bytes()...))
		}
	}
}

func UpdateMongo(client *mongo.Client, data []Candidate, database_name string, collection_name string) {
	collection := client.Database(database_name).Collection(collection_name)

	var ui []interface{}
	for _, candidate := range data {
		ui = append(ui, candidate)
	}

	collection.InsertMany(context.TODO(), ui)
}

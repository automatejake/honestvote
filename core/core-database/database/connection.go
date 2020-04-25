package database

import (
	"context"
	"log"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client
var Collection *mongo.Collection

//Connect to MongoDB
func MongoConnect(remoteip string) *mongo.Client {
	uri := "mongodb://" + remoteip + ":27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.Println("connection.go", "MongoConnect()", err)
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Println("connection.go", "MongoConnect()", err)
	}

	return client
}

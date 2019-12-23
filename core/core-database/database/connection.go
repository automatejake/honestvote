package database

import (
	"context"
	"log"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

//Connect to MongoDB
func MongoConnect() *mongo.Client {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Println("connection.go", "MongoConnect()", err.Error())
	}

	return client
}

package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
)

func AddBlock(block Block) error {
	//Make the block a document and add it to local database
	collection := MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), block)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return err
	}

	return nil
}

package database

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CorrespondingRegistration(v Vote) Registration {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration

	query := bson.M{"transaction.type": "Registration", "transaction.signature": v.Election, "transaction.sender": v.Sender}

	result := collection.FindOne(context.TODO(), query)
	result.Decode(&block)

	annoying_mongo_form := block.Transaction.(primitive.D)

	mapstructure.Decode(annoying_mongo_form.Map(), &registration)

	return registration
}

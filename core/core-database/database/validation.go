package database

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CorrespondingRegistration(v Vote) Registration {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration

	query := bson.M{"transaction.type": "Registration", "transaction.election": v.Election, "transaction.receiver": v.Sender}

	result := collection.FindOne(context.TODO(), query)
	result.Decode(&block)

	annoying_mongo_form := block.Transaction.(primitive.D)
	mapstructure.Decode(annoying_mongo_form.Map(), &registration)

	return registration
}

func ContainsRegistration(sender PublicKey, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration

	query := bson.M{"transaction.type": "Registration", "transaction.sender": sender, "transaction.election": election}
	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&block)
	if err != nil {
		return false
	}

	annoying_mongo_form := block.Transaction.(primitive.D)
	err = mapstructure.Decode(annoying_mongo_form.Map(), &registration)
	if err != nil {

	}
	return true
}

func ContainsVote(sender PublicKey, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var vote Vote

	query := bson.M{"transaction.type": "Vote", "transaction.sender": sender, "transaction.election": election}
	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&block)
	if err != nil {
		return false
	}

	annoying_mongo_form := block.Transaction.(primitive.D)
	err = mapstructure.Decode(annoying_mongo_form.Map(), &vote)
	if err != nil {

	}
	return true
}
func MarkDishonestNode(n Node) {
	n.Role = "bad actor"
	fmt.Println(n)

}

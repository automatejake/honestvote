package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CorrespondingRegistration(v Vote) Registration {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration

	query := bson.M{"transaction.type": "Registration", "transaction.electionId": v.Election, "transaction.receiver": v.Sender}
	result := collection.FindOne(context.TODO(), query)
	result.Decode(&block)

	annoying_mongo_form := block.Transaction.(primitive.D)
	err := mapstructure.Decode(annoying_mongo_form.Map(), &registration)
	if err != nil {
		logger.Println("validation.go", "CorrespondingRegistration", err.Error())
	}

	return registration
}

func ContainsRegistration(receiver string, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration

	query := bson.M{"transaction.type": "Registration", "transaction.receiver": receiver, "transaction.electionId": election}
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

func ContainsVote(sender string, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var vote Vote

	query := bson.M{"transaction.type": "Vote", "transaction.sender": sender, "transaction.electionId": election}
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
func MarkDishonestNode(n Node) error {
	n.Role = "bad actor"
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"publickey": n.PublicKey},
		n,
	)
	if err != nil {
		return err
	}

	return nil
}

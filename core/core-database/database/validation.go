package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"

	"go.mongodb.org/mongo-driver/bson"
)

func CorrespondingRegistration(v Vote) Registration {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "registrations")
	var registration Registration

	query := bson.M{"electionId": v.Election, "receiver": v.Sender}
	result := collection.FindOne(context.TODO(), query)
	result.Decode(&registration)

	return registration
}

func ContainsRegistration(receiver string, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "registrations")
	var registration Registration

	query := bson.M{"receiver": receiver, "electionId": election}
	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&registration)
	if err != nil {
		return false
	}

	return true
}

func ContainsVote(sender string, election string) bool {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "votes")
	var vote Vote

	query := bson.M{"sender": sender, "electionId": election}
	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&vote)
	if err != nil {
		return false
	}

	return true
}

func CheckElectionSignature(sig string) bool {
	var election Election

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "elections")

	query := bson.M{"signature": sig}

	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&election)

	if err == nil {
		return false
	}

	return true
}

func EditNodeRole(n Node, role string) error {
	n.Role = role
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"publickey": n.PublicKey},
		n,
	)
	if err != nil {
		logger.Println("validation.go", "EditNodeRole()", err)
		return err
	}

	return nil
}

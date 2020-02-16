package database

import (
	"context"
	"time"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveRegistrationCode(registrant AwaitingRegistration) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	_, err := collection.InsertOne(context.TODO(), registrant)
	if err != nil {
		logger.Println("email_registration.go", "SaveRegistrationCode()", err.Error())
	}
	logger.Println("email_registration.go", "SaveRegistrationCode()", "inserted document successfully")
}

func IsValidRegistrationCode(code string) (AwaitingRegistration, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	query := bson.M{"code": code}

	// retrieve code from database
	var result AwaitingRegistration
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			logger.Println("routing_table.go", "ExistsInTable()", err.Error())
		}
		return AwaitingRegistration{}, err
	}

	// determine if registration code is young enough
	linkAge, err := time.Parse(time.RFC1123, result.Timestamp)
	if err != nil {
		logger.Println("email_registration.go", "IsValidRegistrationCode()", err.Error())
	}

	expiryTime := 4 * time.Hour
	if linkAge.Add(expiryTime).After(time.Now()) {
		customErr := &CustomError{
			Time:    time.Now(),
			Message: "Registration Code is too young",
		}
		return AwaitingRegistration{}, customErr
	}

	// make sure that public key is correct

	// make sure that election is still ongoing / valid

	return result, nil
}

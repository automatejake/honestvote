package database

import (
	"context"
	"time"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckEmailVerification(registration AwaitingRegistration) error {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	query := bson.M{"email": registration.Email, "verified": "true"}
	registrate := registration
	err := collection.FindOne(context.TODO(), query).Decode(&registrate)
	logger.Println("email_registration.go", "CheckEmailVerification", err)
	if err == nil {
		logger.Println("email_registration.go", "CheckEmailVerification", err.Error())
		return err
	} else {
		collection.UpdateOne(context.TODO(), query, bson.D{{"$set", bson.D{{"verified", "true"}}}})
	}
	return nil
}

func SaveRegistrationCode(registrant AwaitingRegistration) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	_, err := collection.InsertOne(context.TODO(), registrant)
	if err != nil {
		logger.Println("email_registration.go", "SaveRegistrationCode()", err.Error())
	}
	logger.Println("email_registration.go", "SaveRegistrationCode()", "inserted document successfully")
}

func RemoveRegistrationCode(registrant AwaitingRegistration) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	_, err := collection.DeleteOne(context.TODO(), registrant)
	if err != nil {
		logger.Println("email_registration.go", "RemoveRegistrationCode()", err.Error())
	}
	logger.Println("email_registration.go", "RemoveRegistrationCode()", "removed document successfully")
}

func IsValidRegistrationCode(code string) (AwaitingRegistration, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	query := bson.M{"code": code}

	// retrieve code from database
	var result AwaitingRegistration
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			logger.Println("email_registration.go", "IsValidRegistrationCode()", err.Error())
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

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	a := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	// b := time.Now().String()
	start, err := time.Parse(time.RFC1123, a)
	if err != nil {
		log.Println("Error is: ", err)
	}

	var HOURS float64 = 4
	if time.Now().Sub(start).Hours() > HOURS {

	}

	// Mon, 02 Jan 2006 15:04:05 MST
	// start, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", result.Timestamp)
	// if err != nil {
	// 	log.Println("This is the error: ", err) // logger.Println("email_registration.go", "IsValidRegistrationCode()", err.Error())
	// }

	fmt.Println(time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	// _, _ = IsValidRegistrationCode("b8442edc49bc01d335d71b0a7b4c92f23caedc5b3d2d6cbeb09389784d01288572dd6dc4e6a6891b149ff5e07e6969d0c5e12302b3cc733d9f99850149d4a77c7098057cd0741f83850fbfbeb6bb3c2439f7a8c2f7d6a3bfe93b5dce4936f0ef77f11d7d")

}

var DatabaseName string = "honestvote"
var CollectionPrefix string = "a_"
var EmailRegistrants string = "email_registrants"
var MongoDB *mongo.Client = database.MongoConnect()

func IsValidRegistrationCode(code string) (string, bool) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + EmailRegistrants)
	query := bson.M{"code": code}

	// retrieve code from database
	var result database.AwaitingRegistration
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			logger.Println("routing_table.go", "ExistsInTable()", err.Error())
		}
		return "no registration code exists", false
	}

	// determine if registration code is young enoughtime.RFC1123
	a := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	fmt.Println(result.Timestamp)
	fmt.Println(a)

	start, err := time.Parse(time.RFC1123, a)
	if err != nil {
		log.Println("This is the error: ", err) // logger.Println("email_registration.go", "IsValidRegistrationCode()", err.Error())
	}

	var HOURS float64 = 4
	if time.Now().Sub(start).Hours() > HOURS {
		return "registration code has expired", false
	}

	// make sure that public key is correct

	// make sure that election is still ongoing / valid

	return result.PublicKey, true
}

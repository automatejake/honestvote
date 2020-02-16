package database

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetElections() ([]Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var elections []Election
	var election Election

	query := bson.M{"transaction.type": "Election"}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		annoying_mongo_form := block.Transaction.(primitive.D).Map()
		mapstructure.Decode(annoying_mongo_form, &election)

		election.Start = annoying_mongo_form["startDate"].(string)
		election.End = annoying_mongo_form["endDate"].(string)
		election.Institution = annoying_mongo_form["institutionName"].(string)

		elections = append(elections, election)

	}

	result.Close(context.TODO())
	return elections, nil
}

func GetElection(election_signature string) (Election, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var election Election

	query := bson.M{"transaction.type": "Election", "transaction.signature": election_signature}

	result := collection.FindOne(context.TODO(), query)

	err := result.Decode(&block)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
		return election, err
	}

	annoying_mongo_form := block.Transaction.(primitive.D).Map()
	mapstructure.Decode(annoying_mongo_form, &election)

	election.Start = annoying_mongo_form["startDate"].(string)
	election.End = annoying_mongo_form["endDate"].(string)
	election.Institution = annoying_mongo_form["institutionName"].(string)

	election.Positions = nil

	if tran, ok := block.Transaction.(primitive.D); ok {
		tranMap := tran.Map()

		if pos, ok := tranMap["positions"].(primitive.A); ok {
			for _, position := range pos {
				var tempPos Position

				if posInfo, ok := position.(primitive.D); ok {
					posMap := posInfo.Map()

					tempPos.Name = posMap["displayName"].(string)
					tempPos.PositionId = posMap["id"].(string)

					if cand, ok := posMap["candidates"].(primitive.A); ok {
						for _, candidate := range cand {
							var tempCand Candidate

							if candInfo, ok := candidate.(primitive.D); ok {
								candMap := candInfo.Map()

								tempCand.Name = candMap["name"].(string)
								tempCand.Recipient = candMap["key"].(string)

								tempPos.Candidates = append(tempPos.Candidates, tempCand)
							}
						}
					}

					election.Positions = append(election.Positions, tempPos)
				}
			}
		}
	}

	return election, nil
}

func GetVotes(election_signature string) ([]Vote, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block

	var votes []Vote
	var vote Vote

	query := bson.M{"transaction.type": "Vote", "transaction.electionName": election_signature}

	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		annoying_mongo_form := block.Transaction.(primitive.D)
		mapstructure.Decode(annoying_mongo_form.Map(), &vote)

		vote.Receiver = nil

		if tran, ok := block.Transaction.(primitive.D); ok {
			tranMap := tran.Map()

			if votes, ok := tranMap["receiver"].(primitive.A); ok {
				for _, v := range votes {

					var candidate SelectedCandidate
					if obj, ok := v.(primitive.D); ok {

						voteInfo := obj.Map()
						candidate.PositionId = voteInfo["positionid"].(string)
						candidate.Recipient = voteInfo["recipient"].(string)
					}
					vote.Receiver = append(vote.Receiver, candidate)
				}
			}
		}

		votes = append(votes, vote)

	}

	result.Close(context.TODO())
	return votes, nil
}

func GetPermissions(public_key string) ([]string, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "blockchain")
	var block Block
	var registration Registration
	var elections []string

	query := bson.M{"transaction.type": "Registration", "transaction.receiver": public_key}
	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetElection", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		annoying_mongo_form := block.Transaction.(primitive.D)
		err = mapstructure.Decode(annoying_mongo_form.Map(), &registration)
		if err != nil {
			logger.Println("database_web", "GetElection", err.Error())
		}

		elections = append(elections, registration.Election)

	}

	result.Close(context.TODO())
	return elections, nil
}

func GetEndpoint() (string, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + "node_list")

	var nodes []Node
	var node Node

	query := bson.M{}
	result, err := collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("database_web", "GetEndpoint", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&node)
		if err != nil {
			logger.Println("database_web", "GetEndpoint", err.Error())
		}
		nodes = append(nodes, node)
	}

	randNode := rand.Intn(len(nodes))
	port := strconv.Itoa(nodes[randNode].Port)
	endpoint := nodes[randNode].IPAddress + ":" + port

	return endpoint, nil
}

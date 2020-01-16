package database

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCandidates() Candidate {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	// query := bson.M{"$or": bson.A{bson.M{"ipaddress": bson.M{"$ne": requesting_node.IPAddress}}, bson.M{"port": bson.M{"$ne": requesting_node.Port}}}}

	// result, err := collection.Find(context.TODO(), query)
	// if err != nil {
	// 	logger.Println("routing_table.go", "FindNode", err.Error())
	// }

	// for result.Next(context.TODO()) {
	// 	var peer Node
	// 	err = result.Decode(&peer)
	// 	if err != nil {
	// 		logger.Println("routing_table.go", "FindNode", err.Error())
	// 	}

	// 	peers = append(peers, peer)
	// }

	// // Close the cursor once finished
	// result.Close(context.TODO())

	// return peers
	return Candidate{}
}

func GetElections() []Election {

	fmt.Println(reflect.TypeOf(Collection))
	return nil
}

func GetElection(electionid string) API_Election {
	Collection = MongoDB.Database("honestvote").Collection(CollectionPrefix + ElectionHistory)
	query := bson.M{"electionid": electionid}

	result, err := Collection.Find(context.TODO(), query)
	if err != nil {
		logger.Println("routing_table.go", "FindNode", err.Error())
	}

	fmt.Println(reflect.TypeOf(Collection))
	fmt.Println(reflect.TypeOf(result))

	return API_Election{}
}

func GetVoters() API_Voter {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	return API_Voter{}
}

func GetPositions() []API_ElectionPosition {
	return nil
}

func GetTickets() []TicketEntry {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	return nil
}

package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func ExistsInTable(ipaddr string, port string) bool {
	// data, err := ioutil.ReadFile("routingtable.txt")

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + ElectionHistory)

	var result Peer
	err := collection.FindOne(context.TODO(), ipaddr).Decode(&result)
	if err != nil {
		log.Print("Afraid weve reached an impasse: ", err)
		return false
	}
	fmt.Println()

	return true
}

func AddToTable(ipaddr string, port int) {

	newPeer := Peer{
		IPAddress: ipaddr,
		Port:      port,
		// Role:      role,
	}

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	result, err := collection.InsertOne(context.TODO(), newPeer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}

func FindPeers() []Peer {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	var peers []Peer

	result, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.TODO()) {
		var peer Peer
		err = result.Decode(&peer)
		if err != nil {
			log.Fatal(err)
		}

		peers = append(peers, peer)
	}

	// Close the cursor once finished
	// result.Close(context.TODO())

	return peers
}

func FindFullNode() {

}

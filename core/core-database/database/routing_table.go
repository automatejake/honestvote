package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func ExistsInTable(ipaddr string, port int) bool {
	// data, err := ioutil.ReadFile("routingtable.txt")

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	query := bson.M{"ipaddress": ipaddr, "port": port}

	var result Peer
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			log.Println("File: routing_table.go\nFunction:ExistsInTable\n", err)
		}
		log.Println("No documents in result")
		return false
	}

	log.Println("Exists")
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

func FindPeers(exclude_requesting_peer Peer) []Peer {
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
	result.Close(context.TODO())

	return peers
}

func FindFullNode() {

}

package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

/**
* Exist in table - Simple Process
*
* 1) Checks to see if the given connection exists in the table of connections
*
**/
func ExistsInTable(ipaddr string, port int) bool {

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

/**
* Add to Table - Simple Process
*
* 1) Adds the node to the database of connections
**/
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

/**
* Find Peers - 2 Step Process
*
* 1) Query database for all peers besides requesting peer
* 2) Return list of peers to the requesting peer
*
**/
func FindPeers(requesting_peer Peer) []Peer {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	var peers []Peer

	// Mongo shell format:
	// {$or: [ { ipaddress: { $ne: "127.0.0.1" } },{ port: { $ne: 7002 } }]}
	query := bson.M{"$or": bson.A{bson.M{"ipaddress": bson.M{"$ne": requesting_peer.IPAddress}}, bson.M{"port": bson.M{"$ne": requesting_peer.Port}}}}

	result, err := collection.Find(context.TODO(), query)
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

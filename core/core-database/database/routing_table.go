package database

import (
	"context"
	"log"
)

func ExistsInTable(ipaddr string, port string) bool {
	// data, err := ioutil.ReadFile("routingtable.txt")

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + ElectionHistory)

	var result Peer
	err := collection.FindOne(context.TODO(), ipaddr).Decode(&result)
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func AddToTable(ipaddr string, port int, role string, connections int) {

	newPeer := Peer{
		IPAddress:         ipaddr,
		Port:              port,
		Role:              role,
		NumberConnections: connections,
	}

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	_, err := collection.InsertOne(context.TODO(), newPeer)
	if err != nil {
		log.Fatal(err)
	}

}

package database

import (
	"context"
	"log"
	"net"
)

func ExistsInTable(ipaddr string, port int, database_name string, collection_name string) bool {
	// data, err := ioutil.ReadFile("routingtable.txt")

	collection := MongoDB.Database(database_name).Collection(collection_name)

	var result Peer
	err := collection.FindOne(context.TODO(), ipaddr).Decode(&result)
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func AddToTable(ipaddr string, port int, socket net.Conn, role string, database_name string, collection_name string) {

	newPeer := Peer{
		IPAddress: ipaddr,
		Port:      port,
		Socket:    socket,
		Role:      role,
	}

	collection := MongoDB.Database(database_name).Collection(collection_name)
	_, err := collection.InsertOne(context.TODO(), newPeer)
	if err != nil {
		log.Fatal(err)
	}

}

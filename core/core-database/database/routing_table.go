package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
)

/**
* Exist in table - Simple Process
*
* 1) Checks to see if the given connection exists in the table of connections
*
**/
func DoesNodeExist(node Node) bool {

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	query := bson.M{"publickey": node.PublicKey}
	// query := bson.M{"publickey": node.PublicKey}

	var result Node
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			logger.Println("routing_table.go", "ExistsInTable()", "Node does not exist: "+err.Error())
		}
		return false
	}

	logger.Println("routing_table.go", "ExistsInTable()", "Node already exists")
	return true
}

/**
* Add to Table - Simple Process
*
* 1) Adds the node to the database of connections
**/
func AddNode(newNode Node) {

	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
	_, err := collection.InsertOne(context.TODO(), newNode)
	if err != nil {
		logger.Println("routing_table.go", "AddNode", err.Error())
	}

}

/**
* Find Nodes - 2 Step Process
*
* 1) Query database for all Nodes besides requesting Node
* 2) Return list of Nodes to the requesting Node
*
**/
func FindNodes() []Node {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	var peers []Node

	result, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.Println("routing_table.go", "FindNode", err.Error())
	}

	for result.Next(context.TODO()) {
		var peer Node
		err = result.Decode(&peer)
		if err != nil {
			logger.Println("routing_table.go", "FindNode", err.Error())
		}

		peers = append(peers, peer)
	}

	// Close the cursor once finished
	result.Close(context.TODO())

	return peers
}

func FindNode(public_key string) (Node, error) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	var node Node

	query := bson.M{"publickey": public_key}
	result := collection.FindOne(context.TODO(), query)
	err := result.Decode(&node)
	if err != nil {
		logger.Println("routing_table.go", "FindNode()", err)
		return Node{}, err
	}

	return node, nil
}

func DeleteNode(node Node) {
	collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)

	query := bson.M{"ipaddress": node.IPAddress, "port": node.Port}
	_, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		logger.Println("routing_table.go", "DeleteNode", err.Error())
	}

}

func ConnectFullNode() {

}

package main

import (
	"context"
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDoesNodeExist(t *testing.T) {
	//Test case 1: passing in a node that does exist
	database.MongoDB = database.MongoConnect("localhost")
	var node1 database.Node
	//database.AddNode(node1) //node is added
	got := database.DoesNodeExist(node1)
	if got != true {
		t.Error()
	}
	//test case 2: passing in a invalid node

}

func TestAddNode(t *testing.T) {
	//gets the node_list length, add a node, then makes sure the current node_list length is one higher
	database.MongoDB = database.MongoConnect("localhost")

	//collection := database.MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")
	collection := database.MongoDB.Database("honestvote").Collection("node_list")
	index, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		var node1 database.Node
		database.AddNode(node1)
	}
	index2, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		if index2 != index+1 {
			t.Error()
		}

	}

}

func TestFindNodes(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	nodes := database.FindNodes()

	if nodes == nil {
		t.Log("There were no nodes returned from the database, this could be an error.")
	}

	t.Log("There were no errors when looking for nodes from the database: ", nodes)
}

func TestFindNode(t *testing.T) {
	database.MongoDB = database.MongoConnect("localhost")

	node, err := database.FindNode(" ") //TODO: This should not be empty

	if err != nil {
		t.Error("There shouldn't be an error when trying to grab the node from the database. Error: ", err)
	}

	t.Log("There wasn't an error when grabbing the node: ", node)

}

func TestDeleteNode(t *testing.T) { //Fails

	database.MongoDB = database.MongoConnect("localhost")

	var nodex database.Node
	database.AddNode(nodex)

	//collection := database.MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")
	collection := database.MongoDB.Database("honestvote").Collection("node_list")
	index, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		t.Error()
	}

	database.DeleteNode(nodex)

	index2, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err == nil {
		if index2 != index-1 {
			t.Error()
		}

	}

}

// func TestConnectFullNode(t *testing.T) {
// 	database.ConnectFullNode()
// }

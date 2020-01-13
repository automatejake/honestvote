package database

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

func GetElections() {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
}

func GetElection() {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
}

func GetVoters() {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
}

func GetTickets() {
	// collection := MongoDB.Database(DatabaseName).Collection(CollectionPrefix + Connections)
}

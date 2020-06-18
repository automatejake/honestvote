package p2p

import (
	"encoding/json"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Adds new connection to database and local Node array
func AcceptConnectMessage(node database.Node, conn net.Conn) {

	node.IPAddress = conn.RemoteAddr().String()
	if !database.DoesNodeExist(node) {
		database.AddNode(node)
	}

	Nodes = append(Nodes, conn)

}

//Decoding the data sent from another peer, this data is from a database
func DecodeBlockData(data []byte) {
	var block database.Block

	err := json.Unmarshal(data, &block)
	if err != nil {
		logger.Println("read_functions.go", "DecodeData()", err)
		return
	}

	PreviousBlock = block
	database.UpdateBlockMongo(database.MongoDB, block)
}

func DecodeTransactionData(data []byte, tranType string) {

	switch tranType {
	case "elections":
		var election database.Election

		err := json.Unmarshal(data, &election)
		if err != nil {
			logger.Println("read_functions.go", "DecodeData()", err)
			return
		}

		database.UpdateElectionMongo(database.MongoDB, election)
	case "registrations":
		var registration database.Registration

		err := json.Unmarshal(data, &registration)
		if err != nil {
			logger.Println("read_functions.go", "DecodeData()", err)
			return
		}
		database.UpdateRegistrationMongo(database.MongoDB, registration)
	case "votes":
		var vote database.Vote

		err := json.Unmarshal(data, &vote)
		if err != nil {
			logger.Println("read_functions.go", "DecodeData()", err)
			return
		}
		database.UpdateVoteMongo(database.MongoDB, vote)
	}
}

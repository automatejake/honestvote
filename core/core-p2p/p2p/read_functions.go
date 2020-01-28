package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Adds new connection to database and local Node array
func AcceptConnectMessage(node database.Node, conn net.Conn) {

	node.IPAddress = conn.RemoteAddr().String()[0:9]
	if !database.DoesNodeExist(node) {
		database.AddNode(node)
	}

	Nodes = append(Nodes, conn)

	fmt.Println(Nodes)
}

//Decoding the data sent from another peer, this data is from a database
func DecodeData(buffer *bytes.Buffer) {
	tmpArray := new([]database.Block)
	js := json.NewDecoder(buffer)
	err := js.Decode(tmpArray)
	if err == nil {
		database.UpdateMongo(database.MongoDB, *tmpArray)
	}
}

//Get vote from full node and turn it into a block and propose
func ReceiveTransaction(mType string, data []byte) error {

	fmt.Println("recieved")
	var valid bool
	valid = true
	switch mType {
	case "Vote":
		var vote database.Vote
		json.Unmarshal(data, &vote)
		if valid {
			AddToBlock(data)
		} else {
			logger.Println("", "", "")
		}
		// valid, err = validation.IsValidVote(vote)
	case "Election":
		fmt.Println("identified election")
		election := &database.Election{}
		json.Unmarshal(data, election)
		// valid, err = validation.IsValidElection(election)
		if valid {
			AddToBlock(election)
		} else {
			logger.Println("", "", "")
		}
	case "Registration":
		registration := &database.Registration{}
		json.Unmarshal(data, &registration)
		// valid, err = validation.IsValidRegistration(registration)
		if valid {
			AddToBlock(registration)
		} else {
			logger.Println("", "", "")
		}
	}

	return nil
}

func AddToBlock(transaction interface{}) {
	block := consensus.GenerateBlock(PreviousBlock, transaction, PublicKey)

	fmt.Println("created block")
	//Check if there is a proposed block currently, if so, add to the queue

	logger.Println("peer_routes.go", "HandleConn()", "Empty, proposing this block.")

	err := database.AddBlock(block)
	if err != nil {
	} else {
		PreviousBlock = block
		ProposeBlock(block)
	}

}

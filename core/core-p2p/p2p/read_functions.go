package p2p

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-validation/validation"
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
func DecodeData(data []byte) {
	var block database.Block

	err := json.Unmarshal(data, &block)
	if err == nil {
		database.UpdateMongo(database.MongoDB, block)
	}
}

//Get vote from full node and turn it into a block and propose
func ReceiveTransaction(mType string, data []byte) error {

	fmt.Println("recieved")
	var valid bool
	switch mType {
	case "Vote":
		vote := &database.Vote{}
		err := json.Unmarshal(data, vote)
		if err != nil {

		}

		valid, err = validation.IsValidVote(*vote)
		if valid {
			AddToBlock(vote)
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}

	case "Election":
		election := &database.Election{}
		err := json.Unmarshal(data, election)
		if err != nil {

		}

		valid, err = validation.IsValidElection(*election)
		if valid {
			AddToBlock(election)
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}
	case "Registration":
		registration := &database.Registration{}
		err := json.Unmarshal(data, &registration)
		if err != nil {

		}

		valid, err = validation.IsValidRegistration(*registration)
		if valid {
			AddToBlock(registration)
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}
	}

	return nil
}

func AddToBlock(transaction interface{}) {
	block := consensus.GenerateBlock(PreviousBlock, transaction, PublicKey, PrivateKey)

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

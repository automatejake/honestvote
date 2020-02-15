package p2p

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
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

	// fmt.Println(Nodes)
}

//Decoding the data sent from another peer, this data is from a database
func DecodeData(data []byte) {
	var block database.Block

	err := json.Unmarshal(data, &block)
	if err != nil {
		return
	}
	// fmt.Println(block)
	PreviousBlock = block
	database.UpdateMongo(database.MongoDB, block)

}

//Get vote from full node and turn it into a block and propose
func ReceiveTransaction(mType string, data []byte) error {

	var valid bool
	switch mType {
	case "Vote":
		vote := &database.Vote{}
		err := json.Unmarshal(data, vote)
		if err != nil {

		}

		valid, err = validation.IsValidVote(*vote)

		if valid {
			// websocket.BroadcastVote(*vote)
			AddToBlock(vote, crypto.CalculateHash([]byte(vote.Signature)))
		} else {
			fmt.Println(err)
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}

	case "Election":
		election := &database.Election{}
		err := json.Unmarshal(data, election)
		if err != nil {

		}

		valid, err = validation.IsValidElection(*election)
		if valid {
			AddToBlock(election, crypto.CalculateHash([]byte(election.Signature)))
		} else {
			fmt.Println(err.Error())
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}
	case "Registration":
		fmt.Println("recieved registration")

		registration := &database.Registration{}
		err := json.Unmarshal(data, &registration)
		if err != nil {

		}

		valid, err = validation.IsValidRegistration(*registration)

		if valid {
			// websocket.SendRegistration(*registration)
			AddToBlock(registration, crypto.CalculateHash([]byte(registration.Signature)))
		} else {
			fmt.Println(err)
			logger.Println("read_functions.go", "RecieveTransaction()", err.Error())
		}
	}

	return nil
}

func AddToBlock(transaction interface{}, hash string) {
	block, err := consensus.GenerateBlock(PreviousBlock, transaction, PublicKey, PrivateKey)
	if err != nil {
		logger.Println("read_function.go", "AddToBlock()", err.Error())
	}

	block.MerkleRoot = hash

	fmt.Println("created block")
	//Check if there is a proposed block currently, if so, add to the queue

	logger.Println("peer_routes.go", "HandleConn()", "Empty, proposing this block.")

	fmt.Println("Transaction Type:\n", reflect.TypeOf(block.Transaction))

	err = database.AddBlock(block)
	if err != nil {
	} else {
		PreviousBlock = block
		ProposeBlock(block)
	}

}

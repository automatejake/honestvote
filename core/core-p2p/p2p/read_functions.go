package p2p

import (
	"encoding/hex"
	"encoding/json"
	"net"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-validation/validation"
	"github.com/jneubaum/honestvote/core/core-websocket/websocket"
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
func DecodeData(data []byte) {
	var block database.Block

	err := json.Unmarshal(data, &block)
	if err != nil {
		logger.Println("read_functions.go", "DecodeData()", err)
		return
	}

	PreviousBlock = block
	database.UpdateMongo(database.MongoDB, block)

}

//Get vote from full node and turn it into a block and propose
func ReceiveTransaction(mType string, data []byte) error {
	logger.Println("read_functions.go", "RecieveTransactions()", "Processing "+mType+" transaction")
	var valid bool
	switch mType {
	case "Vote":
		vote := &database.Vote{}
		err := json.Unmarshal(data, vote)
		if err != nil {
			logger.Println("read_functions.go", "RecieveTransactions()", err)
		}

		valid, err = validation.IsValidVote(*vote)

		if valid {
			logger.Println("read_functions.go", "RecieveTransaction()", "Passed validation")
			websocket.BroadcastVote(*vote)
			AddToBlock(vote, hex.EncodeToString(crypto.CalculateHash([]byte(vote.Signature))))
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err)
		}

	case "Election":
		election := &database.Election{}
		err := json.Unmarshal(data, election)
		if err != nil {
			logger.Println("read_functions.go", "RecieveTransactions()", err)
		}
		logger.Println("read_functions.go", "RecieveTransactions()", "Received transaction")
		logger.Println("read_functions.go", "RecieveTransactions()", election)

		valid, err = validation.IsValidElection(*election)
		if valid {
			AddToBlock(election, hex.EncodeToString(crypto.CalculateHash([]byte(election.Signature))))
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err)
		}
	case "Registration":

		registration := &database.Registration{}
		err := json.Unmarshal(data, &registration)
		if err != nil {
			logger.Println("read_functions.go", "RecieveTransactions()", err)
		}

		valid, err = validation.IsValidRegistration(*registration)

		if valid {
			logger.Println("", "", "Sending Registration")
			websocket.SendRegistration(*registration)
			AddToBlock(registration, hex.EncodeToString(crypto.CalculateHash([]byte(registration.Signature))))
		} else {
			logger.Println("read_functions.go", "RecieveTransaction()", err)
		}
	}

	return nil
}

func AddToBlock(transaction interface{}, hash string) {
	block, err := consensus.GenerateBlock(PreviousBlock, transaction, PublicKey, PrivateKey)
	if err != nil {
		logger.Println("read_function.go", "AddToBlock()", err.Error())
	}

	// block.MerkleRoot = hash

	//Check if there is a proposed block currently, if so, add to the queue

	logger.Println("peer_routes.go", "HandleConn()", "Empty, proposing this block.")

	err = database.AddBlock(block)
	if err != nil {
		logger.Println("read_functions.go", "AddToBlock()", err)
	} else {
		PreviousBlock = block
		ProposeBlock(block)
	}

}

package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-websocket/websocket"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Adds new connection to database and local Node array
func AcceptConnectMessage(node database.Node, conn net.Conn) {
	//ADD TO DATABASE AS WELL
	// Nodes[port] = true

	node.IPAddress = conn.RemoteAddr().String()[0:9]
	if !database.DoesNodeExist(node) {
		database.AddNode(node)
	}

	// var message Message
	// data, err := json.Marshal(Self)
	// if err != nil {
	// 	logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	// }
	// message.Message = "connect response"
	// message.Data = data
	// data, err = json.Marshal(message)
	// if err != nil {
	// 	logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	// }

	// conn.Write(data)

	Nodes = append(Nodes, conn)

	fmt.Println(Nodes)
}

//Decoding the data sent from another peer, this data is from a database
func DecodeData(buffer *bytes.Buffer) {
	tmpArray := new([]database.Candidate)
	js := json.NewDecoder(buffer)
	err := js.Decode(tmpArray)
	if err == nil {
		database.UpdateMongo(database.MongoDB, *tmpArray, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
	}
}

//Get vote from full node and turn it into a block and propose
func ReceiveTransaction(data []byte, mType string) {

	var transaction interface{}
	var blockType string

	switch mType {
	case "Vote":
		vote := database.Vote{Sender: "0xcheese", Receiver: map[string]string{"1": "0xsugar", "2": "0xpeanut"}}
		vote.Signature = consensus.CreateSignature(vote, PrivateKey)
		transaction = vote
		websocket.Broadcast(vote)
		blockType = "Vote"
	case "Register":
		register := database.Registration{Election: "0xelection", Sender: "0xadmin", Receiver: "0xcheese"}
		register.Signature = consensus.CreateSignature(register, PrivateKey)
		blockType = "Register"
	case "Election":
		//Temporary Variable, will be data unmarshalled
		election := database.Election{ElectionName: "WCU", Start: "3/23/2020", End: "3/30/2020"}
		election.Signature = consensus.CreateSignature(election, PrivateKey)
		transaction = election
		blockType = "Election"
	}

	block := consensus.GenerateBlock(PrevIndex, PrevHash, transaction, PublicKey)
	block.Type = blockType

	//Check if there is a proposed block currently, if so, add to the queue
	if reflect.DeepEqual(ProposedBlock, database.Block{}) {
		logger.Println("peer_routes.go", "HandleConn()", "Empty, proposing this block.")
		ProposedBlock = block
		ProposeBlock(ProposedBlock, Nodes)
	} else {
		logger.Println("peer_routes.go", "HandleConn()", "Not Empty, sending to queue.")
		BlockQueue = append(BlockQueue, block)
		fmt.Println(BlockQueue)
	}
}

//Receive the responses given by all other peers deciding if a block is valid
func ReceiveResponses(answer bool, sMap map[string]string) {
	/*
		Use answer and pair it with sMap which allows for accountability
		of their choices

		TODO: assumes that sMap is length 1, could be an issue????
	*/
	for k, v := range sMap {
		SignatureMap[k] = make(map[string]bool)
		SignatureMap[k][v] = answer
	}

	if len(SignatureMap) == len(Nodes) {
		logger.Println("peer_routes.go", "HandleConn()", "Received Responses!!")
		CheckResponses(len(SignatureMap)) //Go through the responses and see if block valid
		SignatureMap = nil
		SignatureMap = make(map[string]map[string]bool)
		ProposedBlock = database.Block{}
	} else {
		logger.Println("peer_routes.go", "HandleConn()", "Didn't Receiving Responses")
		return
	}

	if len(BlockQueue) > 0 {
		//Propose the next block
		ProposedBlock = BlockQueue[0]
		//TODO: get rid of first item in slice
		ProposeBlock(ProposedBlock, Nodes)
	} else {
		//Wait for the next vote
		logger.Println("peer_routes.go", "HandleConn()", "Everything is up to date.")
	}
}

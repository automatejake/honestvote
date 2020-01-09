package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
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

	var message Message
	data, err := json.Marshal(Self)
	if err != nil {
		logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	}
	message.Message = "connect response"
	message.Data = data
	data, err = json.Marshal(message)
	if err != nil {
		logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	}

	conn.Write(data)

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

	if mType == "Vote" {
		//Temporary Variable, will be data unmarshalled
		transaction = database.Vote{Sender: "0xcheese", Value: 1, Receiver: map[int]string{1: "0xsugar", 2: "0xpeanut"}}
	} else if mType == "Election" {
		//Temporary Variable, will be data unmarshalled
		transaction = database.Election{Name: "WCU", EligibleVoters: 16345, Start: "3/23/2020", End: "3/30/2020"}
	}

	// if !crypto.Verify(){
	//return
	// }

	block := consensus.GenerateBlock(PrevIndex, PrevHash, transaction, PublicKey)

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
	SignatureMap = make(map[bool]map[string]string)
	for k, v := range sMap {
		SignatureMap[answer] = make(map[string]string)
		SignatureMap[answer][k] = v
	}

	logger.Println("peer_routes.go", "HandleConn()", "Receiving Responses")

	if len(SignatureMap) == len(Nodes) {
		CheckResponses(len(SignatureMap)) //Go through the responses and see if block valid
		SignatureMap = nil
		ProposedBlock = database.Block{}
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

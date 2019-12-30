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
	//ADD TO DATABASE AS WELL
	// Nodes[port] = true
	tmpNode := database.TempNode{
		IPAddress: conn.RemoteAddr().String(),
		Socket:    conn,
	}
	node.IPAddress = conn.RemoteAddr().String()[0:9]
	if !database.DoesNodeExist(node) {
		database.AddNode(node)
	}

	var message Message
	data, err := json.Marshal(Self)
	if err != nil {
		logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	}
	message.Message = "get id"
	message.Data = data
	data, err = json.Marshal(message)
	if err != nil {
		logger.Println("read_functions.go", "AcceptConnectMessage()", err.Error())
	}

	conn.Write(data)

	Nodes = append(Nodes, tmpNode)

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
func ReceiveVote(vote int) {
	block := consensus.GenerateBlock(PrevIndex, PrevHash, database.Transaction{
		Sender:   "",
		Vote:     vote,
		Receiver: "",
	}, Port, PublicKey)

	//Check if there is a proposed block currently, if so, add to the queue
	if ProposedBlock == (database.Block{}) {
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
func ReceiveResponses(block *database.Block) {
	ValidatorResponses = append(ValidatorResponses, *block) //Keep track of all responses to check and compare
	logger.Println("peer_routes.go", "HandleConn()", "Receiving Responses")
	if len(ValidatorResponses) == len(Nodes) {
		CheckResponses(ValidatorResponses, len(ValidatorResponses)) //Go through the responses and see if block valid
		ValidatorResponses = nil
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

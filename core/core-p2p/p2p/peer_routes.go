package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	for {
		d := json.NewDecoder(conn)
		var write database.Write
		d.Decode(&write)

		if write.Message == "connect" {
			port, err := strconv.Atoi(string(write.Data))
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")
			if err == nil {
				ConnectMessage(port, conn)
			}
		} else if write.Message == "recieve data" {
			buffer := bytes.NewBuffer(write.Data)
			DecodeData(buffer)
		} else if write.Message == "get data" {
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		} else if write.Message == "vote" { //Get a vote and make a block out of it
			vote := write.Vote
			ReceiveVote(vote)
		} else if write.Message == "verify" { //Verifying that the sent block is correct(sign/reject)
			block := new(database.Block)
			json.Unmarshal(write.Data, block)
			logger.Println("peer_routes.go", "HandleConn()", "Verifying")
			VerifyBlock(*block)
		} else if write.Message == "sign" { //Response from all Nodes verifying block
			block := new(database.Block)
			err := json.Unmarshal(write.Data, &block)
			if err == nil {
				ReceiveResponses(block)
			} else {
				fmt.Println(err)
			}
		} else if write.Message == "update" {
			block := new(database.Block)
			json.Unmarshal(write.Data, block)
			if database.UpdateBlockchain(database.MongoDB, *block) {
				PrevHash = block.Hash
				PrevIndex = block.Index
				fmt.Println(string(PrevIndex) + " " + PrevHash)
			}
		}
	}
}

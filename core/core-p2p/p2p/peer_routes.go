package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		length, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {
			port, err := strconv.Atoi(string(buf[8:length]))
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")
			if err == nil {
				ConnectMessage(port, conn)
			}
		} else if string(buf[0:12]) == "recieve data" {
			buffer := bytes.NewBuffer(buf[13:length])
			DecodeData(buffer)
		} else if string(buf[0:8]) == "get data" {
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		} else if string(buf[0:4]) == "vote" { //Get a vote and make a block out of it
			sVote := string(buf[5:length])
			sVote = strings.TrimSuffix(sVote, "\n")
			vote, err := strconv.Atoi(sVote)
			if err == nil {
				ReceiveVote(vote)
			}
		} else if string(buf[0:6]) == "verify" { //Verifying that the sent block is correct(sign/reject)
			block := new(database.Block)
			json.Unmarshal(buf[7:length], block)
			logger.Println("peer_routes.go", "HandleConn()", "Verifying")
			VerifyBlock(*block)
		} else if string(buf[0:4]) == "sign" { //Response from all Nodes verifying block
			block := new(database.Block)
			json.Unmarshal(buf[5:length], block)
			ReceiveResponses(block)
		} else if string(buf[0:6]) == "update" {
			block := new(database.Block)
			json.Unmarshal(buf[7:length], block)
			if database.UpdateBlockchain(database.MongoDB, *block) {
				PrevHash = block.Hash
				PrevIndex = block.Index
				fmt.Println(string(PrevIndex) + " " + PrevHash)
			}
		}
	}
}

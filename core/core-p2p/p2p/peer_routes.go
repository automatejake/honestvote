package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
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

			//ADD TO DATABASE AS WELL

			port, err := strconv.Atoi(string(buf[8:length]))

			if err == nil {
				Nodes[port] = true
				tmpPeer := database.TempPeer{
					IPAddress: "127.0.0.1",
					Port:      port,
					Socket:    conn,
				}
				Peers = append(Peers, tmpPeer)
				fmt.Println(Peers)
				// permPeer := database.Peer{}
				// database.AddToTable(permPeer.IPAddress, permPeer.Port)
			}
		} else if string(buf[0:12]) == "recieve data" {
			buffer := bytes.NewBuffer(buf[13:length])
			tmpArray := new([]database.Candidate)
			js := json.NewDecoder(buffer)
			err := js.Decode(tmpArray)
			if err == nil {
				database.UpdateMongo(database.MongoDB, *tmpArray, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
			}
		} else if string(buf[0:8]) == "get data" {
			database.MoveDocuments(Peers, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		} else if string(buf[0:4]) == "vote" {
			sVote := string(buf[5:length])
			sVote = strings.TrimSuffix(sVote, "\n")
			vote, err := strconv.Atoi(sVote)
			if err == nil {
				block := consensus.GenerateBlock(database.Block{}, database.Transaction{
					Sender:   "",
					Vote:     vote,
					Receiver: "",
				})

				//Check if there is a proposed block currently, if so, add to the queue
				if ProposedBlock == (database.Block{}) {
					fmt.Println("Empty, proposing this block.")
					ProposedBlock = block
					ProposeBlock(ProposedBlock, Peers)
				} else {
					fmt.Println("Not Empty, sending to queue.")
					BlockQueue = append(BlockQueue, block)
					fmt.Println(BlockQueue)
				}
			}
		} else if string(buf[0:6]) == "verify" {
			//TODO: Verify the block is correct
			block := new(database.Block)
			json.Unmarshal(buf[7:length], block)
			if (consensus.VerifyHash(database.Block{}, *block)) {
				fmt.Println("Block verified!")
			}
		}
	}
}

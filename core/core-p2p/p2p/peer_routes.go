package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//decode json data
	d := json.NewDecoder(conn)

	for {

		var write Message
		err := d.Decode(&write)

		if err != nil {
			logger.Println("peer_routes.go", "HandleConn()", err.Error())
			return
		}

		switch write.Message {
		case "connect":
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")

			var node database.Node
			json.Unmarshal(write.Data, &node)

			AcceptConnectMessage(node, conn)
		case "send connected nodes":
			var node database.Node
			json.Unmarshal(write.Data, &node)
			tmp_peers := database.FindNodes()
			fmt.Println(tmp_peers)
			if tmp_peers != nil {
				peers_json, err := json.Marshal(tmp_peers)
				if err != nil {
					logger.Println("peer_routes.go", "RegisterNode", err.Error())
				}
				_, err = conn.Write(peers_json)
				if err != nil {
					logger.Println("peer_routes.go", "RegisterNode", err.Error())
				}
			}
		case "recieve data":
			buffer := bytes.NewBuffer(write.Data)
			DecodeData(buffer)
		case "get data":
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		case "transaction":
			ReceiveTransaction(write.Data, write.Type)
		case "register":
			tcp_port := strconv.Itoa(TCP_PORT)
			registration.EmailRegistration("jacob@neubaum.com (senders_email)", "election_name", "senders_public_key", PublicIP, tcp_port)
		case "become peer":
			var node database.Node
			json.Unmarshal(write.Data, &node)
			// administrator.ProposePeer(node)
		case "new election":
			//Create a new election
		case "verify":
			DecideType(write.Data, write.Type, conn)
		case "sign":
			answer, err := strconv.ParseBool(string(write.Data))

			if err == nil {
				ReceiveResponses(answer, write.Signature)
			}
		case "update":
			block := new(database.Block)
			json.Unmarshal(write.Data, block)
			if database.UpdateBlockchain(database.MongoDB, *block) {
				PrevHash = block.Hash
				PrevIndex = block.Index
				logger.Println("peer_routes.go", "HandleConn()", string(PrevIndex)+" "+PrevHash)
			}
		case "find":
		default:
			logger.Println("peer_routes.go", "HandleConn", "Recieved Bad Message")
			conn.Close()
			break
			// database.FindDocument(database.MongoDB, database.CollectionPrefix+"blockchain", database.Vote{Value: 1}, "Vote")
		}
	}
}

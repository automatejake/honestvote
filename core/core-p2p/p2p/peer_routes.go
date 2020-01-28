package p2p

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//decode json data
	d := json.NewDecoder(conn)

	for {

		var message Message
		err := d.Decode(&message)

		if err != nil {
			logger.Println("peer_routes.go", "HandleConn()", err.Error())
			return
		}
		fmt.Println(message)

		switch message.Message {
		case "connect":
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")

			var node database.Node
			json.Unmarshal(message.Data, &node)

			AcceptConnectMessage(node, conn)
		case "send connected nodes":
			var node database.Node
			json.Unmarshal(message.Data, &node)
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
			buffer := bytes.NewBuffer(message.Data)
			DecodeData(buffer)
		case "get data":
			database.MoveDocuments(Nodes, database.DatabaseName, database.CollectionPrefix+database.ElectionHistory)
		case "transaction":
			fmt.Println("recieved transaction")
			ReceiveTransaction(message.Type, message.Data)
		case "register":
			tcp_port := strconv.Itoa(TCP_PORT)
			registration.EmailRegistration("jacob@neubaum.com (senders_email)", "election_name", "senders_public_key", PublicIP, tcp_port)
		case "become peer":
			var node database.Node
			json.Unmarshal(message.Data, &node)
			// administrator.ProposePeer(node)
		case "verify transaction":
			var block database.Block
			err := json.Unmarshal(message.Data, &block)
			if err != nil {
			}
			if consensus.IsBlockValid(PreviousBlock, block) {
				err = database.AddBlock(block)
				if err != nil {
				}
				PreviousBlock = block
			}
		default:
			logger.Println("peer_routes.go", "HandleConn", "Recieved Bad Message")
			conn.Close()
			break
			// database.FindDocument(database.MongoDB, database.CollectionPrefix+"blockchain", database.Vote{Value: 1}, "Vote")
		}
	}
}

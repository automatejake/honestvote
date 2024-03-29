package p2p

import (
	"encoding/json"
	"net"

	"github.com/jneubaum/honestvote/core/core-validation/validation"

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

		switch message.Message {
		case "connect":
			logger.Println("peer_routes.go", "HandleConn()", "Recieved Connect Message")

			var node database.Node
			json.Unmarshal(message.Data, &node)

			AcceptConnectMessage(node, conn)
		case "send connected nodes":
			var node database.Node
			json.Unmarshal(message.Data, &node)
			if !database.DoesNodeExist(node) {
				node.Role = "full node"
				database.AddNode(node)
			}
			tmp_peers := database.FindNodes()

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
			// ReceiveTransaction(message.Type, message.Data)
		case "register":
			var registrant database.AwaitingRegistration
			err := json.Unmarshal(message.Data, &registrant)
			if err != nil {
				logger.Println("peer_routes.go", "HandleConn()", err)
			}
			if registration.IsValidEmailRegistrant(&registrant) {
				registration.SendRegistrationCode(registrant, Self.IPAddress, HTTP_Port, Email_Address, Email_Password)
			}

			// tcp_port := strconv.Itoa(TCP_PORT)
			// registration.EmailRegistration("jacob@neubaum.com (senders_email)", "election_name", "senders_public_key", PublicIP, tcp_port)
		case "become peer":
			var node database.Node
			json.Unmarshal(message.Data, &node)
		case "grab data":
			GrabDocuments(database.MongoDB, conn, string(message.Data))
		case "receive block":
			DecodeBlockData(message.Data)
		case "receive transaction":
			DecodeTransactionData(message.Data, message.Type)
		case "send transaction":
			switch message.Type {
			case "elections":
				var election database.Election
				err := json.Unmarshal(message.Data, &election)
				if err != nil {
					logger.Println("peer_routes.go", "HandleConn()", err)
				}

				valid, err := validation.IsValidElection(election)
				if valid && err == nil {
					database.AddTransaction(election, message.Type)
				}
			case "registrations":
				var registration database.Registration
				err := json.Unmarshal(message.Data, &registration)
				if err != nil {
					logger.Println("peer_routes.go", "HandleConn()", err)
				}

				valid, err := validation.IsValidRegistration(registration)
				if valid && err == nil {
					database.AddTransaction(registration, message.Type)
				}
			case "votes":
				var vote database.Vote
				err := json.Unmarshal(message.Data, &vote)
				if err != nil {
					logger.Println("peer_routes.go", "HandleConn()", err)
				}

				valid, err := validation.IsValidVote(vote)
				if valid && err == nil {
					database.AddTransaction(vote, message.Type)
				}
			}
		case "verify block":
			var block database.Block

			err := json.Unmarshal(message.Data, &block)
			if err != nil {
				logger.Println("peer_routes.go", "HandleConn()", err)
			}

			sigValid := consensus.CheckSignature(block)
			if !sigValid {
				logger.Println("peer_routes.go", "HandleConn()", "The signature is invalid and someone is impersonating the sender")
				return
			}

			verified, err := consensus.IsBlockValid(PreviousBlock, block)
			if verified {
				err = database.AddBlock(block)
				if err != nil {
					logger.Println("peer_routes.go", "HandleConn()", err)
				}
				PreviousBlock = block
			} else {
				// calls a fn to change the role of a dishonest node(hon-117)
				var node database.Node
				database.MarkDishonestNode(node)
				logger.Println("peer_routes.go", "HandleConn", err.Error())
			}
		default:
			logger.Println("peer_routes.go", "HandleConn", "Recieved Bad Message")
			conn.Close()
			break
			// database.FindDocument(database.MongoDB, database.CollectionPrefix+"blockchain", database.Vote{Value: 1}, "Vote")
		}
	}
}

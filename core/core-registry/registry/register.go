package registry

import (
	"encoding/json"
	"log"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

/**
* Register Node - 2 Step Process
*
* 1) Adds the node to the database of connections
* 2) Returns to node the list of nodes to speak with, IP Address and Port
*
**/

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr, tcp_port int) {

	// Adds the node to the database of connections as a full node.  Nodes do not become peers until accpetance by the network
	database.AddToTable(addr.IP.String(), tcp_port)

	// Returns to node the list of nodes to speak with, IP Address and Port
	tmp_peers := database.FindPeers()

	peers_json, err := json.Marshal(tmp_peers)
	if err != nil {
		log.Println("File: register.go\nFunction:RegisterNode\n", err)
	}

	_, err = conn.WriteToUDP(peers_json, addr)
	if err != nil {
		log.Println("File: register.go\nFunction:RegisterNode\n", err)
	}

}

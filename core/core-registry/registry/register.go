package registry

import (
	"encoding/json"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

/**
* Register Node - 2 Step Process
*
* 1) Checks if the node already exists in the database
* 2) If so, adds the node to the database of connections
* 3) Returns to node the list of nodes to speak with, IP Address and Port contained in a JSON object
*
**/

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr, tcp_port int) {

	//Checks if the node already exists in the database
	tempNode := database.Node{
		IPAddress: addr.IP.String(),
		Port:      tcp_port,
	}
	if !database.DoesNodeExist(tempNode) {
		// Adds the node to the database of connections as a full node.  Nodes do not become peers until accpetance by the network
		database.AddNode(tempNode)
	}

	tmp_peers := database.FindNodes()

	peers_json, err := json.Marshal(tmp_peers)
	if err != nil {
		logger.Println("register.go", "RegisterNode\n", err.Error())
	}

	_, err = conn.WriteToUDP(peers_json, addr)
	if err != nil {
		logger.Println("register.go", "RegisterNode\n", err.Error())
	}

}

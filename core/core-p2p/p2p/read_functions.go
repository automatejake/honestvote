package p2p

import (
	"encoding/json"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Adds new connection to database and local Node array
func AcceptConnectMessage(node database.Node, conn net.Conn) {

	node.IPAddress = conn.RemoteAddr().String()
	if !database.DoesNodeExist(node) {
		database.AddNode(node)
	}

	Nodes = append(Nodes, conn)

}

//Decoding the data sent from another peer, this data is from a database
func DecodeData(data []byte) {
	var block database.Block

	err := json.Unmarshal(data, &block)
	if err != nil {
		logger.Println("read_functions.go", "DecodeData()", err)
		return
	}

	PreviousBlock = block
	database.UpdateMongo(database.MongoDB, block)

}

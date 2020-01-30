package discovery

import (
	"encoding/json"
	"net"
	"strconv"
	"sync"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Used to update the blockchain only once per producer
var doOnce sync.Once

/***
* Find Nodes to talk to in the network, 2 Step Process
*
* 1) Send findnode message to registry node over raw udp socket (include TCP socket that you will be listening in on)
* 2) Parse node from JSON to struct
* 3) Send Connect Message to node
*
**/
func FetchLatestPeers(registry_ip string, registry_port string, tcp_port string) {
	// logger.Println("find_peer.go", "FetchLatestPeers()", "Contacting: "+registry_ip+":"+registry_port)

	conn, err := net.Dial("tcp", registry_ip+":"+registry_port)
	if err != nil {
		logger.Println("find_peers.go", "FetchLatestPeers", err.Error())
	}
	if conn != nil {
		logger.Println("find_peers.go", "FetchLatestPeers", "Dial Successful!")

		byteSelf, err := json.Marshal(p2p.Self)
		if err != nil {
			logger.Println("find_peer.go", "FetchLatestPeers", err.Error())
		}

		var write p2p.Message
		write.Message = "send connected nodes"
		write.Data = byteSelf

		byteSelf, err = json.Marshal(write)
		conn.Write(byteSelf)

		peers_json := make([]byte, 2048)
		n, err := conn.Read(peers_json) // n, udp_address, error

		var peers []database.Node
		err = json.Unmarshal(peers_json[0:n], &peers)

		for _, peer := range peers {
			if !database.DoesNodeExist(peer) {
				database.AddNode(peer)
			}
			ConnectMessage(peer)
		}
	}

}

// /*
// * 1) Attempt to connect to peer
// * 2) If unsuccessful, report to registry node
// * 3) If succsessful, Add Peer to database and connection to memory
//  */
func ConnectMessage(peer database.Node) { //is run
	port := strconv.Itoa(peer.Port)

	conn, err := net.Dial("tcp", peer.IPAddress+":"+port) //dials up the other two nodes
	if err != nil {
		logger.Println("find_peer.go", "ConnectMessage", err.Error())
	}
	if conn != nil { //you dont want to run this unless at least one connection is made

		//Catch up on latest blockchain and only run it once
		doOnce.Do(func() { p2p.SendIndex(database.LastIndex(database.MongoDB), conn) })
		//when you have new cons you have multithreading
		//
		doOnce.Do(func() { p2p.LatestHashAndIndex(database.MongoDB) })
		//

		logger.Println("find_peer.go", "ConnectMessage", "Dial Successful!")

		write := new(p2p.Message)

		byteSelf, err := json.Marshal(p2p.Self)
		if err != nil {
			logger.Println("find_peer.go", "ConnectMessage", err.Error())
		}

		write.Message = "connect"
		write.Data = byteSelf
		jWrite, err := json.Marshal(write)
		if err != nil {
			logger.Println("find_peer.go", "ConnectMessage()", err.Error())
		}
		conn.Write(jWrite)

		p2p.Nodes = append(p2p.Nodes, conn)

		go p2p.HandleConn(conn)
	}
}

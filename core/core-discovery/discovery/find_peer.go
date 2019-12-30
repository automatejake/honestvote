package discovery

import (
	"encoding/json"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

/***
* Find Nodes to talk to in the network, 2 Step Process
*
* 1) Send findnode message to registry node over raw udp socket (include TCP socket that you will be listening in on)
* 2) Parse node from JSON to struct
* 3) Send Connect Message to node
*
**/
func FindPeer(registry_ip string, registry_port string, tcp_port string) {
	logger.Println("find_peer.go", "FindPeer()", "Contacting: "+registry_ip+":"+registry_port)

	// Send findpeer message to registry node over raw udp socket (include TCP socket that you will be listening in on)
	remote_address, err := net.ResolveUDPAddr("udp", registry_ip+":"+registry_port)
	if err != nil {
		logger.Println("find_peer.go", "FindPeer", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, remote_address)
	if err != nil {
		logger.Println("find_peer.go", "FindPeer", err.Error())
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("findpeer" + tcp_port))
	if err != nil {
		logger.Println("find_peer.go", "FindPeer", err.Error())
	}

	// Parse Peer from JSON to struct
	peers_json := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(peers_json) // n, udp_address, error

	var peers []database.Node
	_ = json.Unmarshal(peers_json[0:n], &peers)

	//Send connect message to peer
	for _, peer := range peers {
		ConnectMessage(peer, tcp_port)
	}

}

/*
* 1) Attempt to connect to peer
* 2) If unsuccessful, report to registry node
* 3) If succsessful, Add Peer to database and connection to memory
 */
func ConnectMessage(peer database.Node, tcp_port string) {
	port := strconv.Itoa(peer.Port)

	conn, err := net.Dial("tcp", peer.IPAddress+":"+port)
	if err != nil {
		logger.Println("find_peer.go", "ConnectMessage", err.Error())
	}
	if conn != nil {

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

		tmpNode := database.TempNode{
			IPAddress: peer.IPAddress,
			Port:      peer.Port,
			Socket:    conn,
		}
		p2p.Nodes = append(p2p.Nodes, tmpNode)
		if !database.DoesNodeExist(database.Node{
			IPAddress: peer.IPAddress,
			Port:      peer.Port,
		}) {
			database.AddNode(peer)
		}

		go p2p.HandleConn(conn)
	}
}

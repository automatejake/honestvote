package discovery

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

/***
* Find Peers to talk to in the network, 2 Step Process
*
* 1) Send findpeer message to registry node over raw udp socket (include TCP socket that you will be listening in on)
* 2) Parse peer from JSON to struct
* 3) Send Connect Message to Peer
*
**/
func FindPeer(registry_ip string, registry_port string, tcp_port string) {

	// Send findpeer message to registry node over raw udp socket (include TCP socket that you will be listening in on)
	remote_address, err := net.ResolveUDPAddr("udp", registry_ip+":"+registry_port)
	if err != nil {
		log.Println("File: find_peer.go\nFunction:FindPeer\n", err)
	}

	conn, err := net.DialUDP("udp", nil, remote_address)
	if err != nil {
		log.Println("File: find_peer.go\nFunction:FindPeer\n", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("findpeer" + tcp_port))
	if err != nil {
		log.Println("File: find_peer.go\nFunction:FindPeer\n", err)
	}

	// Parse Peer from JSON to struct
	peers_json := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(peers_json) // n, udp_address, error

	var peers []database.Peer
	_ = json.Unmarshal(peers_json[0:n], &peers)

	//Send connect message to peer
	for _, peer := range peers {
		ConnectMessage(peer)
	}

}

/*
* 1) Attempt to connect to peer
* 2) If unsuccessful, report to registry node
* 3) If succsessful, Add Peer to database and connection to memory
 */
func ConnectMessage(peer database.Peer) {
	port := strconv.Itoa(peer.Port)

	conn, err := net.Dial("tcp", peer.IPAddress+":"+port)
	if err != nil {
		log.Println("File: find_peer.go\nFunction:ConnectMessage\n", err)
	}

	if conn != nil {
		fmt.Println("Dial Successful!")

		conn.Write([]byte("connect " + port))

		tmpPeer := database.TempPeer{
			IPAddress: "127.0.0.1",
			Port:      peer.Port,
			Socket:    conn,
		}
		p2p.Peers = append(p2p.Peers, tmpPeer)
		database.AddToTable(peer.IPAddress, peer.Port)
		go p2p.HandleConn(conn)
	}

}

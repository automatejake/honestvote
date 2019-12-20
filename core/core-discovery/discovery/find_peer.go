package discovery

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

/***
* Find Peers to talk to in the network, 2 Step Process
*
* 1) Send findpeer message to registry node over raw udp socket (include TCP socket that you will be listening in on)
* 2)
*
**/

func FindPeer(registry_ip string, registry_port string, tcp_port string) {

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

	// fmt.Fprintf(conn, "findpeer"+tcp_port)
	_, err = conn.Write([]byte("findpeer" + tcp_port))
	if err != nil {
		log.Println("File: find_peer.go\nFunction:FindPeer\n", err)
	}

	peers_json := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(peers_json) // n, udp_address, error
	// fmt.Println("recieved ", n, " bytes")

	var peers []database.Peer
	_ = json.Unmarshal(peers_json[0:n], &peers)

	// fmt.Println(peers[0].IPAddress)
	log.Println("Got him " + peers[0].IPAddress)

	// _, err = bufio.NewReader(conn).Read(new_peer)
	// if err == nil {
	// 	go DialPeer(string(new_peer))
	// } else {
	// 	fmt.Print("Some error %v\n", err)
	// }

}

func DialPeer(peer string) {
	p := strings.Trim(peer, "\x00")
	fmt.Printf("%q\n", p)
	log.Printf("Here Before")
	conn, err := net.Dial("tcp", "127.0.0.1:"+p)
	log.Printf("Here After")
	if err != nil {
		log.Print(err)
	}

	port, _ := strconv.Atoi(peer)

	if conn != nil {
		fmt.Println("Dial Successful!")
		tmpPeer := database.TempPeer{
			IPAddress: "127.0.0.1",
			Port:      port,
			Socket:    conn,
		}
		p2p.Peers = append(p2p.Peers, tmpPeer)

		conn.Write([]byte("connect " + peer))
		go p2p.HandleConn(conn)
	}

}

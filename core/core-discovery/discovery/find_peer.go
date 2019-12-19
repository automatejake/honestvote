package discovery

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

/***
* Find Peers in the network
**/

func FindPeer(registry_ip string, registry_port string, tcp_port string) {

	new_peer := make([]byte, 2048)

	// Dial Connection
	conn, err := net.Dial("udp", registry_ip+":"+registry_port)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	// Read Connection
	//BREAKS WINDOWS
	//fmt.Fprintf(conn, "findpeer"+tcp_port)
	log.Printf("Here Before")
	_, err = bufio.NewReader(conn).Read(new_peer)
	log.Printf("Here After")
	if err == nil {
		fmt.Printf("%s\n", new_peer)
		DialPeer(string(new_peer))

	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()

}

func DialPeer(peer string) {
	fmt.Println(peer)
	conn, err := net.Dial("tcp", "127.0.0.1:"+peer)
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

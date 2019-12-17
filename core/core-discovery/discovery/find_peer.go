package discovery

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

/***
* Find Peers in the network
**/

func FindPeer(args string) {

	// buffer := make([]byte, 2048)

	// // Dial Connection
	// conn, err := net.Dial("udp", "127.0.0.1:7700")
	// if err != nil {
	// 	fmt.Printf("Some error %v", err)
	// 	return
	// }

	// // Read Connection
	// fmt.Fprintf(conn, "hello")
	// _, err = bufio.NewReader(conn).Read(buffer)
	// if err == nil {
	// 	fmt.Printf("%s\n", buffer)
	// } else {
	// 	fmt.Printf("Some error %v\n", err)
	// }
	// conn.Close()

	ignore, _ := strconv.Atoi(args)
	p2p.Nodes[ignore] = true

	for {
		for port := 7000; port <= 7001; port++ {
			if !p2p.Nodes[port] {
				//fmt.Println("Checking...")
				sPort := strconv.Itoa(port)
				conn, _ := net.Dial("tcp", "127.0.0.1:"+sPort)
				if conn != nil {
					fmt.Println("Dial Successful!")
					tmpPeer := database.Peer{
						IPAddress: "127.0.0.1",
						Port:      port,
						Socket:    conn,
					}
					p2p.Peers = append(p2p.Peers, tmpPeer)
					p2p.Nodes[port] = true

					conn.Write([]byte("connect " + strconv.Itoa(ignore)))
					go p2p.HandleConn(conn)
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

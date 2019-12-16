package discovery

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/joho/godotenv"
)

/***
* Find Peers in the network
**/

func FindPeer() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	ignore, _ := strconv.Atoi(os.Getenv("PORT"))
	p2p.Nodes[ignore] = true

	// conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 100, Zone: ""})
	// defer conn.Close()
	// conn.Write([]byte("hello"))
	// conn.Read(make([]byte, 1024))

	for {
		for port := 7000; port <= 7001; port++ {
			if !p2p.Nodes[port] {
				// fmt.Println("Checking...")
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

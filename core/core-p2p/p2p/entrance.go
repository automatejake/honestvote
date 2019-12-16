package p2p

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/joho/godotenv"
)

var nodes = make(map[int]bool)
var Peers []database.Peer

func PeerToPeer(args string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	ignore, _ := strconv.Atoi(args)
	nodes[ignore] = true

	for {
		for port := 7000; port <= 7001; port++ {
			if !nodes[port] {
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
					Peers = append(Peers, tmpPeer)
					nodes[port] = true

					conn.Write([]byte("connect " + strconv.Itoa(ignore)))
					go HandleConn(conn)
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

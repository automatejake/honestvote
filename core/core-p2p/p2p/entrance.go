package p2p

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/joho/godotenv"
)

var nodes = make(map[int]bool)
var Peers []database.Peer

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	ignore, _ := strconv.Atoi(os.Getenv("PORT"))
	nodes[ignore] = true

	go ListenConn()
	go http.CreateServer()

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

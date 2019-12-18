package discovery

import (
	"bufio"
	"fmt"
	"net"
)

/***
* Find Peers in the network
**/

func FindPeer(registry_ip string, registry_port string) {

	new_peer := make([]byte, 2048)

	// Dial Connection
	conn, err := net.Dial("udp", registry_ip+registry_port)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	// Read Connection
	fmt.Fprintf(conn, "hello")
	_, err = bufio.NewReader(conn).Read(new_peer)
	if err == nil {
		fmt.Printf("%s\n", new_peer)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()

	// ignore, _ := strconv.Atoi(IPAddress)
	// p2p.Nodes[ignore] = true

	// for {
	// 	for port := 7000; port <= 7001; port++ {
	// 		if !p2p.Nodes[port] {
	// 			//fmt.Println("Checking...")
	// 			sPort := strconv.Itoa(port)
	// 			conn, _ := net.Dial("tcp", "127.0.0.1:"+sPort)
	// 			if conn != nil {
	// 				fmt.Println("Dial Successful!")
	// 				tmpPeer := database.Peer{
	// 					IPAddress: "127.0.0.1",
	// 					Port:      port,
	// 					Socket:    conn,
	// 				}
	// 				p2p.Peers = append(p2p.Peers, tmpPeer)
	// 				p2p.Nodes[port] = true

	// 				// conn.Write([]byte("connect " + strconv.Itoa(ignore)))
	// 				go p2p.HandleConn(conn)
	// 			}
	// 		}
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
}

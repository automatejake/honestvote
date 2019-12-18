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
	conn, err := net.Dial("udp", registry_ip+":"+registry_port)
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

}

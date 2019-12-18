package registry

import (
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr, tcp_port string) {
	// defer conn.Close()
	fmt.Println(tcp_port)

	database.AddToTable(addr.IP.String(), tcp_port)
	tmp_peers := database.FindPeer()

	// var network bytes.Buffer // Stand-in for a network connection
	// enc := gob.NewEncoder(&network)

	for _, elem := range tmp_peers {

		// err := enc.Encode(elem)
		// if err != nil {
		// 	log.Fatal("encode error:", err)
		// }

		_, err := conn.WriteToUDP([]byte(elem.IPAddress+strconv.Itoa(elem.Port)), addr)
		if err != nil {
			fmt.Printf("Couldn't send response %v", err)
		}

	}

}

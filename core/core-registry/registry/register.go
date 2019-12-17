package registry

import (
	"fmt"
	"net"

	"github.com/jneubaum/honestvote-registry/core/core-database/database"
)

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr) {
	database.ExistsInTable(addr.IP.String())

	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

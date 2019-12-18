package registry

import (
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr, collection_prefix string) {
	port := strconv.Itoa(addr.Port)

	database.ExistsInTable(addr.IP.String(), port, collection_prefix)

	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

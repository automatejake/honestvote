package registry

import (
	"fmt"
	"net"
)

func RegisterNode(conn *net.UDPConn, addr *net.UDPAddr) {
	// port := strconv.Itoa(addr.Port)

	// database.ExistsInTable(addr.IP.String(), port)

	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

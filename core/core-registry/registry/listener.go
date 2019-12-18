package registry

import (
	"fmt"
	"net"
	"strconv"
)

func ListenConnections(port string, collection_prefix string) {

	port_integer, _ := strconv.Atoi(port)
	buffer := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port_integer,
		IP:   net.ParseIP("127.0.0.1"),
	}

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	for {
		n, remoteaddr, err := ser.ReadFromUDP(buffer)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, buffer)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		fmt.Println(string(buffer[0:n]))

		if string(buffer[0:n]) == "hello" {
			go RegisterNode(ser, remoteaddr, collection_prefix)
		}

	}
}

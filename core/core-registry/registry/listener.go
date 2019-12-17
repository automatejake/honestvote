package registry

import (
	"fmt"
	"net"
)

func ListenConnections() {

	buffer := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 7700,
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
			go RegisterNode(ser, remoteaddr)
		}

	}
}

package registry

import (
	"fmt"
	"net"
	"strconv"
)

func ListenConnections(udp_service string) {

	port, err := strconv.Atoi(udp_service)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	buffer := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	fmt.Println("Listening UDP on port ", addr.Port)
	for {

		n, remoteaddr, err := ser.ReadFromUDP(buffer) // n is length of bytes, remoteaddr is ip and port of message sender
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

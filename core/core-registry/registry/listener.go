package registry

import (
	"log"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
)

func ListenConnections(udp_port string) {

	// Sets up server to accept incoming connections
	port, err := strconv.Atoi(udp_port)
	if err != nil {
		logger.Println("listener.go", "ListenConnections", err.Error())
	}

	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		logger.Println("listener.go", "ListenConnections", err.Error())
		return
	}

	logger.Println("listener.go", "ListenConnections()", "Registry port running on port: "+udp_port)

	// The only UDP Route is called findpeer and is sent in order to tell the registry port that they want to talk to someone
	//
	defer listener.Close()
	buffer := make([]byte, 4096)

	for {

		// n is length of bytes, remoteaddr is ip and port of message sender
		n, remote_address, err := listener.ReadFromUDP(buffer)
		if err != nil {
			log.Println("File: listener.go\nFunction:ListenConnections2\n", err)
			continue
		}

		log.Println("Registry receiving message from node: ", buffer[0:n])

		if string(buffer[0:8]) == "findpeer" {
			// default tcp port is 7632, otherwise it should be specified explicitly
			non_default_port, _ := strconv.Atoi(string(buffer[8:n]))
			go RegisterNode(listener, remote_address, non_default_port)
		}

	}
}

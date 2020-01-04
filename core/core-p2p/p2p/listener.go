package p2p

import (
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
)

type Message struct {
	Message   string            `json:"message"`
	Data      []byte            `json:"data"`
	Signature map[string]string `json:"signature"`
	Vote      int               `json:"vote"` //Used to send vote, should be changed
}

func ListenConn(port string, role string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Println("listener.go", "ListenConn()", err.Error())
	}

	logger.Println("listener.go", "ListenConn()", "Peer running on port: "+port)

	TCP_SERVICE, err = strconv.Atoi(port)

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Println("listener.go", "ListenConn", err.Error())
		}

		// defined in peer_routes.go
		go HandleConn(conn)
	}
}

package p2p

import (
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
)

func ListenConn(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Println("listener.go", "ListenConn()", err.Error())
	}

	logger.Println("listener.go", "ListenConn()", "Peer running on port: "+port)

	Port, err = strconv.Atoi(port)

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

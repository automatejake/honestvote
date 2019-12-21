package p2p

import (
	"log"
	"net"
	"strconv"

	logger "github.com/jneubaum/honestvote/tests/logging"
)

func ListenConn(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Println("listener.go", "ListenConn()", err.Error())
	}

	logger.Println("listener.go", "ListenConn()", port)

	Port, err = strconv.Atoi(port)

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// defined in peer_routes.go
		go HandleConn(conn)
	}
}

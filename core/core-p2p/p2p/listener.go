package p2p

import (
	"log"
	"net"
	"strconv"
)

func ListenConn(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

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

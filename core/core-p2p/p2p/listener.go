package p2p

import (
	"log"
	"net"
)

func ListenConn(args string) {
	portString := args
	listen, err := net.Listen("tcp", portString)

	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}

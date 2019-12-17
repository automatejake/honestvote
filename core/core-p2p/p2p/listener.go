package p2p

import (
	"log"
	"net"
)

func ListenConn(port string) {
	listen, err := net.Listen("tcp", port)
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

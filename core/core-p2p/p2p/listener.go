package p2p

import (
	"log"
	"net"
	"os"
)

func ListenConn() {
	portString := ":" + os.Getenv("PORT")
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

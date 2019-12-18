package p2p

import (
	"log"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

var Nodes = make(map[int]bool)
var Peers []database.Peer

func ListenConn(port string, collection_prefix string) {
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

		// defined in peer_routes.go
		go HandleConn(conn, collection_prefix)
	}
}

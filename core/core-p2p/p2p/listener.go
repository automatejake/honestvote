package p2p

import (
	"log"
	"net"
)

<<<<<<< HEAD
=======
var Nodes = make(map[int]bool)
var Peers []database.TempPeer
var ProposedBlock database.Block

//Block queue starts to fill if ProposedBlock is not nil
var BlockQueue []database.Block

>>>>>>> 3961aefaabcfbd23a5b413a3d349dd0f9f6efdd5
func ListenConn(port string) {
	listen, err := net.Listen("tcp", ":"+port)
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
		go HandleConn(conn)
	}
}

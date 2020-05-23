package p2p

import (
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"

	"github.com/jneubaum/honestvote/tests/logger"
)

func ListenConn(port string, role string) error {
	PreviousBlock = LatestHashAndIndex(database.MongoDB)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Println("listener.go", "ListenConn()", err.Error())
	}

	logger.Println("listener.go", "ListenConn()", "p2p service running on port: "+port)

	defer listen.Close()

	go BroadcastScheduler()
	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Println("listener.go", "ListenConn", err.Error())
		}

		// defined in peer_routes.go
		go HandleConn(conn)
	}
}

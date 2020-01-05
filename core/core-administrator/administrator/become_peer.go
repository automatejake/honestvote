package administrator

import (
	"encoding/json"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

func RequestPeerPrivileges(peer database.Node) {
	port := strconv.Itoa(peer.Port)

	conn, err := net.Dial("tcp", peer.IPAddress+":"+port)
	if err != nil {
		logger.Println("become_peer.go", "BecomePeer()", err.Error())
	}
	if conn != nil {

		logger.Println("find_peer.go", "BecomePeer()", "Dial Successful!")

		write := new(p2p.Message)

		byteSelf, err := json.Marshal(p2p.Self)
		if err != nil {
			logger.Println("find_peer.go", "BecomePeer()", err.Error())
		}

		write.Message = "become peer"
		write.Data = byteSelf
		jWrite, err := json.Marshal(write)
		if err != nil {
			logger.Println("find_peer.go", "BecomePeer()()", err.Error())
		}
		conn.Write(jWrite)

		p2p.Nodes = append(p2p.Nodes, conn)

		go p2p.HandleConn(conn)
	}
}

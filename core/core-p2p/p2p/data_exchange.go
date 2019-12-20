package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func ProposeBlock(block database.Block, peers []database.TempPeer) {
	j, err := json.Marshal(block)

	if err == nil {
		fmt.Println(len(peers))
		for _, peer := range peers {
			peer.Socket.Write(append([]byte("verify "), j...))
		}
	}
}

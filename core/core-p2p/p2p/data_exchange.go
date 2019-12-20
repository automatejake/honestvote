package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
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

func VerifyBlock(block database.Block) {
	if (consensus.VerifyHash(database.Block{}, block)) {
		for _, peer := range Peers {
			if peer.Port == block.Port {
				peer.Socket.Write([]byte("sign"))
			}
		}
	}
}

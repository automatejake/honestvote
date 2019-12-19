package p2p

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func ProposeBlock(block database.Block, peers []database.TempPeer) {
	buffer := new(bytes.Buffer)
	tmpStruct := block
	gobobj := gob.NewEncoder(buffer)
	err := gobobj.Encode(tmpStruct)
	if err == nil {
		fmt.Println(len(peers))
		for _, peer := range peers {
			peer.Socket.Write(append([]byte("propose "), buffer.Bytes()...))
		}
	}
}

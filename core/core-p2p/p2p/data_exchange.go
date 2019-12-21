package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func ProposeBlock(block database.Block, peers []database.TempNode) {
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
		block.Signiture = "" //Put the Validator's signiture here so peer knows who signed it
		block.Valid = true
	} else {
		block.Signiture = "" //Put the Validator's signiture here so peer knows who signed it
		block.Valid = false
	}

	j, err := json.Marshal(block)

	if err == nil {
		for _, node := range Nodes {
			if node.Port == block.Port {
				node.Socket.Write(append([]byte("sign "), j...))
			}
		}
	}
}

func CheckResponses(responses []database.Block, size int) {
	counter := size
	for _, response := range responses{
		if response.Valid{
			continue
		}else{
			counter--
		}
	}

	if size == counter{
		database.UpdateBlockchain(database.MongoDB, ProposedBlock) //Update the mongo database with the new block
	}else{
		fmt.Println("Someone is a bad actor or this block is wrong.")
	}
}


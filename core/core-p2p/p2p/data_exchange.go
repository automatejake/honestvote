package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Send a block out to be verified by other peers
func ProposeBlock(block database.Block, peers []database.TempNode) {
	j, err := json.Marshal(block)

	if err == nil {
		fmt.Println(len(peers))
		for _, peer := range peers {
			peer.Socket.Write(append([]byte("verify "), j...))
		}
	}
}

//Decide if the block sent is valid
func VerifyBlock(block database.Block) {
	if consensus.VerifyHash(PrevIndex, PrevHash, block) {
		block.Signiture = "" //Put the Validator's signiture here so peer knows who signed it
		block.Valid = true
	} else {
		block.Signiture = "" //Put the Validator's signiture here so peer knows who signed it
		block.Valid = false
	}

	j, err := json.Marshal(block)

	if err == nil {
		logger.Println("peer_routes.go", "HandleConn()", "Sending response")
		for _, node := range Nodes {
			if node.Port == block.Port {
				node.Socket.Write(append([]byte("sign "), j...))
			}
		}
	}
}

//Go through all responses from other peers and see the result
func CheckResponses(responses []database.Block, size int) {
	counter := size
	for _, response := range responses {
		if response.Valid {
			continue
		} else {
			counter--
		}
	}

	if size == counter {
		j, err := json.Marshal(ProposedBlock)

		if err == nil {
			if database.UpdateBlockchain(database.MongoDB, ProposedBlock) {
				PrevHash = ProposedBlock.Hash
				PrevIndex = ProposedBlock.Index
				fmt.Println(string(PrevIndex) + " " + PrevHash)
			}

			for _, node := range Nodes {
				node.Socket.Write(append([]byte("update "), j...))
			}
		}
	} else {
		logger.Println("data_exchange.go", "CheckResponses", "Someone is a bad actor or this block is wrong.")
	}
}

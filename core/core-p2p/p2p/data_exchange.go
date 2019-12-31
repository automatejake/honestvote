package p2p

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Send a block out to be verified by other peers
func ProposeBlock(block database.Block, peers []net.Conn) {
	j, err := json.Marshal(block)

	write := new(Message)
	write.Message = "verify"
	write.Data = j

	jWrite, err := json.Marshal(write)

	if err == nil {
		fmt.Println(len(peers))
		for _, peer := range peers {
			peer.Write(jWrite)
		}
	}
}

//Decide if the block sent is valid
func VerifyBlock(block database.Block, conn net.Conn) {
	// if consensus.VerifyHash(PrevIndex, PrevHash, block) {
	// 	Valid = true
	// } else {
	// 	Valid = false
	// }

	j, err := json.Marshal(block)

	write := new(Message)
	write.Message = "sign"
	write.Data = j
	write.Signature = make(map[string]string)
	write.Signature[PublicKey], err = crypto.Sign(j, PrivateKey)

	if err != nil {
		write.Signature[PublicKey] = "Signature"
	}

	jWrite, err := json.Marshal(write)

	if err == nil {
		logger.Println("peer_routes.go", "HandleConn()", "Sending response")
		conn.Write(jWrite)
	}
}

//Go through all responses from other peers and see the result
func CheckResponses(size int) {
	counter := size

	checkBlock, err := json.Marshal(ProposedBlock)

	if err != nil {
		return
	}

	for k, v := range SignatureMap {
		valid, err := crypto.Verify(checkBlock, k, v)
		if valid && err == nil {
			ProposedBlock.Signatures[k] = v
		}
	}

	// for _, response := range responses {
	// 	if response.Valid {
	// 		if response.Valid {
	// 			continue
	// 		} else {
	// 			counter--
	// 		}
	// 	}
	// }

	if size == counter {
		j, err := json.Marshal(ProposedBlock)

		write := new(Message)
		write.Message = "update"
		write.Data = j

		jWrite, err := json.Marshal(write)

		if err == nil {
			if database.UpdateBlockchain(database.MongoDB, ProposedBlock) {
				PrevHash = ProposedBlock.Hash
				PrevIndex = ProposedBlock.Index
				fmt.Println(string(PrevIndex) + " " + PrevHash)
			}

			for _, node := range Nodes {
				node.Write(jWrite)
			}
		}
	} else {
		logger.Println("data_exchange.go", "CheckResponses", "Someone is a bad actor or this block is wrong.")
	}
}

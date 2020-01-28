package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

//Send a block out to be verified by other peers
func ProposeBlock(block database.Block) {
	j, err := json.Marshal(block)

	fmt.Println("proposed block")
	write := new(Message)
	write.Message = "verify transaction"
	write.Data = j
	write.Type = TransactionType(block.Transaction)

	jWrite, err := json.Marshal(write)

	if err == nil {
		for _, node := range Nodes {
			node.Write(jWrite)
		}
	}
	ProposedBlock = database.Block{}
}

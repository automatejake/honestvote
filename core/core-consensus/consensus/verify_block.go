package consensus

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func VerifyHash(prevIndex int, prevHash string, block database.Block) bool {
	if prevHash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	}
	// else if CalculateHash(GenerateHeader(block)) != block.Hash {
	// 	fmt.Println("Block hash is wrong!", CalculateHash(GenerateHeader(block)))
	// 	return false
	// }

	return true
}

func IsBlockValid(prevBlock database.Block, block database.Block) bool {
	return true
}

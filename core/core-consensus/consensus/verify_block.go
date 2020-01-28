package consensus

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsBlockValid(prevBlock database.Block, block database.Block) bool {
	if prevBlock.Hash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	}
	// else if CalculateHash(GenerateHeader(block)) != block.Hash {
	// 	fmt.Println("Block hash is wrong!", CalculateHash(GenerateHeader(block)))
	// 	return false
	// }

	//add code to verify each transaction

	return true
}

func VerifySignature(block database.Block) bool {
	//VerifySignature for making sure the sender is who they say they are

	return false
}

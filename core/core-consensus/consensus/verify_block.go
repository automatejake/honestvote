package consensus

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsBlockValid(prevBlock database.Block, block database.Block) (bool, error) {

	// Make sure that block's index is correct
	if prevBlock.Index != block.Index {

	}

	// Make sure that block's previous hash is the last block
	if prevBlock.Hash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false, nil
	}

	// Make sure the validator is a valid producer
	validator, err := database.FindNode(block.Validator)
	if err != nil {
		return false, err
	} else if validator.Role != "producer" {
		return false, nil
	}

	// Iterate through transactions contained in block and make sure that they are valid

	// Make sure that the merkle root is correct

	// Make sure that the block signature is correct
	header := GenerateBlockHeader(block)
	valid, err := crypto.Verify([]byte(header), database.PublicKey(block.Validator), block.Signature)
	if err != nil {
		return false, err
	} else if !valid {
		return false, err
	}

	return true, nil
}

func VerifySignature(block database.Block) bool {
	//VerifySignature for making sure the sender is who they say they are

	return false
}

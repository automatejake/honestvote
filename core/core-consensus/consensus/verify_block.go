package consensus

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-validation/validation"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsBlockValid(prevBlock database.Block, block database.Block) (bool, error) {
	customErr := &ConsensusError{
		Time: time.Now(),
	}
	ending := "  Block is invalid."

	// Make sure that block's index is correct
	if prevBlock.Index != block.Index {

	}

	// Make sure that block's previous hash is the last block
	if prevBlock.Hash != block.PrevHash {
		customErr.Message = "Previous hash is wrong!" + ending
		return false, customErr
	}

	// Make sure the validator is a valid producer
	validator, err := database.FindNode(block.Validator)
	if err != nil {
		return false, err
	} else if validator.Role != "producer" {
		customErr.Message = "Actor proposing this block is not a valid producer."
		return false, customErr
	}

	// Iterate through transactions contained in block and make sure that they are valid
	var honestTransaction bool
	switch database.TransactionType(block.Transaction) {
	case "Election":
		honestTransaction, err = validation.IsValidElection(block.Transaction.(database.Election))
	case "Registration":
		honestTransaction, err = validation.IsValidElection(block.Transaction.(database.Election))
	case "Vote":
		honestTransaction, err = validation.IsValidElection(block.Transaction.(database.Election))
	}
	if !honestTransaction {
		return false, err
	}

	// Make sure that the merkle root is correct
	if CalculateMerkleRoot(block) != block.MerkleRoot {
		customErr.Message = "Merkle root is incorrect."
		return false, customErr
	}

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

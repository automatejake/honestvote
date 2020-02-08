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
	ending := ", invalid block rejected."

	// Make sure that block's index is correct
	if prevBlock.Index != block.Index {
		customErr.Message = "Block index is incorrect" + ending
		return false, customErr
	}

	// Make sure that block's previous hash is the last block
	if prevBlock.Hash != block.PrevHash {
		customErr.Message = "Block's previous hash is incorrect" + ending
		return false, customErr
	}

	// Make sure the validator is a valid producer
	validator, err := database.FindNode(block.Validator)
	if err != nil {
		return false, err
	} else if validator.Role != "producer" {
		customErr.Message = "Actor proposing this block is not a valid producer." + ending
		return false, customErr
	}

	// Iterate through transactions contained in block and make sure that they are valid
	var honestTransaction bool
	switch database.TransactionType(block.Transaction) {
	case "Election":
		honestTransaction, err = validation.IsValidElection(block.Transaction.(database.Election))
	case "Registration":
		honestTransaction, err = validation.IsValidRegistration(block.Transaction.(database.Registration))
	case "Vote":
		honestTransaction, err = validation.IsValidVote(block.Transaction.(database.Vote))
	}
	if !honestTransaction {
		customErr.Message = "Block contains an invalid transaction:\n |" + err.Error() + "\nInvalid block is rejected."
		return false, customErr
	}

	// Make sure that the merkle root is correct
	if CalculateMerkleRoot(block) != block.MerkleRoot {
		customErr.Message = "Block's merkle root is incorrect" + ending
		return false, customErr
	}

	// Make sure that the block hash is correct
	header, err := block.Encode()
	if err != nil {
		return false, err
	}
	hash := crypto.CalculateHash(header)
	if hash != block.Hash {
		customErr.Message = "Block's hash is incorrect" + ending
		return false, customErr
	}

	// Make sure that the block signature is correct
	valid, err := crypto.Verify([]byte(hash), database.PublicKey(block.Validator), block.Signature)
	if err != nil {
		customErr.Message = "Block's signature is invalid\n |" + err.Error() + "\n" + ending
		return false, customErr
	} else if !valid {
		customErr.Message = "Block's signature is invalid" + ending
		return false, customErr
	}

	return true, nil
}

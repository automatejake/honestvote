package consensus

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func IsBlockValid(prevBlock database.Block, block database.Block) (bool, error) {
	customErr := &ConsensusError{
		Time: time.Now(),
	}
	ending := ", invalid block rejected."

	// Make sure that block's index is correct
	if prevBlock.Index+1 != block.Index {
		customErr.Message = "Block index is incorrect" + ending
		logger.Println("verify_block.go", "IsBlockValid()", customErr.Message)
		return false, customErr
	}

	// Make sure that block's previous hash is the last block
	if prevBlock.Hash != block.PrevHash {
		customErr.Message = "Block's previous hash is incorrect" + ending
		logger.Println("verify_block.go", "IsBlockValid()", customErr.Message)
		return false, customErr
	}

	// Make sure the validator is a valid producer
	validator, err := database.FindNode(block.Validator)
	if err != nil {
		logger.Println("verify_block.go", "IsBlockValid()", err)
		return false, err
	} else if validator.Role != "producer" {
		customErr.Message = "Actor proposing this block is not a valid producer." + ending
		logger.Println("verify_block.go", "IsBlockValid()", customErr.Message)
		return false, customErr
	}

	// // Make sure that the block hash is correct
	header, err := block.Encode()
	if err != nil {
		logger.Println("verify_block.go", "IsBlockValid()", err)
		return false, err
	}
	hash := crypto.CalculateHash(header)
	hashString := hex.EncodeToString(hash)
	if hashString != block.Hash {
		customErr.Message = "Block's hash is incorrect" + ending
		logger.Println("verify_block.go", "IsBlockValid()", customErr.Message)
		return false, customErr
	}

	return true, nil
}

func CheckSignature(block database.Block) bool {
	header, err := block.Encode()
	hash := crypto.CalculateHash(header)
	// Make sure that the block signature is correct
	valid, err := crypto.Verify([]byte(hash), block.Validator, block.Signature)
	if err != nil {
		fmt.Println(err)
		return false
	} else if !valid {
		fmt.Println(valid)
		return false
	}

	return true
}

package consensus

import (
	"encoding/hex"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func GenerateBlock(prevBlock database.Block, transactions []string, pubKey string, privKey string) (database.Block, error) {
	var newBlock database.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Format(time.RFC1123)
	newBlock.MerkleRoot = crypto.NewMerkleRoot(transactions)
	newBlock.Validator = pubKey
	newBlock.PrevHash = prevBlock.Hash

	header, err := newBlock.Encode()
	if err != nil {
		logger.Println("create_block.go", "GenerateBlock()", err)
		return database.Block{}, err
	}
	hash := crypto.CalculateHash(header)
	newBlock.Hash = hex.EncodeToString(hash)

	signature, err := crypto.Sign(hash, privKey)
	if err != nil {
		logger.Println("create_block.go", "GenerateBlock()", err)
		newBlock.Signature = "None"
	} else {
		newBlock.Signature = signature
	}

	return newBlock, nil
}

package consensus

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GenerateBlock(prevBlock database.Block, transaction interface{}, pubKey string, privKey string) (database.Block, error) {
	var newBlock database.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	newBlock.Transaction = transaction
	newBlock.MerkleRoot = CalculateMerkleRoot(transaction)
	newBlock.Validator = pubKey
	newBlock.PrevHash = prevBlock.Hash

	header, err := newBlock.Encode()
	if err != nil {
		return database.Block{}, err
	}
	newBlock.Hash = crypto.CalculateHash(header)

	signature, err := crypto.Sign([]byte(newBlock.Hash), privKey)
	if err != nil {
		newBlock.Signature = "None"
	} else {
		newBlock.Signature = signature
	}

	return newBlock, nil
}

func CalculateMerkleRoot(transaction interface{}) string {
	var hash string
	switch database.TransactionType(transaction) {
	case "Registration":
		hash = transaction.(database.Registration).Signature
	case "Vote":
		hash = transaction.(database.Vote).Signature
	case "Election":
		hash = transaction.(database.Election).Signature
	}

	merkleroot := crypto.CalculateHash([]byte(hash))
	return merkleroot
}

package consensus

import (
	"encoding/asn1"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GenerateBlock(prevBlock database.Block, transaction interface{}, pubKey string, privKey string) database.Block {
	var newBlock database.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	newBlock.Transaction = transaction
	newBlock.MerkleRoot = CalculateMerkleRoot(transaction)
	newBlock.Validator = pubKey
	newBlock.PrevHash = prevBlock.Hash

	header := GenerateBlockHeader(newBlock)
	newBlock.Hash = crypto.CalculateHash(header)

	sig, err := crypto.SignBlock(header, privKey)
	if err != nil {
		newBlock.Signature = "None"
	} else {
		newBlock.Signature = sig
	}

	return newBlock
}

func GenerateBlockHeader(block database.Block) []byte {
	header, err := asn1.Marshal(block)
	if err != nil {

	}
	return header
}

// Index       int         `json:"index"`
// Timestamp   string      `json:"timestamp"`
// Transaction interface{} `json:"transaction"` // not  included in the hash
// MerkleRoot  string      `json:"merkleRoot"`
// Validator   string      `json:"validator"`
// Signature   string      `json:"signature"`
// PrevHash    string      `json:"prevhash"`
// Hash        string      `json:"hash"`

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

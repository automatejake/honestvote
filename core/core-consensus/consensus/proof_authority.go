package consensus

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func CalculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func GenerateBlock(pIndex int, pHash string, transaction database.Transaction, port int, pKey string) database.Block {
	var newBlock database.Block

	newBlock.Index = pIndex + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = transaction
	newBlock.PrevHash = pHash
	newBlock.Validator = pKey
	newBlock.Port = port

	header := GenerateHeader(newBlock)

	newBlock.Hash = CalculateHash(header)

	return newBlock
}

func VerifyHash(prevIndex int, prevHash string, block database.Block) bool {
	if prevHash != block.PrevHash {
		return false
	} else if CalculateHash(GenerateHeader(block)) != block.Hash {
		return false
	}

	return true
}

func GenerateHeader(block database.Block) string {
	header := string(block.Index) + block.Timestamp +
		block.Transaction.Sender + string(block.Transaction.Vote) +
		block.Transaction.Receiver + block.PrevHash + block.Validator

	return header
}

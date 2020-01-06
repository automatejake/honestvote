package consensus

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func CalculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func GenerateBlock(pIndex int, pHash string, transaction interface{}, pKey string) database.Block {
	var newBlock database.Block

	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	fmt.Println(timestamp)
	newBlock.Index = pIndex + 1
	newBlock.Timestamp = time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	newBlock.Transaction = transaction
	newBlock.PrevHash = pHash

	newBlock.Signatures = make(map[string]string)
	newBlock.Signatures[pKey] = "Validator"

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
	var header string

	if t, ok := block.Transaction.(database.Transaction); ok {
		header = string(block.Index) + block.Timestamp +
			t.Sender + string(t.Vote) + block.PrevHash

		for _, transaction := range t.Receiver {
			header = header + transaction
		}
	} else if t, ok := block.Transaction.(database.Election); ok {
		header = string(block.Index) + block.Timestamp +
			t.Name + t.RegisteredVoters + t.Start + t.End + block.PrevHash
	} else {
		fmt.Println(t)
	}

	return header
}

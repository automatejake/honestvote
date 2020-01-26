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

	header := GenerateHeader(newBlock)

	newBlock.Hash = CalculateHash(header)
	newBlock.Signatures[pKey] = newBlock.Hash

	fmt.Println(pIndex, newBlock)

	return newBlock
}

func VerifyHash(prevIndex int, prevHash string, block database.Block) bool {
	if prevHash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	}
	// else if CalculateHash(GenerateHeader(block)) != block.Hash {
	// 	fmt.Println("Block hash is wrong!", CalculateHash(GenerateHeader(block)))
	// 	return false
	// }

	return true
}

func GenerateHeader(block database.Block) string {
	var header string

	switch t := block.Transaction.(type) {
	case database.Vote:
		header = string(block.Index) + block.Timestamp +
			string(t.Sender) + t.Election + t.Signature +
			t.Type + block.PrevHash

		for k, v := range t.Receiver {
			header = header + k + v
		}
	case *database.Vote:
		header = string(block.Index) + block.Timestamp +
			string(t.Sender) + t.Election + t.Signature +
			t.Type + block.PrevHash

		for k, v := range t.Receiver {
			header = header + k + v
		}
	case database.Election:
		header = string(block.Index) + block.Timestamp +
			t.ElectionName + t.EmailDomain + t.Institution +
			string(t.Sender) + t.Signature + t.Type + t.Start +
			t.End + block.PrevHash
	case *database.Election:
		header = string(block.Index) + block.Timestamp +
			t.ElectionName + t.EmailDomain + t.Institution +
			string(t.Sender) + t.Signature + t.Type + t.Start +
			t.End + block.PrevHash
	case database.Registration:
		header = string(block.Index) + block.Timestamp + t.Election +
			t.Receiver + string(t.Sender) + t.Signature + t.Type
	case *database.Registration:
		header = string(block.Index) + block.Timestamp + t.Election +
			t.Receiver + string(t.Sender) + t.Signature + t.Type
	}

	return header
}

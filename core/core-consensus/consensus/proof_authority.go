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

	fmt.Println("%d", pIndex, newBlock)

	return newBlock
}

func VerifyHash(prevIndex int, prevHash string, block database.Block) bool {
	if prevHash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	}
	// else if strings.Compare(CalculateHash(GenerateHeader(block)), block.Hash) != 0 {
	// 	fmt.Println("Block hash is wrong!", CalculateHash(GenerateHeader(block)))
	// 	return false
	// }

	return true
}

func GenerateHeader(block database.Block) string {
	var header string

	if t, ok := block.Transaction.(database.Vote); ok {
		header = string(block.Index) + block.Timestamp +
			string(t.Sender) + block.PrevHash

		fmt.Println(t)

		for _, v := range t.Receiver {
			header = header + v
		}
	} else if t, ok := block.Transaction.(*database.Vote); ok {
		header = string(block.Index) + block.Timestamp +
			string(t.Sender) + block.PrevHash

		fmt.Println(t)

		for _, v := range t.Receiver {
			header = header + v
		}
	} else if t, ok := block.Transaction.(database.Election); ok {
		header = string(block.Index) + block.Timestamp +
			t.ElectionName + t.Start + t.End + block.PrevHash
	} else if t, ok := block.Transaction.(*database.Election); ok {
		header = string(block.Index) + block.Timestamp +
			t.ElectionName + t.Start + t.End + block.PrevHash
	} else {
		fmt.Println(block.Transaction)
	}

	return header
}

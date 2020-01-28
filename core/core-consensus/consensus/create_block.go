package consensus

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func CalculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func GenerateBlock(prevBlock database.Block, transaction interface{}, pKey string) database.Block {
	var newBlock database.Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	newBlock.Transaction = transaction
	newBlock.MerkleRoot = CalculateMerkleRoot(transaction)
	newBlock.Validator = pKey
	newBlock.PrevHash = prevBlock.Hash

	index := strconv.Itoa(newBlock.Index)
	header := index + newBlock.Timestamp + newBlock.MerkleRoot + newBlock.Validator + newBlock.PrevHash
	newBlock.Hash = crypto.CalculateHash(header)

	return newBlock
}

func CalculateMerkleRoot(transaction interface{}) string {
	// switch t := block.Transaction.(type)
	// case database.Vote
	return ""
}

func VerifyHash(prevIndex int, prevHash string, block database.Block) bool {
	if prevHash != block.PrevHash {
		fmt.Println("Previous hash is wrong!")
		return false
	} else if CalculateHash(GenerateHeader(block)) != block.Hash {
		fmt.Println("Block hash is wrong!", CalculateHash(GenerateHeader(block)))
		return false
	}

	return true
}

func GenerateHeader(block database.Block) string {
	var header string

	header = string(block.Index) + block.Timestamp + block.PrevHash

	return header
}

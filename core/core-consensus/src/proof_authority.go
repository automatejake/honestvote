package coreconsensus

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

type Block struct {
	Index       int
	Timestamp   string
	Transaction Transaction
	Hash        string
	PrevHash    string
	Validator   string
}

type Transaction struct {
	Sender   string
	Vote     int
	Receiver string
}

var Blockchain []Block
var ProposedBlocks []Block

var Validators []string

var Address string

func calculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func generateBlock(block Block, transaction Transaction) Block {
	var newBlock Block

	newBlock.Index = block.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = transaction
	newBlock.PrevHash = block.Hash
	newBlock.Validator = Address

	header := generateHeader(newBlock)

	newBlock.Hash = calculateHash(header)

	return newBlock
}

func verifyHash(prevBlock, block Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	} else if calculateHash(generateHeader(block)) != block.Hash {
		return false
	}

	return true
}

func generateHeader(block Block) string {
	header := string(block.Index) + block.Timestamp +
		block.Transaction.Sender + string(block.Transaction.Vote) +
		block.Transaction.Receiver + block.PrevHash + block.Validator

	return header
}

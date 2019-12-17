package consensus

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

var Blockchain []database.Block
var ProposedBlocks []database.Block

var Validators []string

var Address string

func CalculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func GenerateBlock(block database.Block, transaction database.Transaction) database.Block {
	var newBlock database.Block

	newBlock.Index = block.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = transaction
	newBlock.PrevHash = block.Hash
	newBlock.Validator = Address

	header := GenerateHeader(newBlock)

	newBlock.Hash = CalculateHash(header)

	return newBlock
}

func VerifyHash(prevBlock, block database.Block) bool {
	if prevBlock.Hash != block.PrevHash {
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

package consensus

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	coredb "github.com/jneubaum/honestvote/core/core-database/src"
)

var Blockchain []coredb.Block
var ProposedBlocks []coredb.Block

var Validators []string

var Address string

func calculateHash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return base64.URLEncoding.EncodeToString(sum)
}

func generateBlock(block coredb.Block, transaction coredb.Transaction) coredb.Block {
	var newBlock coredb.Block

	newBlock.Index = block.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = transaction
	newBlock.PrevHash = block.Hash
	newBlock.Validator = Address

	header := generateHeader(newBlock)

	newBlock.Hash = calculateHash(header)

	return newBlock
}

func verifyHash(prevBlock, block coredb.Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	} else if calculateHash(generateHeader(block)) != block.Hash {
		return false
	}

	return true
}

func generateHeader(block coredb.Block) string {
	header := string(block.Index) + block.Timestamp +
		block.Transaction.Sender + string(block.Transaction.Vote) +
		block.Transaction.Receiver + block.PrevHash + block.Validator

	return header
}

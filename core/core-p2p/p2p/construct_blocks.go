package p2p

import (
	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func Dequeue() interface{} {
	if len(TransactionQueue) > 0 {
		earliestTransaction := TransactionQueue[0]
		TransactionQueue = TransactionQueue[1:]
		return earliestTransaction
	} else {
		return nil
	}
}

func Enqueue(transaction interface{}) {
	TransactionQueue = append(TransactionQueue, transaction)
}

type DecodeTransaction struct {
	Type string `json:"type" bson:"type"`
}

func AddToBlock(transaction interface{}, hash string) {
	block, err := consensus.GenerateBlock(PreviousBlock, transaction, PublicKey, PrivateKey)
	if err != nil {
		logger.Println("read_function.go", "AddToBlock()", err.Error())
	}

	// block.MerkleRoot = hash

	//Check if there is a proposed block currently, if so, add to the queue

	logger.Println("peer_routes.go", "HandleConn()", "Empty, proposing this block.")

	err = database.AddBlock(block)
	if err != nil {
		logger.Println("construct_blocks.go", "AddToBlock()", err)
	} else {
		PreviousBlock = block
		ProposeBlock(block)
	}

}

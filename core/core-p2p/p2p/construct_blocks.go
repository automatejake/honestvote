package p2p

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func Dequeue() interface{} {
	if len(TransactionQueue) > 0 {
		earliestTransaction := TransactionQueue[0]
		TransactionQueue = TransactionQueue[1:]
		logger.Println("construct_blocks.go", "Dequeue()", TransactionQueue)
		return earliestTransaction
	} else {
		return nil
	}
}

func Enqueue(transaction interface{}) {
	logger.Println("construct_blocks.go", "Enqueue()", "Enqueue transaction")
	TransactionQueue = append(TransactionQueue, transaction)
}

type DecodeTransaction struct {
	Type string `json:"type" bson:"type"`
}

//Add transaction to list as hex hash and add interface to database collection for corresponding transaction collection
func AddTransactionToList(transaction interface{}, tranType string) {
	hexTransaction := crypto.HashTransaction(transaction)
	fmt.Println("Hex Transaction: ", hexTransaction)
	TransactionsInBlock = append(TransactionsInBlock, hexTransaction)

	switch tranType {
	case "Election":
		database.AddTransaction(transaction, "elections")
		ProposeTransaction(transaction, "elections")
	case "Registration":
		database.AddTransaction(transaction, "registrations")
		ProposeTransaction(transaction, "registrations")
	case "Vote":
		database.AddTransaction(transaction, "votes")
		ProposeTransaction(transaction, "votes")
	}
}

func CreateBlock() {
	block, err := consensus.GenerateBlock(PreviousBlock, TransactionsInBlock, PublicKey, PrivateKey)

	if err != nil {
		logger.Println("read_function.go", "AddToBlock()", err.Error())
	}

	//Clear the list of transactions in a block for the next set
	TransactionsInBlock = nil

	logger.Println("construct_blocks.go", "CreateBlock()", "Empty, proposing this block.")

	err = database.AddBlock(block)
	if err != nil {
		logger.Println("construct_blocks.go", "AddToBlock()", err)
	} else {
		PreviousBlock = block
		ProposeBlock(block)
	}

}

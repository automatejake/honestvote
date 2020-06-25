package p2p

import (
	"encoding/json"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-validation/validation"
	"github.com/jneubaum/honestvote/core/core-websocket/websocket"
	"github.com/jneubaum/honestvote/tests/logger"
)

/*
Winner of block is leader calculated by((time - first) / step) % node

Given 3 nodes and genesis block time 1:
node 0 - ((8-0)/1) % 3 == 2
node 1 - ((9-0)/1) % 3 == 0
node 2 - ((10-0)/1) % 3 == 1
*/
func BroadcastScheduler() {
	for {
		time := time.Now().UnixNano() / 1000000 //time in milliseconds_
		leader := ((time - GenesisBlockTime) / Step) % ConsensusNodes

		if len(TransactionQueue) > 0 && leader == 0 {
			logger.Println("broadcaster.go", "BroadcastScheduler()", "Going through transactions")

			// create a block from validated transactions in transaction quene
			for i := 0; i < len(TransactionQueue); i++ {
				transaction := Dequeue()
				transaction_json, err := json.Marshal(transaction)
				if err != nil {

				}

				transaction_type := &DecodeTransaction{}
				err = json.Unmarshal(transaction_json, transaction_type)
				if err != nil {

				}

				// check validation and then broadcast block to the network
				switch transaction_type.Type {
				case "Election":
					election := &database.Election{}
					err := json.Unmarshal(transaction_json, election)
					if err != nil {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}
					logger.Println("broadcaster.go", "BroadcastScheduler()", "Received transaction")
					logger.Println("broadcaster.go", "BroadcastScheduler()", election)

					valid, err := validation.IsValidElection(*election)
					if valid {
						//Add transaction to list of transactions in block and save block index to make validating later faster
						election.BlockIndex = PreviousBlock.Index + 1
						AddTransactionToList(*election, transaction_type.Type)
					} else {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}
				case "Registration":
					registration := &database.Registration{}
					err := json.Unmarshal(transaction_json, &registration)
					if err != nil {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}

					valid, err := validation.IsValidRegistration(*registration)

					if valid {
						logger.Println("", "", "Sending Registration")
						websocket.SendRegistration(*registration)

						//Add transaction to list of transactions in block and save block index to make validating faster
						registration.BlockIndex = PreviousBlock.Index + 1
						AddTransactionToList(*registration, transaction_type.Type)
					} else {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}
				case "Vote":
					vote := &database.Vote{}
					err := json.Unmarshal(transaction_json, vote)
					if err != nil {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}

					valid, err := validation.IsValidVote(*vote)

					if valid {
						logger.Println("broadcaster.go", "BroadcastScheduler()", "Passed validation")
						websocket.BroadcastVote(*vote)

						//Add transaction to list of transactions in block and save block index to make validating faster
						vote.BlockIndex = PreviousBlock.Index + 1
						AddTransactionToList(*vote, transaction_type.Type)
					} else {
						logger.Println("broadcaster.go", "BroadcastScheduler()", err)
					}
				}

			}

			if len(TransactionsInBlock) > 0 {
				CreateBlock()
			}
		}
	}
}

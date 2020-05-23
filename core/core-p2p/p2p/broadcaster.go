package p2p

import (
	"time"
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

		if TransactionQuene != nil && leader == 0 {
			// create a block from validated transactions in transaction quene and broadcast it to the network
		}

	}
}

package p2p

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

//Send a block out to be verified by other peers
func ProposeBlock(block database.Block, peers []net.Conn) {
	j, err := json.Marshal(block)

	write := new(Message)
	write.Message = "verify"
	write.Data = j
	write.Type = block.Type

	jWrite, err := json.Marshal(write)

	if err == nil {
		fmt.Println(len(peers))
		for _, peer := range peers {
			peer.Write(jWrite)
		}
	}
}

func DecideType(data []byte, mType string, conn net.Conn) {
	var block database.Block

	if mType == "Vote" {
		vote := &database.Vote{}
		block = database.Block{Transaction: vote}
	} else if mType == "Election" {
		election := &database.Election{}
		block = database.Block{Transaction: election}
	} else if mType == "Registration" {
		registration := &database.Registration{}
		block = database.Block{Transaction: registration}
	}

	json.Unmarshal(data, &block)
	logger.Println("peer_routes.go", "HandleConn()", "Verifying")
	VerifyBlock(block, conn)
}

//Decide if the block sent is valid
func VerifyBlock(block database.Block, conn net.Conn) {
	var valid bool

	if consensus.VerifyHash(PrevIndex, PrevHash, block) {
		fmt.Println("Hash was valid %v", block)
		valid = true
	} else {
		fmt.Println("Hash was not valid %v", block)
		valid = false
	}

	j, err := json.Marshal(block)

	write := new(Message)
	write.Message = "sign"
	write.Data = []byte(strconv.FormatBool(valid))
	write.Signature = make(map[string]string)
	write.Signature[PublicKey], err = crypto.Sign(j, PrivateKey)

	if err != nil {
		write.Signature[PublicKey] = "Signature"
	}

	jWrite, err := json.Marshal(write)

	if err == nil {
		logger.Println("peer_routes.go", "HandleConn()", "Sending response")
		fmt.Println("Sending....")
		conn.Write(jWrite)
	}
}

//Go through all responses from other peers and see the result
func CheckResponses(size int) {
	counter := size

	checkBlock, err := json.Marshal(ProposedBlock)

	if err != nil {
		return
	}

	/*
		Iterate through nested map that holds boolean as first arg and
		a map[string]string as its second
	*/
	for pKey, m := range SignatureMap {
		for sig, b := range m {
			valid, err := crypto.Verify(checkBlock, pKey, sig)
			if valid && err == nil && b {
				fmt.Println("Everything is good.")
				ProposedBlock.Signatures[pKey] = sig
			} else {
				fmt.Println("%t %t", valid, b)
				counter--
			}
		}
	}

	fmt.Println(size, counter)

	if size == counter {
		j, err := json.Marshal(ProposedBlock)

		write := new(Message)
		write.Message = "update"
		write.Data = j

		jWrite, err := json.Marshal(write)

		if err == nil {
			if database.UpdateBlockchain(database.MongoDB, ProposedBlock) {
				PrevHash = ProposedBlock.Hash
				PrevIndex = ProposedBlock.Index
				fmt.Println(string(PrevIndex) + " " + PrevHash)
			}

			for _, node := range Nodes {
				node.Write(jWrite)
			}
		}
	} else {
		logger.Println("data_exchange.go", "CheckResponses", "Someone is a bad actor or this block is wrong.")
	}
}

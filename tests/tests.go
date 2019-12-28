package main

import (
	"fmt"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

// var priv, pub = crypto.GenerateKeys()
var priv, pub = crypto.GenerateKeyPair()
var block database.Block = database.Block{
	Index:       1,
	Timestamp:   "now",
	Transaction: database.Transaction{},
	Hash:        consensus.CalculateHash(consensus.GenerateHeader(database.Block{Index: 1, Timestamp: "now", Transaction: database.Transaction{}})),
	PrevHash:    "",
	Signature:   "",
}

var block2 database.Block = database.Block{
	Index:       1,
	Timestamp:   "now",
	Transaction: database.Transaction{},
	Hash:        consensus.CalculateHash(consensus.GenerateHeader(database.Block{Index: 1, Timestamp: "now", Transaction: database.Transaction{}})),
	PrevHash:    "",
	Signature:   "",
}

func main() {
	byteHash := []byte(block.Hash)

	// fmt.Println(priv + "\n\n" + pub)
	signature, _ := crypto.Sign(byteHash, priv)
	verify, _ := crypto.Verify(byteHash, pub, signature)
	fmt.Println(string(byteHash) + "\n")

	fmt.Println(signature + "\n")
	fmt.Println(verify)

}

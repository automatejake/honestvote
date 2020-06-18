package main

import (
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func TestNewMerkleRoot(t *testing.T) {
	stringArray := make([]string, 0)

	for i := 0; i < 49; i++ {
		hash := crypto.CalculateHash([]byte("hey" + strconv.Itoa(i)))
		stringArray = append(stringArray, hex.EncodeToString(hash))
	}

	tree := crypto.NewMerkleRoot(stringArray)

	for i := 0; i < len(stringArray)-1; i++ {
		yes := crypto.MerkleProof(stringArray[i], tree.RootNode)

		t.Log(yes)
	}
}

func TestVerifyTransaction(t *testing.T) {

}

package main

import (
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func TestNewMerkleRoot(t *testing.T) {
	stringArray := make([]string, 0)

	for i := 0; i < 10; i++ {
		hash := crypto.CalculateHash([]byte("hey" + strconv.Itoa(i)))
		stringArray = append(stringArray, hex.EncodeToString(hash))
	}

	t.Log(stringArray)

	tree := crypto.NewMerkleRoot(stringArray)

	yes := crypto.IsIntroverse(stringArray[3], tree.RootNode)

	if yes {
		t.Log("The transaction is verified.")
	} else {
		t.Log("The transaction is not verified.")
	}
}

func TestVerifyTransaction(t *testing.T) {

}

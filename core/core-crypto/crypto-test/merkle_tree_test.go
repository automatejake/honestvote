package main

import (
	"strconv"
	"testing"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func TestNewMerkleRoot(t *testing.T) {
	byteArray := make([][]byte, 0)

	for i := 0; i < 10; i++ {
		hash := crypto.CalculateHash([]byte("hey" + strconv.Itoa(i)))
		byteArray = append(byteArray, hash)
	}

	t.Log(byteArray)

	tree := crypto.NewMerkleRoot(byteArray)

	t.Log(*tree.RootNode)
}

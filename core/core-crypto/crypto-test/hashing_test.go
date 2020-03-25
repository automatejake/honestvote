package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func TestCalculateHash(t *testing.T) {
	test := []byte("Here's a test string.")
	hash := crypto.CalculateHash(test)

	if hash == "" {
		t.Error("The hash shouldn't be nil: ", hash)
		return
	}

	t.Log("The hash is not nil: ", hash)
}

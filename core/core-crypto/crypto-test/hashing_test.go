package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func TestCalculateHash(t *testing.T) {
	test := []byte("Here's a test string.")
	hash := crypto.CalculateHash(test)

	t.Log("The hash is not nil: ", hash)
}

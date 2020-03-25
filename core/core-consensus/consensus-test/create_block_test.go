package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"

	"github.com/jneubaum/honestvote/core/core-p2p/p2p"

	"github.com/jneubaum/honestvote/core/core-consensus/consensus"
)

func TestGenerateBlock(t *testing.T) {
	block, err := consensus.GenerateBlock(p2p.PreviousBlock, database.Block{}, p2p.PublicKey, p2p.PrivateKey)

	if err != nil {
		t.Error()
	}

	t.Log(block)
}

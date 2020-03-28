package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

func TestListenConn(t *testing.T) {
	// p2p.AddToBlock
	err := p2p.ListenConn("7000", "")

	if err != nil {
		t.Error("There shouldn't be an error when trying to listen for a connection.")
	}
}

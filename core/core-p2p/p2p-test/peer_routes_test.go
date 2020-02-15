package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

func TestSum(t *testing.T) {
	p2p.ListenConn("7000", "hello")

	// Errorf("Sum was incorrect, got: %d, want: %d.")

}

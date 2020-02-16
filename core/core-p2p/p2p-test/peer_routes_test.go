package main

import (
	"encoding/json"
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

func TestAcceptConnectMessage(t *testing.T) {
	logger.Mode = "all"
	// p2p.AcceptConnectMessage
	v := database.Vote{}
	vote, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
	}
	err = p2p.ReceiveTransaction("Vote", vote)
	if err != nil {
		t.Error("Functions failed", err)
	}
}

func TestAddToBlock(t *testing.T) {
	// p2p.AddToBlock
}

func TestSendIndex(t *testing.T) {
	// p2p.SendIndex()
}

func TestGrabDocuments(t *testing.T) {
	// p2p.GrabDocuments()
}

func TestMoveDocuments(t *testing.T) {

}

func TestProposeBlock(t *testing.T) {

}

func TestLatestHashAndIndex(t *testing.T) {

}

func TestSendRegistrationTransaction(t *testing.T) {

}

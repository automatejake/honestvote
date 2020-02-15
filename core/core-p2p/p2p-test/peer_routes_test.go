package main

import (
	"encoding/json"
	"testing"

	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

func TestAcceptConnectMessage(t *testing.T) {
	logger.Mode = "all"
	// p2p.AcceptConnectMessage
	v := Vote{}
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

type Vote struct {
	Type      string              `json:"type" bson:"type"`
	Election  string              `json:"electionName" bson:"electionName"` //Data Start
	Receiver  []SelectedCandidate `json:"receivers" bson:"receivers"`       //Data End
	Sender    PublicKey           `json:"sender" bson:"sender"`
	Signature string              `json:"signature" bson:"signature"`
}

type SelectedCandidate struct {
	PositionId string `json:"id" bson:"id"`
	Recipient  string `json:"key" bson:"key"`
}

type PublicKey string

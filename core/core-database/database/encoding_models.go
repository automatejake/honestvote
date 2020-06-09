package database

import (
	"encoding/asn1"

	"github.com/jneubaum/honestvote/tests/logger"
)

type EncodedBlock struct {
	Index      int    `json:"index"`
	Timestamp  string `json:"timestamp"`
	MerkleRoot string `json:"merkleRoot"`
	PrevHash   string `json:"prevhash"`
	Hash       string `json:"hash"`
}

func (b Block) Encode() ([]byte, error) {
	object := EncodedBlock{
		Index:     b.Index,
		Timestamp: b.Timestamp,
		// MerkleRoot: b.MerkleRoot.RootNode.Hash,
		PrevHash: b.PrevHash,
		Hash:     b.Hash,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, err
	}

	return encoded, nil
}

type EncodedElection struct {
	ElectionName string     `json:"electionName"` //Data Start
	Description  string     `json:"description"`
	Start        string     `json:"startDate"`
	End          string     `json:"endDate"`
	EmailDomain  string     `json:"emailDomain"`
	Positions    []Position `json:"positions"` //Data End
}

func (e Election) Encode() ([]byte, error) {
	object := EncodedElection{
		ElectionName: e.ElectionName,
		Description:  e.Description,
		Start:        e.Start,
		End:          e.End,
		EmailDomain:  e.EmailDomain,
		Positions:    e.Positions,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, err
	}

	return encoded, nil
}

type EncodedRegistration struct {
	Election    string `json:"electionId"` //Data Start
	Receiver    string `json:"receiver"`   //Data End
	RecieverSig string `json:"recieverSig"`
}

func (r Registration) Encode() ([]byte, error) {
	object := EncodedRegistration{
		Election:    r.Election,
		Receiver:    r.Receiver,
		RecieverSig: r.RecieverSig,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		logger.Println("encoding_models.go", "Encode(Registration)", err)
		return encoded, err
	}

	return encoded, nil
}

type EncodedVote struct {
	Election string              `json:"electionId"` //Data Start
	Receiver []SelectedCandidate `json:"receivers"`  //Data End
}

func (v Vote) Encode() ([]byte, error) {
	object := EncodedVote{
		Election: v.Election,
		Receiver: v.Receiver,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		logger.Println("encoding_models.go", "EncodeVote()", err)
		return encoded, err
	}

	return encoded, nil
}

type EncodedNode struct {
	Institution string `json:"institution" bson:"institution"`
	IPAddress   string `json:"ipaddress" bson:"ipaddress"`
	Port        int    `json:"port" bson:"port"`
	Timestamp   string `json:"timestamp" bson:"timestamp"`
	Role        string `json:"role" bson:"role"` // peer | full | registry
}

func (n Node) Encode() ([]byte, error) {
	object := EncodedNode{
		Institution: n.Institution,
		IPAddress:   n.IPAddress,
		Port:        n.Port,
		Timestamp:   n.Timestamp,
		Role:        n.Role,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		logger.Println("encoding_models.go", "Encode(Node)", err)
		return encoded, err
	}
	return encoded, nil

}

type EncodedTransaction interface {
	Encode() ([]byte, error)
}

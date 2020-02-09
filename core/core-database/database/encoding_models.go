package database

import (
	"encoding/asn1"
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
		Index:      b.Index,
		Timestamp:  b.Timestamp,
		MerkleRoot: b.MerkleRoot,
		PrevHash:   b.PrevHash,
		Hash:       b.Hash,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, nil
	}

	return nil, nil
}

type EncodedElection struct {
	ElectionName string     `json:"electionName"` //Data Start
	Institution  string     `json:"institutionName"`
	Description  string     `json:"description"`
	Start        string     `json:"startDate"`
	End          string     `json:"endDate"`
	EmailDomain  string     `json:"emailDomain"`
	Positions    []Position `json:"positions"` //Data End
}

func (e Election) Encode() ([]byte, error) {
	object := EncodedElection{
		ElectionName: e.ElectionName,
		Institution:  e.Institution,
		Description:  e.Description,
		Start:        e.Start,
		End:          e.End,
		EmailDomain:  e.EmailDomain,
		Positions:    e.Positions,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, nil
	}

	return encoded, nil
}

type EncodedRegistration struct {
	Election    string    `json:"election"` //Data Start
	Receiver    PublicKey `json:"receiver"` //Data End
	RecieverSig string    `json:"recieverSig"`
}

func (r Registration) Encode() ([]byte, error) {
	object := EncodedRegistration{
		Election:    r.Election,
		Receiver:    r.Receiver,
		RecieverSig: r.RecieverSig,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, nil
	}

	return encoded, nil
}

type EncodedVote struct {
	Election string              `json:"electionName"` //Data Start
	Receiver []SelectedCandidate `json:"receivers"`    //Data End
}

func (v Vote) Encode() ([]byte, error) {
	object := EncodedVote{
		Election: v.Election,
		Receiver: v.Receiver,
	}
	encoded, err := asn1.Marshal(object)
	if err != nil {
		return encoded, nil
	}

	return encoded, nil
}

type EncodedTransaction interface {
	Encode() ([]byte, error)
}

package http

import "github.com/jneubaum/honestvote/core/core-database/database"

type ElectionInfo struct {
}

func Translate(e database.Election) ElectionInfo {
	return ElectionInfo{}
}

type Vote struct {
	Sender    string `json:"voterId"`
	Election  string `json:"electionId"`
	Candidate string `json:"candidateId"`
	Signature string `json:"signature"`
}

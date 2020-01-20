package http

import "github.com/jneubaum/honestvote/core/core-database/database"

type ElectionInfo struct {
}

// id: ElectionId
// electionName: string
// institutionName: string
// description: string
// startDate: string
// endDate: string

func Translate(e database.Election) ElectionInfo {
	return ElectionInfo{}
}

type Vote struct {
	Sender    string `json:"voterId"`
	Election  string `json:"electionId"`
	Candidate string `json:"candidateId"`
	Signature string `json:"signature"`
}

type Payload struct {
	Status    string      `json:"status"`
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data"`
}

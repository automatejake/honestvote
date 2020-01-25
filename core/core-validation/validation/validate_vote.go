package validation

import "github.com/jneubaum/honestvote/core/core-database/database"

func IsValidVote(r database.Vote) bool {

	return true
}

type Vote struct {
	Type      string             `json:"type"`
	Election  string             `json:"election"` //Data Start
	Receiver  map[string]string  `json:"receiver"` //Data End
	Sender    database.PublicKey `json:"sender"`
	Signature string             `json:"signature"`
}

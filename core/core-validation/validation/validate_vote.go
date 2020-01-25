package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidVote(v database.Vote) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	ending := ", invalid transaction fails"

	//Check to see if election is a valid election
	election, err := database.GetElection(v.Election)
	if err != nil {
		customErr.Message = "Vote transactions must specify a valid election" + ending +
			err.Error()
		return false, customErr
	}

	//Check to see if election is an ongoing election
	now := time.Now()
	electionEnd, err := time.Parse(election.End, "Mon, 02 Jan 2006 15:04:05 MST")
	if now.After(electionEnd) {
		customErr.Message = "Vote transactions must occur for elections that are still ongoing" + ending
		return false, customErr
	}

	//Check to see if voter is registered to vote
	registration := database.CorrespondingRegistration(v)
	if registration.Sender != v.Sender {
		customErr.Message = "Vote transactions must have a corresponding registration transaction" + ending
		return false, customErr
	}

	//Check to see if vote went to valid candidates
	for _, position := range election.Positions {
		for _, _ = range position.Candidates {
			// if candidate.Recipient ==
		}
		// if v.Receiver[position.PositionId] !=
	}

	//Check to see if Vote type is correctly stored in transaction
	if v.Type != "Vote" {
		customErr.Message = "Transaction is incorrect type" + ending
		return false, customErr
	}

	customErr = nil
	return true, customErr
}

type Vote struct {
	Type      string             `json:"type"`
	Election  string             `json:"election"` //Data Start
	Receiver  map[string]string  `json:"receiver"` //Data End
	Sender    database.PublicKey `json:"sender"`
	Signature string             `json:"signature"`
}

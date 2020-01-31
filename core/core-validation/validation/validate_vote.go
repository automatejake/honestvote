package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidVote(v database.Vote) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	ending := ", invalid transaction fails"

	//Check to see if signature is valid
	voteHeaders := v.Election
	for key, value := range v.Receiver {
		voteHeaders += key + value
	}

	valid, err := crypto.Verify([]byte(voteHeaders), v.Sender, v.Signature)
	if err != nil {

	}
	if !valid {
		customErr.Message = "Vote transaction contains invalid signature" + ending
		return false, customErr
	}

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
	for i, _ := range election.Positions {
		// if ContainsPositionCandidate(v.Receiver[election.Positions[i].PositionId]) {
		if !ContainsCandidate(election.Positions[i], v.Receiver[election.Positions[i].PositionId]) {
			customErr.Message = "Vote transaction must be for a legitimate candidate" + ending
			return false, customErr
		}
	}

	//Check to see if Vote type is correctly stored in transaction
	if v.Type != "Vote" {
		customErr.Message = "Transaction is incorrect type" + ending
		return false, customErr
	}

	return true, nil
}

func ContainsCandidate(p database.Position, c string) bool {
	for _, candidate := range p.Candidates {
		if c == candidate.Recipient {
			return true
		}
	}
	return false
}

package validation

import (
	"fmt"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GenerateVoteHeaders(v database.Vote) (string, error) {
	encoded, err := v.Encode()
	if err != nil {
		return "", err
	}

	hash := crypto.CalculateHash(encoded)
	return hash, nil

}

func IsValidVote(v database.Vote) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	ending := ", invalid transaction fails"

	//Check to see if signature is valid
	voteHeaders, err := GenerateVoteHeaders(v)
	if err != nil {
		return false, err
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
	fmt.Println(election, "\n", v.Election)
	if err != nil {
		customErr.Message = "Vote transactions must specify a valid election" + ending +
			err.Error()
		return false, customErr
	}

	//Check to see if election is an ongoing election
	now := time.Now()
	electionEnd, err := time.Parse(time.RFC1123, election.End)
	if now.After(electionEnd) {
		customErr.Message = "Vote transactions must occur for elections that are still ongoing" + ending
		return false, customErr
	}

	//Check to see if voter is registered to vote
	registration := database.CorrespondingRegistration(v)
	if registration.Receiver != v.Sender {
		customErr.Message = "Vote transactions must have a corresponding registration transaction" + ending
		return false, customErr
	}

	//Check to see if vote choice is valid (this check is not 100% perfect with a map, but does not poise harm)
	eligibleCandidates := map[string]int{}
	for _, position := range election.Positions {
		for _, candidate := range position.Candidates {
			eligibleCandidates[position.PositionId+candidate.Recipient] = 1
		}
	}
	for _, recipient := range v.Receiver {
		if eligibleCandidates[recipient.PositionId+recipient.Recipient] == 0 {
			customErr.Message = "Vote transactions must be for valid candidates" + ending
			return false, customErr
		}
		if eligibleCandidates[recipient.PositionId+recipient.Recipient] > 1 {
			customErr.Message = "Vote transactions cannot contain multiple selections for a single candidate" + ending
			return false, customErr
		}
		eligibleCandidates[recipient.PositionId+recipient.Recipient]++
	}

	//Check to see if Vote type is correctly stored in transaction
	if v.Type != "Vote" {
		customErr.Message = "Transaction is incorrect type" + ending
		return false, customErr
	}

	//Make sure that vote does not occur twice
	if database.ContainsVote(v.Sender, v.Election) {
		customErr.Message = "Vote transaction has already been cast for this voter" + ending
		return false, customErr
	}

	return true, nil
}

// func ContainsCandidate(p database.Position, c string) bool {
// 	for _, candidate := range p.Candidates {
// 		if c == candidate.Recipient {
// 			return true
// 		}
// 	}
// 	return false
// }

package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func GenerateVoteHeaders(v database.Vote) ([]byte, error) {
	encoded, err := v.Encode()
	if err != nil {
		logger.Println("validate_vote.go", "GenerateVoteHeaders()", err)
		return nil, err
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
		logger.Println("validate_vote.go", "IsValidVote()", err)
		return false, err
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Generated headers")

	valid, err := crypto.Verify(voteHeaders, v.Sender, v.Signature)
	if err != nil {
		logger.Println("validate_vote.go", "IsValidVote()", err)
		return false, err
	}
	if !valid {
		customErr.Message = "Vote transaction contains invalid signature" + ending
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Signature verified as valid")

	//Check to see if election is a valid election
	election, err := database.GetElection(v.Election)
	if err != nil {
		customErr.Message = "Vote transactions must specify a valid election" + ending +
			err.Error()
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Vote is for a valid election")

	//Check to see if election is an ongoing election
	now := time.Now()
	electionEnd, err := time.Parse(time.RFC1123, election.End)
	if now.After(electionEnd) {
		customErr.Message = "Vote transactions must occur for elections that are still ongoing" + ending
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Vote is for an ongoing election")

	//Check to see if voter is registered to vote
	logger.Println("validate_vote.go", "IsValidVote()", v)
	registration := database.CorrespondingRegistration(v)
	if registration.Receiver != v.Sender {
		customErr.Message = "Vote transactions must have a corresponding registration transaction" + ending
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Vote contains a corresponding registration")

	//Check to see if vote choice is valid (this check is not 100% perfect with a map, but does not poise harm)
	eligibleCandidates := map[string]int{}
	for _, position := range election.Positions {
		for _, candidate := range position.Candidates {
			eligibleCandidates[candidate.Name+position.PositionId] = 1
		}
	}
	for _, recipient := range v.Receiver {
		if eligibleCandidates[recipient.Recipient+recipient.PositionId] == 0 {
			customErr.Message = "Vote transactions must be for valid candidates" + ending
			logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
			return false, customErr
		}
		if eligibleCandidates[recipient.Recipient+recipient.PositionId] > 1 {
			customErr.Message = "Vote transactions cannot contain multiple selections for a single candidate" + ending
			logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
			return false, customErr
		}
		eligibleCandidates[recipient.Recipient+recipient.PositionId]++
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Vote is for eligible positions and candidates")

	//Check to see if Vote type is correctly stored in transaction
	if v.Type != "Vote" {
		customErr.Message = "Transaction is incorrect type" + ending
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}

	//Make sure that vote does not occur twice
	if database.ContainsVote(v.Sender, v.Election) {
		customErr.Message = "Vote transaction has already been cast for this voter" + ending
		logger.Println("validate_vote.go", "IsValidVote()", customErr.Message)
		return false, customErr
	}
	logger.Println("validate_vote.go", "IsValidVote()", "Voter has not voted more than once.")
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

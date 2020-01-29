package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidElection(e database.Election) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	end := ", invalid transaction fails"

	//Check to see if signature is valid
	electionHeaders := e.ElectionName + e.Institution + e.Description + e.Start + e.End + e.EmailDomain
	for _, position := range e.Positions {
		electionHeaders += position.PositionId + position.Name
		for _, candidate := range position.Candidates {
			electionHeaders += candidate.Name + candidate.Recipient
		}
	}

	valid, err := crypto.Verify([]byte(electionHeaders), e.Sender, e.Signature)
	if err != nil {

	}
	if !valid {
		customErr.Message = "Election transaction contains invalid signature" + end
		return false, customErr
	}

	//Check to see if sender matches the public key of a legitimate administrator node
	node := database.FindNode(string(e.Sender))
	if node.PublicKey != "producer" {
		customErr.Message = "Election transaction is not permitted by node without administrator capabilities" + end
		return false, customErr
	}

	//Check to see if institution matches public key of sender
	if e.Institution != node.Institution {
		customErr.Message = "Election transaction must come from the correct institution" + end
		return false, customErr
	}

	//Check to see if Election type is correctly stored in transaction
	if e.Type != "Election" {
		customErr.Message = "Election transaction is incorrect type" + end
		return false, customErr
	}

	//Check to see if election end is valid
	now := time.Now()
	electionEnd, er := time.Parse(e.End, "Mon, 02 Jan 2006 15:04:05 MST")
	if er != nil {
		customErr.Message = "Election transaction contains an invalid date format"
		return false, customErr
	}
	if now.Before(electionEnd) {
		customErr.Message = "Election transaction end date is already past" + end
		return false, customErr
	}

	//Check to see if election contains postions with unique ids and candidates with uniqued recipient ids
	positionSet := make(map[string]bool)
	candidateSet := make(map[string]bool)
	for _, position := range e.Positions {

		if positionSet[position.PositionId] {
			customErr.Message = "Election transaction contains multiple position ids for a single transaction" + end
			return false, customErr
		}
		positionSet[position.PositionId] = true

		for _, candidate := range position.Candidates {
			if candidate.Recipient == "" {
				if candidateSet[candidate.Recipient] {
					customErr.Message = "Election transaction contains multiple recipients for a single transaction" + end
					return false, customErr
				}
				candidateSet[candidate.Recipient] = true
			}
		}
	}

	//if all passes, then transaction is valid
	customErr = nil
	return true, customErr
}

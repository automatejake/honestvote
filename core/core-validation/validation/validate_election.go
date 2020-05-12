package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func GenerateElectionHeaders(e database.Election) ([]byte, error) {

	encoded, err := e.Encode()
	if err != nil {
		logger.Println("validate_election.go", "GenerateElectionHeaders()", err)
		return nil, err
	}

	hash := crypto.CalculateHash(encoded)

	return hash, nil
}

func IsValidElection(e database.Election) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	end := ", invalid transaction fails"

	//Check to see if signature is valid
	electionHeaders, err := GenerateElectionHeaders(e)
	if err != nil {
		logger.Println("validate_election.go", "IsValidElection()", err)
		return false, err
	}

	valid, err := crypto.Verify(electionHeaders, e.Sender, e.Signature)
	if err != nil {
		logger.Println("validate_election.go", "IsValidElection()", err)
		return false, customErr
	}
	if !valid {
		customErr.Message = "Election transaction contains invalid signature" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//Check to see if sender matches the public key of a legitimate administrator node
	node, err := database.FindNode(string(e.Sender))
	if err != nil {
		customErr.Message = "Election transaction does not contain a node registered in the database" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}
	if node.Role != "producer" {
		customErr.Message = "Election transaction is not permitted by node without administrator capabilities" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//Check to see if institution matches public key of sender
	if e.Institution != node.Institution {
		customErr.Message = "Election transaction must come from the correct institution" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//Check to see if Election type is correctly stored in transaction
	if e.Type != "Election" {
		customErr.Message = "Election transaction is incorrect type" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//Check to see if election end is valid
	now := time.Now()
	electionEnd, er := time.Parse(time.RFC1123, e.End)
	if er != nil {
		customErr.Message = "Election transaction contains an invalid date format"
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}
	if now.After(electionEnd) {
		customErr.Message = "Election transaction end date is already past" + end
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//Check to see if election contains postions with unique ids and candidates with uniqued recipient ids
	positionSet := make(map[string]bool)
	candidateSet := make(map[string]bool)
	for _, position := range e.Positions {

		if positionSet[position.PositionId] {
			customErr.Message = "Election transaction contains multiple position ids for a single transaction" + end
			logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
			return false, customErr
		}
		positionSet[position.PositionId] = true

		for _, candidate := range position.Candidates {
			if candidate.Recipient == "" {
				if candidateSet[candidate.Recipient] {
					customErr.Message = "Election transaction contains multiple recipients for a single transaction" + end
					logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
					return false, customErr
				}
				candidateSet[candidate.Recipient] = true
			}
		}
	}

	//Check to see if election signature doesn't exist in the database
	newSig := database.CheckElectionSignature(e.Signature)
	if !newSig {
		customErr.Message = "Election signature already exists in the database, this is copy of a previous election."
		logger.Println("validate_election.go", "IsValidElection()", customErr.Message)
		return false, customErr
	}

	//if all passes, then transaction is valid

	return true, nil
}

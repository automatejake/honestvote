package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidElection(e database.Election) (bool, error) {
	err := &ValidationError{
		Time: time.Now(),
	}
	end := ", invalid transaction fails"

	//Check to see if sender matches the public key of a legitimate administrator node
	node := database.FindNode(string(e.Sender))
	if node.PublicKey != "producer" {
		err.Message = "Transaction is not permitted by node without administrator capabilities" + end
		return false, err
	}

	//Check to see if institution matches public key of sender
	if e.Institution != node.Institution {
		err.Message = "Transaction must come from the correct institution" + end
		return false, err
	}

	//Check to see if Election type is correctly stored in transaction
	if e.Type != "Election" {
		err.Message = "Transaction is incorrect type" + end
		return false, err
	}

	//Check to see if election end is valid
	check := time.Time{}
	now, er := time.Parse(e.End, "Mon, 02 Jan 2006 15:04:05 MST")
	if er != nil {
		err.Message = "Transaction contains an invalid date format"
		return false, err
	}
	if check.Before(now) {
		err.Message = "Transaction end date is already past" + end
		return false, err
	}

	//Check to see if election contains postions with unique ids and candidates with uniqued recipient ids
	positionSet := make(map[string]bool)
	candidateSet := make(map[string]bool)
	for _, position := range e.Positions {

		if positionSet[position.PositionId] {
			err.Message = "Transaction contains multiple position ids for a single transaction" + end
			return false, err
		}
		positionSet[position.PositionId] = true

		for _, candidate := range position.Candidates {
			if candidate.Recipient == "" {
				if candidateSet[candidate.Recipient] {
					err.Message = "Transaction contains multiple recipients for a single transaction" + end
					return false, err
				}
				candidateSet[candidate.Recipient] = true
			}
		}
	}

	//if all passes, then transaction is valid
	err = nil
	return true, err
}

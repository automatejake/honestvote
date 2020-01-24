package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidElection(e database.Election) (bool, error) {
	err := &ValidationError{
		Time: time.Now(),
	}
	end := ", transaction is invalid"

	//Check to see if sender matches the public key of a legitimate administrator node
	if e.End != "" {
		err.Message = "" + end
		return false, err
	}

	//Check to see if Election type is correctly stored in transaction
	if e.Type != "Election" {
		err.Message = "Transaction is incorrect type" + end
		return false, err
	}

	//Check to see if institution matches public key of sender
	if e.ElectionName != "Election" {
		err.Message = "Transaction is incorrect type" + end
		return false, err
	}

	//Check to see if election end is valid
	if e.End != "" {
		err.Message = "Transaction end date is already past" + end
		return false, err
	}

	//Check to see if election contains postions with unique ids and candidates with uniqued recipient ids
	if e.End != "" {

		err.Message = "Transaction " + end
		return false, err
	}

	err = nil
	return true, err
}

type Election struct {
	Type         string     `json:"type"`
	ElectionName string     `json:"electionName"` //Data Start
	Institution  string     `json:"institutionName"`
	Description  string     `json:"description"`
	Start        string     `json:"startDate"`
	End          string     `json:"endDate"`
	EmailDomain  string     `json:"emailDomain"`
	Positions    []Position `json:"positions"` //Data End
	Sender       string     `json:"sender"`
	Signature    string     `json:"id"`
}

type Position struct {
	PositionId string      `json:"id"`
	Name       string      `json:"displayName"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name      string `json:"name"`
	Recipient string `json:"key"`
}

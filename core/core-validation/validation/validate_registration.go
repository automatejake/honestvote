package validation

import (
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func IsValidRegistration(r database.Registration) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	ending := ", invalid tranaction fails"

	//Check to see if election exists
	election, err := database.GetElection(r.Election)
	if err != nil {
		customErr.Message = "Registration transactions must specify a valid election" + ending +
			err.Error()
		return false, customErr
	}

	//Check to see if registration was sent by the administrator that declared the election
	if r.Sender != election.Sender {
		customErr.Message = "Registration transactions must be delcared by an administrator" + ending
		return false, customErr
	}

	//Check to see if registration is for a valid public key
	if r.Sender != "" {
		customErr.Message = "Registration transactions must come from a voter with a valid public key" + ending
		return false, customErr
	}

	//Check to see if Registration type is correctly stored in transaction
	if r.Type != "Registration" {
		customErr.Message = "Registration transaction is incorrect type" + ending
		return false, customErr
	}

	//if all passes, then transaction is valid
	customErr = nil
	return true, customErr
}

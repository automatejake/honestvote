package validation

import (
	"fmt"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GenerateRegistrationHeaders(r database.Registration) (string, error) {
	encoded, err := r.Encode()
	if err != nil {
		return "", err
	}

	hash := crypto.CalculateHash(encoded)
	return hash, nil
}

func IsValidRegistration(r database.Registration) (bool, error) {
	customErr := &ValidationError{
		Time: time.Now(),
	}
	ending := ", invalid tranaction fails"

	//Check to see if signature is valid
	registrationHeaders, err := GenerateRegistrationHeaders(r)
	if err != nil {
		return false, err
	}

	valid, err := crypto.Verify([]byte(registrationHeaders), r.Sender, r.Signature)
	if !valid {
		customErr.Message = "Registration transaction contains invalid signature" + ending
		fmt.Println(err)
		return false, customErr

	}

	//Check to see if election exists
	// election, err := database.GetElection(r.Election)
	// if err != nil {
	// 	customErr.Message = "Registration transactions must specify a valid election" + ending +
	// 		err.Error()
	// 	return false, customErr
	// }

	// //Check to see if election is still ongoing
	// now := time.Now()
	// electionEnd, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", election.End)
	// if err != nil {
	// 	customErr.Message = "Registration transaction contains an invalid date format" + ending
	// 	return false, customErr
	// }
	// if now.After(electionEnd) {
	// 	customErr.Message = "Registration transactions must occur for elections that are still ongoing" + ending
	// 	return false, customErr
	// }

	// // Check to see if registration was sent by the administrator that declared the election
	// if r.Sender != election.Sender {
	// 	customErr.Message = "Registration transactions must be delcared by an administrator" + ending
	// 	return false, customErr
	// }

	// Check to see if registration is for a valid public key
	if r.Sender == "" {
		customErr.Message = "Registration transactions must come from a voter with a valid public key" + ending
		return false, customErr
	}

	//Check to see if Registration type is correctly stored in transaction
	if r.Type != "Registration" {
		customErr.Message = "Registration transaction is incorrect type" + ending
		return false, customErr
	}

	//Check to see if same person has already registered to vote
	alreadyVoted := database.ContainsRegistration(r.Sender, r.Election)
	if alreadyVoted {
		customErr.Message = "Registration transaction already exists for this voter" + ending
		return false, customErr
	}

	//if all passes, then transaction is valid
	return true, nil
}

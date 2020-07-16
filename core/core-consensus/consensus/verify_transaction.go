package consensus

import (
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func VerifyTransactions(block database.Block) (bool, error) {
	vElection, err := database.GrabElectionsInBlock(block)

	if !vElection || err != nil {
		return false, err
	}

	vRegistration, err := database.GrabRegistrationsInBlock(block)

	if !vRegistration || err != nil {
		return false, err
	}

	vVote, err := database.GrabVotesInBlock(block)

	if !vVote || err != nil {
		return false, err
	}

	logger.Println("verify_transaction.go", "VerifyTransactions()", "Everything was verified and good!")

	return true, nil
}

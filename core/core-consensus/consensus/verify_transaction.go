package consensus

import (
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func VerifyTransactions(index int) bool {
	var hashedTransactions []string

	elections, err := database.GrabElectionsInBlock(index)

	if err != nil {
		logger.Println("verify_transaction.go", "VerifyTransactions()", err)
	}

	for _, election := range elections {
		hash := crypto.HashTransaction(election)
		hashedTransactions = append(hashedTransactions, hash)
	}

	return true

}

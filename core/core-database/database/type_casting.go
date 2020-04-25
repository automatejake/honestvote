package database

import "github.com/jneubaum/honestvote/tests/logger"

func TransactionType(transaction interface{}) string {

	switch transaction.(type) {
	case Vote:
		return "Vote"
	case Registration:
		return "Registration"
	case Election:
		return "Election"
	default:

	}

	logger.Println("type_casting.go", "TransactionType()", "Transaction is not a vote, registration, or election")

	return ""
}

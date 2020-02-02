package database

import "fmt"

func TransactionType(transaction interface{}) string {

	switch t := transaction.(type) {
	case Vote:
		return "Vote"
	case Registration:
		return "Registration"
	case Election:
		return "Election"
	default:
		fmt.Println(t)
	}

	return ""
}

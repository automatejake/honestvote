package database

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

	return ""
}

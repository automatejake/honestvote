package main

import (
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func TestCustomError(t *testing.T) {
	err := database.CustomError{Message: "This is a test error."}

	rError := err.Error()

	if rError == "" {
		t.Error("Error was not displayed correctly: ", rError)
		return
	}

	t.Log("Here's the error you want to display: ", rError)
}

func TestElectionConvertInfo(t *testing.T) {
	election := new(database.Election)

	conversion := election.ConvertInfo()

	t.Log("Here's the converted election: ", conversion)
}

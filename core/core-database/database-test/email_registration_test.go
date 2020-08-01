package main

import (
	"testing"
	"github.com/jneubaum/honestvote/core/core-database/database"

func TestSaveRegistrationCode(t *testing.T) {
	//database.SaveRegistrationCode()
}
func TestIsValidRegistrationCode(t *testing.T) {
	//database.IsValidRegistrationCode()
}

func TestCheckEmailVerification(t *testing.T){
	database.MongoDB = database.MongoConnect("localhost")
	database.CollectionPrefix = "test"
	var check bool := database.CheckEmailVerification(registrant)
	if check != true{
		t.Log("Already registered")
	}else{
		t.Log("Good to go!")
	}




}

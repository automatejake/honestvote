package registration

import (
	"net/smtp"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/tests/logger"
)

/*
*  Send data to administrator node
*  Make sure election is still ongoing
*  Make sure email checks regex format and has not already been registered
*  Create registration code
*  Save AwaitingRegistration object to database
*  Send confirmation email with code and link
*  Recieve confirmation link response
*  Verify registrant is a student from national clearinghouse
*  Delete student from database and store record of their email in hashed format
*  Send registration transaction
 */

func IsValidEmailRegistrant(registrant *database.AwaitingRegistration) bool {
	if !isElectionOngoing(registrant.Email) {
		logger.Println("email_registration.go", "IsValidEmailRegistrant()", "Election is not ongoing")
		return false
	}

	if !isValidEmail(registrant.Email) {
		logger.Println("email_registration.go", "IsValidEmailRegistrant()", "Email is not valid")
		return false
	}

	if hasRegisteredWithEmail(registrant.Email) {
		logger.Println("email_registration.go", "IsValidEmailRegistrant()", "Voter has already registered to vote")
		return false
	}

	return true

}

//Make sure election is ongoing
func isElectionOngoing(email string) bool {
	return true
}

//Make sure email checks regex format and has not already been registered
func isValidEmail(email string) bool {
	return true
}

func hasRegisteredWithEmail(email string) bool {
	return false
}

// Create registration code, save to database, send email with code and link
func SendRegistrationCode(registrant database.AwaitingRegistration, hostname string, tcp_port string, email_address string, email_password string) {
	registrant.Timestamp = time.Now().Format(time.RFC3339)
	registrant.Code, _ = crypto.RandomHex(100)
	database.SaveRegistrationCode(registrant)

	email := registrant.Email
	from := email_address  //should be environmental variable that is updated by administrator
	pass := email_password //should be environmental variable that is updated by administrator
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Code" + "\n\n" +
		"Click this link if you requested to register for the upcoming student election: \nhttps://" + hostname + "/verifyCode/code=" + registrant.Code + "&verified=true\n" +
		"If this is incorrect, please click here:\nhttps://" + hostname + "/verifyCode/code=" + registrant.Code + "&verified=false"

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		logger.Println("email_registration.go", "SendRegistrationCode", err.Error())
		return
	}

}

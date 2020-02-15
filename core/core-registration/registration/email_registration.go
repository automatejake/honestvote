package registration

import (
	"fmt"
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

func IsValidRegistrant(registrant *database.AwaitingRegistration) bool {
	if !isElectionOngoing(registrant.Email) {
		return false
	}

	if !isValidEmail(registrant.Email) {
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

// Create registration code, save to database, send email with code and link
func SendRegistrationCode(registrant database.AwaitingRegistration, public_ip string, tcp_port string) {
	registrant.Timestamp = time.Now().Format(time.RFC3339)
	registrant.Code, _ = crypto.RandomHex(100)
	database.SaveRegistrationCode(registrant)

	fmt.Println("Sending registration")

	email := registrant.Email
	from := "testhonestvote.io@gmail.com" //should be environmental variable that is updated by administrator
	pass := "Passw0rd123!"                //should be environmental variable that is updated by administrator
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Code" + "\n\n" +
		"Click this link if you requested to register for the upcoming student election: \n" + public_ip + ":" + tcp_port + "/verifyCode/code=" + registrant.Code + "&verified=true\n" +
		"If this is incorrect, please click here:\n" + public_ip + ":" + tcp_port + "/verifyCode/code=" + registrant.Code + "&verified=false"

	fmt.Println("\n\n" + msg + "\n\n")
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		logger.Println("email_registration.go", "SendRegistrationCode", err.Error())
		return
	}

}

func SendWarningEmail(email string, election string) {
	from := "testhonestvote.io@gmail.com" //should be environmental variable that is updated by administrator
	pass := "Passw0rd123!"                //should be environmental variable that is updated by administrator
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Warning" + "\n\n" +
		"You indicated that someone attempted to register falsely with your school email.  Please register to vote as soon as possible."

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		logger.Println("email_registration.go", "SendRegistrationCode", err.Error())
		return
	}

}

package registration

import (
	"net/smtp"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/tests/logger"
)

func EmailRegistration(registrantEmail string, election string, public_key string, public_ip string, tcp_port string) {

	//check if valid election
	if !isValidElection(election) {
		return
	}

	//regex check and check if student has already voted
	if !isValidEmail(registrantEmail) {
		return
	}

	// save registration code in database
	code, _ := crypto.RandomHex(100)

	registrant := database.AwaitingRegistration{
		Election:  election,
		PublicKey: public_key,
		Code:      code,
		Timestamp: time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"),
	}

	database.SaveRegistrationCode(registrant)

	// send email verification code
	SendRegistrationCode(registrantEmail, election, code, public_ip, tcp_port)

}

func isValidEmail(email string) bool {
	return true
}

func isValidElection(election string) bool {
	return true
}

func SendRegistrationCode(email string, election string, code string, public_ip string, tcp_port string) {

	from := "testhonestvote.io@gmail.com" //should be environmental variable that is updated by administrator
	pass := "Passw0rd123!"                //should be environmental variable that is updated by administrator
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Code" + "\n\n" +
		"Click this link if you requested to register for the upcoming" + election + "election: \n" + public_ip + ":" + tcp_port + "/verifyCode/code=" + code + "&verified=true\n" +
		"If this is incorrect, please click here:\n" + public_ip + ":" + tcp_port + "/verifyCode/code=" + code + "&verified=false"

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		logger.Println("email_registration.go", "SendRegistrationCode", err.Error())
		return
	}

}

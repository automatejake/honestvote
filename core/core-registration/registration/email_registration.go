package registration

import (
	"fmt"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/tests/logger"
)

func EmailRegistration(registrantEmail string, election string, public_key string) {

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
		Timestamp: time.Now().String(),
	}

	database.SaveRegistrationCode(registrant)

	// send email verification code
	SendRegistrationCode(registrantEmail, election, code)

}

func isValidEmail(email string) bool {
	return true
}

func isValidElection(election string) bool {
	return true
}

func VerifyRegistrationCode(code string) {
	//check if registration link has expired (should expire after x time, e.g. 1 hour or less for extra security)
	valid, public_key := database.IsValidRegistrationCode(code)
	if valid {
		fmt.Println(public_key)
	}

}

func SendRegistrationCode(email string, election string, code string) {

	from := "testhonestvote.io@gmail.com" //should be environmental variable that is updated by administrator
	pass := "Passw0rd123!"                //should be environmental variable that is updated by administrator
	to := email

	tcp_service := strconv.Itoa(p2p.TCP_SERVICE)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Code" + "\n\n" +
		"Click this link if you requested to register for the upcoming" + election + "election: \n" + p2p.PublicIP + tcp_service + "/verifyCode/" + code

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		logger.Println("email_registration.go", "SendRegistrationCode", err.Error())
		return
	}

}

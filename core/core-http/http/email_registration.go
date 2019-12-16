package http

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func EmailRegistration(registrantEmail string) {
	//regex check

	//check if valid election

	//check if registration link has expired (should expire after x time, e.g. 1 hour or less for extra security)

	//send email verification code
	sendRegistrationCode(registrantEmail)
}

func isValidEmail(email string) bool {
	return true
}

func isValidElection(email string) bool {
	return true
}

func sendRegistrationCode(email string) {

	code, _ := crypto.RandomHex(100)

	from := "testhonestvote.io@gmail.com" //should be environmental variable that is updated by administrator
	pass := "Passw0rd123!"                //should be environmental variable that is updated by administrator
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:  " + "HonestVote Registration Code" + "\n\n" +
		"Click this link to register for the election: \n http://portainer.honestvote.io:7001/verifyCode/" + code

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Click this link to register for the election: \n http://portainer.honestvote.io:7001/verifyCode/" + code)

	// log.Print("sent mail")
}

package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)

	registrant, err := database.IsValidRegistrationCode(params["id"])
	if err != nil {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", err.Error())
		return
	}
	var check bool
	check = database.CheckEmailVerification(registrant)
	if check == false {
		w.Write([]byte("You have already registered."))
		return
	}

	switch p2p.REGISTRATION_TYPE {
	case "EXTERNAL_WHITELIST":
		if !registration.OnWhitelist(registrant.Email, p2p.Whitelist) {
			w.Write([]byte("You are not permitted to participate in this election.  Please talk to your election administrator if you think that this is a mistake."))
			return
		}
	case "DEFAULT_WHITELIST":

	}

	if params["verified"] == "true" {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", string(registrant.Sender)+" is registered to vote for "+registrant.ElectionName)
		if registration.VerifyStudent(registrant) {
			err := p2p.SendRegistrationTransaction(registrant)
			if err != nil {
				logger.Println("producer_handlers.go", "VerifyEmailHandler()", "Registration Transaction not sent correctly. "+err.Error())
			} else {
				w.Write([]byte("You have registered successfully!  Go back to the app to vote."))
			}

		}
	} else if params["verified"] == "false" {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", string(registrant.Sender)+" is not supposed to be registered to vote for "+registrant.ElectionName)
		database.RemoveRegistrationCode(registrant)
		w.Write([]byte("You indicated that you did not register to vote.  It may be possible that a malicious actor is attempting to register with your identity.  It is recommended to register as soon as possible."))
	} else {

	}

}

func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	endpoint, err := database.GetEndpoint()
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
		logger.Println("producer_handlers.go", "GetEndpoint()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = endpoint

	}
	json.NewEncoder(w).Encode(payload)

}

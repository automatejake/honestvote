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
		registration.SendWarningEmail(registrant.ElectionName, registrant.Email, p2p.Email_Address, p2p.Email_Password)
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

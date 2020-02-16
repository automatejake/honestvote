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
		return
	}

	if params["verified"] == "true" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", string(registrant.Sender)+" is registered to vote for "+registrant.ElectionName)
		if registration.VerifyStudent(registrant) {
			err := p2p.SendRegistrationTransaction(registrant)
			if err != nil {

			} else {
				w.Write([]byte("You have registered successfully!  Go back to the app to vote."))
			}
		}
	} else if params["verified"] == "false" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", string(registrant.Sender)+" is not supposed to be registered to vote for "+registrant.ElectionName)
		registration.SendWarningEmail(registrant.Email, registrant.ElectionName)
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
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = endpoint

	}
	json.NewEncoder(w).Encode(payload)

}

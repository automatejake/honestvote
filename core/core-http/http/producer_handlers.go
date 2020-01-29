package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	registrant, err := database.IsValidRegistrationCode(params["id"])
	if err != nil {
		return
	}
	if params["verified"] == "true" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", string(registrant.Sender)+" is registered to vote for "+registrant.ElectionName)
		if registration.VerifyStudent(registrant) {

		}
	} else if params["verified"] == "false" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", string(registrant.Sender)+" is not supposed to be registered to vote for "+registrant.ElectionName)
		registration.SendWarningEmail(params["email"], registrant.ElectionName)
	}

}

func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	endpoint, err := database.GetEndpoint()
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
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

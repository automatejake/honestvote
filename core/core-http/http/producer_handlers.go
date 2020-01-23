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
	public_key, election, valid := database.IsValidRegistrationCode(params["id"])

	if valid && params["verified"] == "true" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", public_key+" is registered to vote for "+election)

		// p2p.ReceiveTransaction(1)
	} else if params["verified"] == "false" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", public_key+" is not supposed to be registered to vote for "+election)
		// implement logic to allow email to register again
		registration.SendWarningEmail(params["email"], election)
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

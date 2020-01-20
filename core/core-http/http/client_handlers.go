package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	// elections := database.GetElections()
	// fmt.Println(elections)
	// json.NewEncoder(w).Encode(elections)
}

func GetElectionHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	election := database.GetElection(params["electionid"])
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Status:    "OK",
		Timestamp: timestamp,
		Data:      election,
	}
	json.NewEncoder(w).Encode(payload)
}

func GetVotesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	voters := database.GetVotes(params["electionid"])
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Status:    "OK",
		Timestamp: timestamp,
		Data:      voters,
	}
	json.NewEncoder(w).Encode(payload)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Println("positions")
	// positions := database.GetPositions()
	// json.NewEncoder(w).Encode(positions)

}

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

func PostPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

func PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

func PostElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

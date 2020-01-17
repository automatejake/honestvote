package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Println("candidates")
	candidates := database.GetCandidates()
	jsonCandidates, err := json.Marshal(candidates)
	if err != nil {
		logger.Println("full_http_routes.go", "GetCandidatesHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonCandidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	elections := database.GetElections()
	fmt.Println(elections)
	json.NewEncoder(w).Encode(elections)
}

func GetElectionHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	election := database.GetElection(params["electionid"])
	json.NewEncoder(w).Encode([]database.API_Election{election})
}

func GetVotesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Println("voters")
	voters := database.GetVoters()
	json.NewEncoder(w).Encode(voters)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	fmt.Println("positions")
	positions := database.GetPositions()
	json.NewEncoder(w).Encode(positions)

}

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	tickets := database.GetTickets()
	json.NewEncoder(w).Encode(tickets)
}

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {

}

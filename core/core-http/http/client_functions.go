package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
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
	jsonElections, err := json.Marshal(elections)
	if err != nil {
		logger.Println("full_http_routes.go", "GetElectionsHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonElections)
}

func GetElectionHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	election := database.GetElection(params["id"])
	jsonElection, err := json.Marshal(election)
	if err != nil {
		logger.Println("full_http_routes.go", "GetElectionsHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonElection)
}

func GetVotersHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	voters := database.GetVoters()
	jsonVoters, err := json.Marshal(voters)
	if err != nil {
		logger.Println("full_http_routes.go", "GetElectionsHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonVoters)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	positions := database.GetPositions()
	jsonPositions, err := json.Marshal(positions)
	if err != nil {
		logger.Println("full_http_routes.go", "GetElectionsHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonPositions)

}

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	tickets := database.GetTickets()
	jsonTickets, err := json.Marshal(tickets)
	if err != nil {
		logger.Println("full_http_routes.go", "GetElectionsHandler()", err.Error())
	}
	json.NewEncoder(w).Encode(jsonTickets)
}

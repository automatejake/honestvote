package corehttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	coredb "github.com/jneubaum/honestvote.io/core/core-database/src"
)

var Router = mux.NewRouter()

var MockCandidates = []coredb.Candidate{
	coredb.Candidate{Name: "Jimmy", PublicKey: "0x54khfn4", Election: "Spring 2020"},
	coredb.Candidate{Name: "Janice", PublicKey: "0xflkh45n", Election: "Spring 2020"},
	coredb.Candidate{Name: "Larry", PublicKey: "0xij04ng3", Election: "Spring 2020"}}

var MockElections = []coredb.Election{
	coredb.Election{Name: "West Chester University", RegisteredVoters: "1023"},
	coredb.Election{Name: "Temple", RegisteredVoters: "103"},
	coredb.Election{Name: "Drexel", RegisteredVoters: "6433"},
	coredb.Election{Name: "UPenn", RegisteredVoters: "9023"}}

func HandleRoutes() {
	Router.HandleFunc("/getCandidates", GetCandidatesHandler)
	Router.HandleFunc("/getElections", GetElectionsHandler)
	http.Handle("/", Router)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockCandidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockElections)
}

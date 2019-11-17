package corehttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	coredb "github.com/jneubaum/honestvote.io/core/core-database/src"
)

var Router = mux.NewRouter()

var MockData = []coredb.Candidate{
	coredb.Candidate{Name: "Jimmy", PublicKey: "0x54khfn4", Election: "Spring 2020"},
	coredb.Candidate{Name: "Janice", PublicKey: "0xflkh45n", Election: "Spring 2020"},
	coredb.Candidate{Name: "Larry", PublicKey: "0xij04ng3", Election: "Spring 2020"}}

func HandleRoutes() {
	Router.HandleFunc("/getCandidates", GetCandidatesHandler)
	http.Handle("/", Router)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockData)
}

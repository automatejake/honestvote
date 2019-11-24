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

var Codes []string

//temporary function for demo data, get rid of when real database implemented
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func HandleRoutes() {
	Router.HandleFunc("/getCandidates", GetCandidatesHandler).Methods("GET")
	Router.HandleFunc("/getElections", GetElectionsHandler).Methods("GET")
	Router.HandleFunc("/verifyCode", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
	http.Handle("/", Router)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockCandidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MockElections)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	registrant := r.FormValue("email")
	EmailRegistration(registrant)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {

	//should be with database
	if contains(Codes, "r.code request") {
		//distributes vote to public key

	}

}

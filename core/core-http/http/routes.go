package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

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
	Router.HandleFunc("/getVoters", GetVotersHandler).Methods("GET")
	Router.HandleFunc("/getPositions", GetPositionsHandler).Methods("GET")
	Router.HandleFunc("/getTickets", GetTicketsHandler).Methods("GET")
	Router.HandleFunc("/verifyCode", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
	http.Handle("/", Router)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Candidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Elections)
}

func GetVotersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Voters)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Positions)
}

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tickets)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	registrant := r.FormValue("email")
	EmailRegistration(registrant)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {

	// //should be with database
	// if contains(Codes, "r.code request") {
	// 	//distributes vote to public key

	// }

}

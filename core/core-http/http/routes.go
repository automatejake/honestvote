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
	Router.HandleFunc("/candidates", GetCandidatesHandler).Methods("GET")
	Router.HandleFunc("/elections", GetElectionsHandler).Methods("GET")
	Router.HandleFunc("/voters", GetVotersHandler).Methods("GET")
	Router.HandleFunc("/positions", GetPositionsHandler).Methods("GET")
	Router.HandleFunc("/tickets", GetTicketsHandler).Methods("GET")
	Router.HandleFunc("/verifyCode", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
	http.Handle("/", Router)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	json.NewEncoder(w).Encode(Candidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	json.NewEncoder(w).Encode(Elections)
}

func GetVotersHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	json.NewEncoder(w).Encode(Voters)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	json.NewEncoder(w).Encode(Positions)
}

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	json.NewEncoder(w).Encode(Tickets)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	registrant := r.FormValue("email")
	EmailRegistration(registrant)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {

	// //should be with database
	// if contains(Codes, "r.code request") {
	// 	//distributes vote to public key

	// }

}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var FullRouter = mux.NewRouter()

func HandleFullRoutes() {
	FullRouter.HandleFunc("/candidates", GetCandidatesHandler).Methods("GET")
	FullRouter.HandleFunc("/elections", GetElectionsHandler).Methods("GET")
	FullRouter.HandleFunc("/voters", GetVotersHandler).Methods("GET")
	FullRouter.HandleFunc("/positions", GetPositionsHandler).Methods("GET")
	FullRouter.HandleFunc("/tickets", GetTicketsHandler).Methods("GET")
	FullRouter.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
	http.Handle("/", FullRouter)
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
	//registrant := r.FormValue("email")
	//EmailRegistration(registrant)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

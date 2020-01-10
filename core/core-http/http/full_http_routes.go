package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

var FullRouter = mux.NewRouter()

func HandleFullRoutes() {
	FullRouter.HandleFunc("/candidates", GetCandidatesHandler).Methods("GET")
	FullRouter.HandleFunc("/election", GetElectionsHandler).Methods("GET")
	FullRouter.HandleFunc("/voters", GetVotersHandler).Methods("GET")
	FullRouter.HandleFunc("/positions", GetPositionsHandler).Methods("GET")
	FullRouter.HandleFunc("/tickets", GetTicketsHandler).Methods("GET")

	// <Full Node IP Address>:<Full Node Port>/candidates
	// <Full Node IP Address>:<Full Node Port>/election?id=<ElectionId>
	// <Full Node IP Address>:<Full Node Port>/voters
	// <Full Node IP Address>:<Full Node Port>/positions
	// <Full Node IP Address>:<Full Node Port>/tickets

	//this needs to be encrypted (send admin's public key and encrypted message containing email and public key)
	FullRouter.HandleFunc("/registerElection/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")

	FullRouter.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
	http.Handle("/", FullRouter)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// json.NewEncoder(w).Encode(Candidates)
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// json.NewEncoder(w).Encode(Elections)
}

func GetVotersHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// json.NewEncoder(w).Encode(Voters)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// json.NewEncoder(w).Encode(Positions)
}

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// json.NewEncoder(w).Encode(Tickets)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	// simply send message to peer node in future
	//port := strconv.Itoa(p2p.Self.Port)
	//registration.EmailRegistration(params["email"], params["election"], params["public_key"], p2p.Self.IPAddress, port)

	registration := database.AwaitingRegistration{Email: params["email"], Election: params["election"], PublicKey: params["public_key"]}
	j, err := json.Marshal(registration)

	message := new(p2p.Message)
	message.Message = "registration"
	message.Data = j

	jMessage, err := json.Marshal(message)

	if err == nil {
		p2p.Nodes[0].Write(jMessage)
	}

	//registrant := r.FormValue("email")
	//EmailRegistration(registrant)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

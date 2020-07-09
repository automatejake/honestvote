package http

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/core/core-validation/validation"
	"github.com/jneubaum/honestvote/tests/logger"
)

func PostRegisterHandler(w http.ResponseWriter, r *http.Request) {

	SetupResponse(&w, r)
	decoder := json.NewDecoder(r.Body)
	var registrant database.AwaitingRegistration
	err := decoder.Decode(&registrant)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error decoding registrant - "+err.Error())
	}

	registrant_json, err := json.Marshal(registrant)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error marshalling registrant into JSON - "+err.Error())
	}
	var message p2p.Message = p2p.Message{
		Message: "register",
		Data:    registrant_json,
		// Type:    "",
	}

	byte_message, err := json.Marshal(message)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error converting message to json - "+err.Error())
	}

	admin, err := database.FindNode(registrant.ElectionAdmin)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error finding node from database - "+err.Error())
	}

	port := strconv.Itoa(admin.Port)
	conn, err := net.Dial("tcp", admin.IPAddress+":"+port)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error dialing administrator node - "+err.Error())
	}

	conn.Write(byte_message)

}

func PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	decoder := json.NewDecoder(r.Body)
	var vote database.Vote
	err := decoder.Decode(&vote)

	if err != nil {
		logger.Println("client_handler.go", "PostVoteHandler", "Error decoding vote - "+err.Error())
	}
	vote.Type = "Vote"

	// Add transaction to quene
	valid, err := validation.IsValidVote(vote)
	if valid {
		logger.Println("client_handlers.go", "PostVoteHandler", "Vote transaction is being added to the queue")
		logger.Println("client_handlers.go", "PostVoteHandler", vote)
		p2p.Enqueue(vote)
	} else {
		logger.Println("broadcaster.go", "BroadcastScheduler()", err)
	}
}

func PostElectionHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	decoder := json.NewDecoder(r.Body)
	var election database.Election
	err := decoder.Decode(&election)
	if err != nil {
		logger.Println("client_handler.go", "PostElectionHandler", "Error decoding election - "+err.Error())
	}
	election.Type = "Election"

	// Add transaction to quene
	valid, err := validation.IsValidElection(election)
	if valid {
		logger.Println("client_handler.go", "PostElectionHandler", "Election transaction is being added to the queue")
		logger.Println("client_handler.go", "PostElectionHandler", election)
		p2p.Enqueue(election)
	} else {
		logger.Println("client_handler.go", "PostElectionHandler", "Election transaction is invalid, not added to queue")
	}

}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	elections, err := database.GetElections()
	var electionInfos []database.ElectionInfo
	for _, election := range elections {
		electionInfos = append(electionInfos, election.ConvertInfo())
	}
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}

	if electionInfos == nil {
		electionInfos = []database.ElectionInfo{}
	}

	if err != nil {
		logger.Println("client_handlers.go", "GetElectionsHandler()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = electionInfos
	}
	json.NewEncoder(w).Encode(payload)
}

func GetElectionHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)

	election, err := database.GetElection(params["electionid"])
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
		logger.Println("client_handlers.go", "GetElectionHandler()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = election
	}
	json.NewEncoder(w).Encode(payload)
}

func GetVotesHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)
	votes, err := database.GetVotes(params["electionid"])

	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if votes == nil {
		votes = []database.Vote{}
	}
	if err != nil {
		logger.Println("client_handlers.go", "GetVotesHandler()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = votes
	}
	json.NewEncoder(w).Encode(payload)
}

func GetPositionsHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

}

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)
	permissions, err := database.GetPermissions(params["publickey"])
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if permissions == nil {
		permissions = []string{}
	}
	if err != nil {
		logger.Println("client_handlers.go", "GetPermissionsHandler()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = permissions
	}
	json.NewEncoder(w).Encode(payload)
}

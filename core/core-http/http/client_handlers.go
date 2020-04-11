package http

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
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
	v, err := json.Marshal(vote)
	if err != nil {
		logger.Println("client_handler.go", "PostVoteHandler", "Error converting vote to json - "+err.Error())
	}

	p2p.ReceiveTransaction("Vote", v)
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
	e, err := json.Marshal(election)
	if err != nil {
		logger.Println("client_handler.go", "PostElectionHandler", "Error converting election to json - "+err.Error())
	}

	p2p.ReceiveTransaction("Election", e)

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

	fmt.Println(r, "\n\nElection id", params["electionid"], "\n", params)

	election, err := database.GetElection(params["electionid"])
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
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
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = permissions
	}
	json.NewEncoder(w).Encode(payload)
}

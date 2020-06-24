package http

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

func PostRequestAdminPrivileges(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	decoder := json.NewDecoder(r.Body)
	var request database.RequestAdminPrivileges
	err := decoder.Decode(&request)

	logger.Println("client_handlers.go", "PostRequestPrivileges", request)

	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", "Error decoding registrant - "+err.Error())
	}

	message := []byte("requesting administrator privileges")
	valid_request, err := crypto.Verify(message, request.PublicKey, request.Signature)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", err)
		return
	}
	if !valid_request {
		logger.Println("client_handler.go", "PostRegisterHandler", "Invalid Signature")
		return
	}

	logger.Println("client_handler.go", "PostRegisterHandler", "Valid signature, nominating full node as an administrator")
	// var nomination database.Election{
	// 	ElectionName: "",
	// 	Institution: "",

	// }

	// p2p.Enqueue(nomination)

}

// type Election struct {
// 	Type            string          `json:"type" bson:"type"`
// 	ElectionName    string          `json:"electionName" bson:"electionName"` //Data Start
// 	Institution     string          `json:"institutionName" bson:"institutionName"`
// 	Description     string          `json:"description" bson:"description"`
// 	Start           string          `json:"startDate" bson:"startDate"`
// 	End             string          `json:"endDate" bson:"endDate"`
// 	EmailDomain     string          `json:"emailDomain" bson:"emailDomain"`
// 	ElectionOptions ElectionOptions `json:"electionOptions" bson:"electionOptions"`
// 	Positions       []Position      `json:"positions" bson:"positions"` //Data End
// 	Sender          string          `json:"sender" bson:"sender"`
// 	Signature       string          `json:"signature" bson:"signature"`
// 	BlockIndex      int             `json:"blockIndex" bson:"blockIndex"`
// }

// type ElectionOptions struct {
// 	ElectionType             string `json:"electionType" bson:"electionType"`                         // (producer nomination | default), producer nomination election is a special election declared to elect a node
// 	ShowDataDuringElection   string `json:"showDataDuringElection" bson:"showDataDuringElection"`     // (during | after voting | after election end), logic handled by client
// 	AllowedVotesPerVoter     int    `json:"allowedVotesPerVoter" bson:"allowedVotesPerVoter"`         // amount of votes allowed to be cast by each voter, default is 1
// 	MultipleVotesPerPosition bool   `json:"multipleVotesPerPosition" bson:"multipleVotesPerPosition"` // default false
// }

// type Position struct {
// 	PositionId string      `json:"id" bson:"id"`
// 	Name       string      `json:"displayName" bson:"displayName"`
// 	Candidates []Candidate `json:"candidates" bson:"candidates"`
// }

// type Candidate struct {
// 	Recipient string `json:"key" bson:"key"`
// 	Name      string `json:"name" bson:"name"`
// }

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
	logger.Println("client_handlers.go", "PostVoteHandler", "Vote transaction is being added to the queue")
	logger.Println("client_handlers.go", "PostVoteHandler", vote)

	// Add transaction to quene
	p2p.Enqueue(vote)
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
	logger.Println("client_handler.go", "PostElectionHandler", "Election transaction is being added to the queue")
	logger.Println("client_handler.go", "PostElectionHandler", election)

	// Add transaction to quene
	p2p.Enqueue(election)

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

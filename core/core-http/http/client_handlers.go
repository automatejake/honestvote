package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

func PostPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

}

func PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var vote database.Vote
	err := decoder.Decode(&vote)
	if err != nil {
		panic(err)
	}
	vote.Type = "Vote"
	v, err := json.Marshal(vote)
	if err != nil {

	}

	p2p.ReceiveTransaction(v, "Vote")
}

func PostElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var election database.Election
	err := decoder.Decode(&election)
	if err != nil {
		panic(err)
	}
	election.Type = "Election"
	e, err := json.Marshal(election)
	if err != nil {

	}

	p2p.ReceiveTransaction(e, "Election")

}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	elections, err := database.GetElections()
	var electionInfos []database.ElectionInfo
	for _, election := range elections {
		electionInfos = append(electionInfos, election.ConvertInfo())
		fmt.Println(election.ConvertInfo())
	}
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Timestamp: timestamp,
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
	EnableCors(&w)
	params := mux.Vars(r)
	election, err := database.GetElection(params["electionid"])
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
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
	EnableCors(&w)
	params := mux.Vars(r)
	votes, err := database.GetVotes(params["electionid"])
	// var voteInfos []database.VoteInfo
	// for _, vote := range votes {
	// 	voteInfos = append(voteInfos, vote.ConvertInfo())
	// }

	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Timestamp: timestamp,
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
	EnableCors(&w)

}

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	permissions, err := database.GetPermissions(params["publickey"])
	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = permissions
	}
	json.NewEncoder(w).Encode(payload)
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

package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	elections, err := database.GetElections()
	var electionInfos []database.ElectionInfo
	for _, election := range elections {
		electionInfos = append(electionInfos, election.ConvertInfo())
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
	var voteInfos []database.VoteInfo
	for _, vote := range votes {
		voteInfos = append(voteInfos, vote.ConvertInfo())
	}

	timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = voteInfos
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

func PostPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

func PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

func PostElectionsHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
}

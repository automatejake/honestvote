package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	// timestamp := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	private_key, public_key := crypto.GenerateKeyPair()
	fmt.Println("Admin Private Key:\n" + private_key + "\n")
	fmt.Println("Admin Public Key\n" + public_key + "\n")

	start := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	end := time.Now().AddDate(0, 0, 200).Format("Mon, 02 Jan 2006 15:04:05 MST") //200 days in  the future

	var election database.Election = database.Election{
		Type:         "Election",
		ElectionName: "Student Government Elections",
		Institution:  "West Chester University",
		Description:  "Spring Elections",
		Start:        start,
		End:          end,
		EmailDomain:  "wcupa.edu",
		Sender:       database.PublicKey(public_key),
	}

	election.Positions = []database.Position{
		database.Position{
			PositionId: "demfrmeororev",
			Name:       "Student Government President",
			Candidates: []database.Candidate{
				database.Candidate{
					Name:      "John Doe",
					Recipient: "test1",
				},
				database.Candidate{
					Name:      "Sarah Jennings",
					Recipient: "test2",
				},
				database.Candidate{
					Name:      "Maximus Footless",
					Recipient: "test3",
				},
			},
		},
	}

	jsonElection, _ := json.Marshal(election)
	signature, err := crypto.Sign(jsonElection, private_key)
	if err != nil {
		fmt.Println(err)
	}
	election.Signature = signature

	var registration database.Registration = database.Registration{
		Type:     "Registration",
		Election: election.Signature,
		Receiver: "test3",
		Sender:   database.PublicKey(public_key),
	}

	jsonRegistration, _ := json.Marshal(registration)
	signature, err = crypto.Sign(jsonRegistration, private_key)
	if err != nil {
		fmt.Println(err)
	}

	registration.Signature = signature

	private_key, public_key = crypto.GenerateKeyPair()
	var vote database.Vote = database.Vote{
		Type:     "Vote",
		Election: election.Signature,
		Receiver: map[string]string{"demfrmeororev": "test1"},
		Sender:   database.PublicKey(public_key),
	}

	fmt.Println("Voter Private Key:\n" + private_key + "\n")
	fmt.Println("Voter Public Key\n" + public_key + "\n")

	jsonVote, _ := json.Marshal(vote)
	signature, err = crypto.Sign(jsonVote, private_key)
	if err != nil {
		fmt.Println(err)
	}
	vote.Signature = signature

	jsonElection, _ = json.MarshalIndent(election, "", "\t")
	jsonRegistration, _ = json.MarshalIndent(registration, "", "\t")
	jsonVote, _ = json.MarshalIndent(vote, "", "\t")

	// jsonData, _ := json.Marshal(jsonArray)

	filename := "_mock_data.json"
	_ = ioutil.WriteFile("election"+filename, jsonElection, 777)
	_ = ioutil.WriteFile("registration"+filename, jsonRegistration, 777)
	_ = ioutil.WriteFile("vote"+filename, jsonVote, 777)
}

// valid votes have a corresponding registration transaction with the public key
// type Vote struct {
// 	Type      string            `json:"type"`
// 	Election  string            `json:"election"` //Data Start
// 	Receiver  map[string]string `json:"receiver"` //Data End
// 	Sender    string            `json:"sender"`
// 	Signature string            `json:"signature"`
// }

// type Registration struct {
// 	Type      string `json:"type"`
// 	Election  string `json:"election"` //Data Start
// 	Receiver  string `json:"receiver"` //Data End
// 	Sender    string `json:"sender"`
// 	Signature string `json:"signature"`
// }

// type Election struct {
// 	Type         string     `json:"type"`
// 	ElectionName string     `json:"electionName"` //Data Start
// 	Institution  string     `json:"institutionName"`
// 	Description  string     `json:"description"`
// 	Start        string     `json:"startDate"`
// 	End          string     `json:"endDate"`
// 	EmailDomain  string     `json:"emailDomain"`
// 	Positions    []Position `json:"positions"` //Data End
// 	Sender       string     `json:"sender"`
// 	Signature    string     `json:"id"`
// }

// type Position struct {
// 	PositionId string      `json:"id"`
// 	Name       string      `json:"displayName"`
// 	Candidates []Candidate `json:"candidates"`
// }

// type Candidate struct {
// 	Name      string `json:"name"`
// 	Recipient string `json:"key"`
// }

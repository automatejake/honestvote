package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jneubaum/honestvote/core/core-validation/validation"

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

	headers := validation.GenerateElectionHeaders(election)
	election.Signature, _ = crypto.Sign([]byte(headers), private_key)

	var registration database.Registration = database.Registration{
		Type:     "Registration",
		Election: election.Signature,
		Receiver: "test3",
		Sender:   database.PublicKey(public_key),
	}

	jsonRegistration, _ := json.Marshal(registration)
	headers = validation.GenerateRegistrationHeaders(registration)
	registration.Signature, _ = crypto.Sign([]byte(headers), private_key)

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
	// vote.Signature = p2p.CreateSignature(vote, private_key)

	jsonElection, _ = json.MarshalIndent(election, "", "\t")
	jsonRegistration, _ = json.MarshalIndent(registration, "", "\t")
	jsonVote, _ = json.MarshalIndent(vote, "", "\t")

	// jsonData, _ := json.Marshal(jsonArray)

	filename := "mock_data.json"
	// _ = ioutil.WriteFile(filename, jsonElection, 0644)
	// _ = ioutil.WriteFile(filename, jsonRegistration, 0644)
	// _ = ioutil.WriteFile(filename, jsonVote, 0644)

	file, _ := os.Create(filename)
	defer file.Close()

	_, _ = io.WriteString(file, "Election Transaction:\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonElection)+"' http://localhost:7003/election\n\n\n\n")
	_, _ = io.WriteString(file, "Registration Transaction:\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonRegistration)+"' http://localhost:7003/test/registration\n\n\n\n")
	_, _ = io.WriteString(file, "Vote Transaction:\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonVote)+"' http://localhost:7003/test/vote\n\n\n\n")

}

// os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644

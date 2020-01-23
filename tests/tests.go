package main

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
)

func main() {
	private_key, public_key := crypto.GenerateKeyPair()
	fmt.Println("Private Key" + private_key)
	fmt.Println("Public Key" + public_key)
	var election Election = Election{
		Type:         "Election\n",
		ElectionName: "Student Government Elections\n",
		Institution:  "West Chester University\n",
		Description:  "Spring Elections\n",
		Start:        "\n",
		End:          "\n",
		EmailDomain:  "wcupa.edu\n",
		Sender:       public_key,
	}

	election.Positions = []Position{
		Position{
			PositionId: "demfrmeororev",
			Name:       "Student Government President\n",
			Candidates: []Candidate{
				Candidate{
					Name:      "John Doe",
					Recipient: "test1",
				},
				Candidate{
					Name:      "Sarah Jennings",
					Recipient: "test2",
				},
				Candidate{
					Name:      "Maximus Footless",
					Recipient: "test3\n",
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

	fmt.Printf("%+v\n", election)
}

type Election struct {
	Type         string     `json:"type"`
	ElectionName string     `json:"electionName"` //Data Start
	Institution  string     `json:"institutionName"`
	Description  string     `json:"description"`
	Start        string     `json:"startDate"`
	End          string     `json:"endDate"`
	EmailDomain  string     `json:"emailDomain"`
	Positions    []Position `json:"positions"` //Data End
	Sender       string     `json:"sender"`
	Signature    string     `json:"id"`
}

type Position struct {
	PositionId string      `json:"id"`
	Name       string      `json:"displayName"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name      string `json:"name"`
	Recipient string `json:"key"`
}

type Registration struct {
	Type      string `json:"type"`
	Election  string `json:"election"` //Data Start
	Receiver  string `json:"receiver"` //Data End
	Sender    string `json:"sender"`
	Signature string `json:"signature"`
}

// valid votes have a corresponding registration transaction with the public key
type Vote struct {
	Type      string            `json:"type"`
	Election  string            `json:"election"` //Data Start
	Receiver  map[string]string `json:"receiver"` //Data End
	Sender    string            `json:"sender"`
	Signature string            `json:"signature"`
}

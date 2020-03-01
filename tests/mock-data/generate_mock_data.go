package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/joho/godotenv"
)

func main() {
	// timestamp := time.Now().Format(time.RFC1123)

	err := os.Chdir("../../build")
	if err != nil {
		fmt.Println(err)
		return
	}
	godotenv.Load()
	admin_private_key := os.Getenv("PRIVATE_KEY")
	admin_public_key := os.Getenv("PUBLIC_KEY")
	if admin_private_key == "" || admin_public_key == "" {
		fmt.Println("Generating new public/private key pair...")
		admin_private_key, admin_public_key = crypto.GenerateKeyPair()

		envfile := ".env"
		f, err := os.OpenFile(envfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			return
		}
		_, _ = f.WriteString("\nPRIVATE_KEY=" + admin_private_key + "\nPUBLIC_KEY=" + admin_public_key)
	}

	fmt.Println("Admin Private Key:\n" + admin_private_key + "\n")
	fmt.Println("Admin Public Key\n" + admin_public_key + "\n")

	start := time.Now().Format(time.RFC1123)
	end := time.Now().AddDate(0, 0, 200).Format(time.RFC1123) //200 days in  the future

	var election database.Election = database.Election{
		Type:         "Election",
		ElectionName: "Vote for Charity",
		Institution:  "Honestvote",
		Description:  "Whichever charities get the most votes, will be donated $50 each by Honestvote",
		Start:        start,
		End:          end,
		EmailDomain:  "^\\w{2}\\d{6}@wcupa\\.edu$",
		Sender:       database.PublicKey(admin_public_key),
	}

	election.Positions = []database.Position{
		database.Position{
			PositionId: "demfrmeororev",
			Name:       "Youth Charities",
			Candidates: []database.Candidate{
				database.Candidate{
					Name:      "Beverlys Birthdays",
					Recipient: "test1",
				},
				database.Candidate{
					Name:      "Brittanys Hope",
					Recipient: "test2",
				},
				database.Candidate{
					Name:      "Ronald McDonald House",
					Recipient: "test3",
				},
			},
		},
		database.Position{
			PositionId: "defmrfmrkmef",
			Name:       "Sustainability",
			Candidates: []database.Candidate{
				database.Candidate{
					Name:      "Art of Recycle",
					Recipient: "test4",
				},
				database.Candidate{
					Name:      "Safe Harbor of Chester County",
					Recipient: "test5",
				},
				database.Candidate{
					Name:      "Global Links",
					Recipient: "test6",
				},
			},
		},
	}

	jsonElection, _ := json.Marshal(election)

	encoded, err := election.Encode()
	if err != nil {
		fmt.Println(err)
		return
	}

	hash := crypto.CalculateHash(encoded)
	election.Signature, err = crypto.Sign([]byte(hash), admin_private_key)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	private_key, public_key := crypto.GenerateKeyPair()
	var registration database.AwaitingRegistration = database.AwaitingRegistration{
		Email:         "jacob@neubaum.com",
		FirstName:     "Jacob",
		LastName:      "Neubaum",
		DateOfBirth:   "3/9/1999",
		ElectionName:  election.Signature,
		ElectionAdmin: string(election.Sender),
		Sender:        database.PublicKey(public_key),
		SenderSig:     "",
		Code:          "",
		Timestamp:     "",
	}

	jsonRegistration, _ := json.Marshal(registration)

	var vote database.Vote = database.Vote{
		Type:     "Vote",
		Election: election.Signature,
		Receiver: []database.SelectedCandidate{
			database.SelectedCandidate{
				PositionId: "demfrmeororev",
				Recipient:  "test1",
			},
		},
		Sender: database.PublicKey(public_key),
	}

	encoded, err = vote.Encode()
	if err != nil {
		fmt.Println(err)
		return
	}
	hash = crypto.CalculateHash(encoded)
	vote.Signature, _ = crypto.Sign([]byte(hash), private_key)

	fmt.Println("Voter Private Key:\n" + private_key + "\n")
	fmt.Println("Voter Public Key\n" + public_key + "\n")

	jsonVote, _ := json.Marshal(vote)

	jsonElection, _ = json.MarshalIndent(election, "", "\t")
	jsonRegistration, _ = json.MarshalIndent(registration, "", "\t")
	jsonVote, _ = json.MarshalIndent(vote, "", "\t")

	// jsonData, _ := json.Marshal(jsonArray)

	filename := "../tests/mock-data/mock_data.sh"
	scriptname := "../scripts/deploy-local-chain.sh"

	// _ = ioutil.WriteFile(filename, jsonRegistration, 0644)
	// _ = ioutil.WriteFile(filename, jsonVote, 0644)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	script, err := os.Create(scriptname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, _ = io.WriteString(file, "echo \"Election Transaction:\"\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonElection)+"' http://localhost:7003/election\n\n\n\n")
	_, _ = io.WriteString(file, "echo \"Registration Transaction:\"\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonRegistration)+"' http://localhost:7003/election/test/register\n\n\n\n")
	_, _ = io.WriteString(file, "echo \"Vote Transaction:\"\n\ncurl --header \"Content-Type: application/json\" --request POST --data '"+
		string(jsonVote)+"' http://localhost:7003/election/test/vote\n\n\n\n")

	_, _ = io.WriteString(script, "go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name \""+election.Institution+"\n\n")

	_, _ = io.WriteString(script, "sleep 5\n\n")

	public_key, private_key = crypto.GenerateKeyPair()
	_, _ = io.WriteString(script, "go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key \""+private_key+"\" --public-key \""+public_key+"\" & \\\n\n")
	_, _ = io.WriteString(script, "sleep 10\n\n")
	public_key, private_key = crypto.GenerateKeyPair()
	_, _ = io.WriteString(script, "go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key \""+private_key+"\" --public-key \""+public_key+"\" & \\")

}

// os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644

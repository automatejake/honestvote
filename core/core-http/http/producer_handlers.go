package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
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
	if request.Domain == "" {
		logger.Println("client_handler.go", "PostRegisterHandler", "Domain field is empty")
		return
	} else if request.Institution == "" {
		logger.Println("client_handler.go", "PostRegisterHandler", "Institution name field is empty")
		return
	}

	message := []byte("requesting administrator privileges")
	valid_request, err := crypto.Verify(message, request.PublicKey, request.Signature)
	if err != nil {
		logger.Println("client_handler.go", "PostRegisterHandler", err)
		return
	}
	if !valid_request {
		logger.Println("client_handler.go", "PostRegisterHandler", "Invalid Signature for nomiation request")
		return
	}

	logger.Println("client_handler.go", "PostRegisterHandler", "Valid signature, nominating full node as an administrator")

	var election_options database.ElectionOptions = database.ElectionOptions{
		ElectionType:             "producer nomination", // (producer nomination | default), producer nomination election is a special election declared to elect a node
		ShowDataDuringElection:   "during",              // election results are shown before and after deciding on a nominee
		AllowedVotesPerVoter:     1,                     // each voter has one vote
		RequireVoteEveryPosition: true,                  // voter should be requred to cast a vote for every position
		MultipleVotesPerPosition: false,
	}
	var candidates []database.Candidate = []database.Candidate{
		database.Candidate{
			Recipient: "YES",
			Name:      "YES",
		},
		database.Candidate{
			Recipient: "NO",
			Name:      "NO",
		},
	}

	var positions []database.Position = []database.Position{
		database.Position{
			PositionId: request.PublicKey,
			Name:       "Should " + request.Institution + " be admitted into the network as a trusted election administrator and honestvote admin?",
			Candidates: candidates,
		},
	}
	var nomination database.Election = database.Election{
		Type:            "Election",
		ElectionName:    "Producer Nomination",
		Institution:     request.Institution,
		Description:     p2p.Self.Institution + " nominating " + request.Institution + " as a producer node",
		Start:           "",
		End:             "",
		EmailDomain:     request.Domain,
		ElectionOptions: election_options,
		Positions:       positions,
	}

	encoded, err := nomination.Encode()
	if err != nil {
		return
	}

	hash := crypto.CalculateHash(encoded)
	nomination.Signature, err = crypto.Sign([]byte(hash), p2p.PrivateKey)
	if err != nil {
		return
	}

	p2p.Enqueue(nomination)

	nodes := database.FindNodes()

	for _, node := range nodes {
		if node.Role == "producer" {

			registrant := database.Registration{
				Type:     "Registration",
				Election: nomination.Signature,
				Receiver: node.PublicKey,
				Sender:   p2p.PublicKey,
			}
			encoded, err := registrant.Encode()
			if err != nil {
				logger.Println("client_handler.go", "PostRegisterHandler", err)
			}

			hash := crypto.CalculateHash(encoded)
			registrant.Signature, err = crypto.Sign([]byte(hash), p2p.PrivateKey)
			if err != nil {
				logger.Println("client_handler.go", "PostRegisterHandler", err)
			}

			logger.Println("client_handler.go", "PostRegisterHandler", "Creating a registration")
			p2p.Enqueue(registrant)
		}
	}

}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)

	registrant, err := database.IsValidRegistrationCode(params["id"])
	if err != nil {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", err.Error())
		return
	}
	// err = database.CheckEmailVerification(registrant)
	// if err != nil {
	// 	logger.Println("producer_handlers.go", "VerifyEmailHandler()", err.Error())
	// 	w.Write([]byte("You have already registered."))
	// 	return
	// }

	switch p2p.REGISTRATION_TYPE {
	case "EXTERNAL_WHITELIST":
		if !registration.OnWhitelist(registrant.Email, p2p.Whitelist) {
			w.Write([]byte("You are not permitted to participate in this election.  Please talk to your election administrator if you think that this is a mistake."))
			return
		}
	case "DEFAULT_WHITELIST":

	}

	if params["verified"] == "true" {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", string(registrant.Sender)+" is registered to vote for "+registrant.ElectionName)
		if registration.VerifyStudent(registrant) {
			err := p2p.SendRegistrationTransaction(registrant)
			if err != nil {
				logger.Println("producer_handlers.go", "VerifyEmailHandler()", "Registration Transaction not sent correctly. "+err.Error())
			} else {
				w.Write([]byte("You have registered successfully!  Go back to the app to vote."))
			}

		}
	} else if params["verified"] == "false" {
		logger.Println("producer_handlers.go", "VerifyEmailHandler()", string(registrant.Sender)+" is not supposed to be registered to vote for "+registrant.ElectionName)
		database.RemoveRegistrationCode(registrant)
		w.Write([]byte("You indicated that you did not register to vote.  It may be possible that a malicious actor is attempting to register with your identity.  It is recommended to register as soon as possible."))
	} else {

	}

}

func GetEndpoint(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	endpoint, err := database.GetEndpoint()
	timestamp := time.Now().Format(time.RFC1123)
	payload := Payload{
		Timestamp: timestamp,
	}
	if err != nil {
		logger.Println("producer_handlers.go", "GetEndpoint()", err)
		payload.Status = "Bad Request"
	} else {
		payload.Status = "OK"
		payload.Data = endpoint

	}
	json.NewEncoder(w).Encode(payload)

}

package http

import (
	"net/http"

	"github.com/jneubaum/honestvote/core/core-websocket/websocket"
)

func HandleFullRoutes() {
	Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	//GETS
	Router.HandleFunc("/elections", GetElectionsHandler).Methods("GET")                     //good
	Router.HandleFunc("/election/{electionid}", GetElectionHandler).Methods("GET")          //good
	Router.HandleFunc("/election/{electionid}/votes", GetVotesHandler).Methods("GET")       //good
	Router.HandleFunc("/userpermissions/{publickey}", GetPermissionsHandler).Methods("GET") //good

	//POSTS
	Router.HandleFunc("/election/{electionid}/vote", PostVoteHandler).Methods("POST")
	Router.HandleFunc("/userpermissions/{publickey}/request", PostPermissionsHandler).Methods("POST")
	Router.HandleFunc("/elections", PostElectionsHandler).Methods("POST")

	//WEBSOCKET
	Router.HandleFunc("/websocket", websocket.WebsocketHandler) //good

	//this needs to be encrypted (send admin's public key and encrypted message containing email and public key)
	Router.HandleFunc("/registerElection/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
	Router.HandleFunc("/registerElection", RegisterHandler).Methods("POST")
}

func HandleProducerRoutes() {
	Router.HandleFunc("/verifyCode/code={id}&verified={verified}&email={email}", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/endpoint", GetEndpoint).Methods("GET") //good
	// Should not actually be here, only for testing.  Router.HandleFunc("/test/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
}

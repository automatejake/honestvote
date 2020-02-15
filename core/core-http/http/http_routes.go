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
	Router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("here there")

	}).Methods("GET")

	//POSTS
	Router.HandleFunc("/election/{electionid}/vote", PostVoteHandler).Methods("POST")
	Router.HandleFunc("/election/{electionid}/register", PostRegisterHandler).Methods("POST")
	Router.HandleFunc("/election", PostElectionHandler).Methods("POST")

	//WEBSOCKET
	Router.HandleFunc("/websocket", websocket.WebsocketHandler) //good

}

func HandleProducerRoutes() {
	Router.HandleFunc("/verifyCode/code={id}&verified={verified}", VerifyEmailHandler).Methods("GET")
	Router.HandleFunc("/endpoint", GetEndpoint).Methods("GET") //good
	// Should not actually be here, only for testing.  Router.HandleFunc("/test/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
}

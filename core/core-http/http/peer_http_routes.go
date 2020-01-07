package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-registration/registration"
	"github.com/jneubaum/honestvote/tests/logger"
)

var PeerRouter = mux.NewRouter()

func HandlePeerRoutes() {
	PeerRouter.HandleFunc("/verifyCode/code={id}&verified={verified}&email={email}", VerifyEmailHandler).Methods("GET")
	// Should not actually be here, only for testing.  PeerRouter.HandleFunc("/test/email={email}&public_key={public_key}&election={election}", RegisterHandler).Methods("GET")
	http.Handle("/", PeerRouter)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)
	fmt.Println("got here")
	public_key, election, valid := database.IsValidRegistrationCode(params["id"])

	if valid && params["verified"] == "true" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", public_key+" is registered to vote for "+election)

		// p2p.ReceiveTransaction(1)
	} else if params["verified"] == "false" {
		logger.Println("peer_http_routes.go", "VerifyEmailHandler()", public_key+" is not supposed to be registered to vote for "+election)
		// implement logic to allow email to register again
		registration.SendWarningEmail(params["email"], election)
	}

}

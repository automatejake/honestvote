package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
)

var PeerRouter = mux.NewRouter()

func HandlePeerRoutes() {
	PeerRouter.HandleFunc("/verifyCode/code={id}&verified={verified}", VerifyEmailHandler).Methods("GET")
	http.Handle("/", PeerRouter)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	params := mux.Vars(r)

	public_key, valid := database.IsValidRegistrationCode(params["id"])
	if valid && params["verified"] == "true" {
		logger.Println("peer_routes.go", "VerifyEmailHandler()", public_key+" is registered to vote")
		p2p.ReceiveVote(1)
	}

}

package http

import (
	"net/http"

<<<<<<< HEAD:core/core-http/http/peer_routes.go
	"github.com/jneubaum/honestvote/tests/logger"

	"github.com/jneubaum/honestvote/core/core-database/database"

=======
>>>>>>> 5955ad42bede2e39c4650199c44f464c535921e1:core/core-http/http/peer_http_routes.go
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

	public_key, election, valid := database.IsValidRegistrationCode(params["id"])
	if valid && params["verified"] == "true" {
<<<<<<< HEAD:core/core-http/http/peer_routes.go
		logger.Println("peer_routes.go", "VerifyEmailHandler()", public_key+" is registered to vote")
		//p2p.ReceiveVote(1)
=======
		logger.Println("peer_routes.go", "VerifyEmailHandler()", public_key+" is registered to vote for "+election)
		p2p.ReceiveVote(1)
	} else if params["verified"] == "false" {
		logger.Println("peer_routes.go", "VerifyEmailHandler()", public_key+" is not supposed to be registered to vote for "+election)

>>>>>>> 5955ad42bede2e39c4650199c44f464c535921e1:core/core-http/http/peer_http_routes.go
	}

}

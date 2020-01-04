package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

var PeerRouter = mux.NewRouter()

func HandlePeerRoutes() {
	PeerRouter.HandleFunc("/verifyCode", VerifyEmailHandler).Methods("GET")
	http.Handle("/", PeerRouter)
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	// //should be with database
	// if contains(Codes, "r.code request") {
	// 	//distributes vote to public key

	// }

}

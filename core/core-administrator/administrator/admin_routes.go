package administrator

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-database/database"
)

var AdminRouter = mux.NewRouter()

func HandleFullRoutes() {
	AdminRouter.HandleFunc("/become-peer", BecomePeer).Methods("GET")
	AdminRouter.HandleFunc("/become-peer/{auth-token}", BecomePeer).Methods("POST")
	http.Handle("/", AdminRouter)
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

}

func BecomePeer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	//if authentication token checks out, execute the following
	RequestPeerPrivileges(database.Node{})
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

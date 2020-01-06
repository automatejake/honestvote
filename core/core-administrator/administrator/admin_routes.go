package administrator

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/tests/logger"
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

/*
* When administrator requests to become a peer, they send the message to the peer that they are connected to
 */
func BecomePeer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	//if authentication token checks out, execute the following

	json_peer, err := json.Marshal(p2p.Self)
	if err != nil {
		logger.Println("admin_routes.go", "BecomePeer()", err.Error())
	}

	var message p2p.Message
	message.Message = "become peer"
	message.Data = json_peer

	json_message, err := json.Marshal(message)
	if err != nil {
		logger.Println("admin_routes.go", "BecomePeer()", err.Error())
	}

	_, err = p2p.Nodes[0].Write(json_message)
	if err != nil {
		logger.Println("admin_routes.go", "BecomePeer()", err.Error())
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

var Connections map[database.PublicKey]*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func BroadcastVote(vote database.Vote) {

	payload := Payload{
		Type:    "VOTES_ADD",
		Payload: vote,
	}

	jsonVote, err := json.Marshal(payload)
	if err != nil {
		logger.Println("broadcast.go", "WebsocketsHandler", err.Error())
	}

	for pubkey, conn := range Connections {
		if err := conn.WriteMessage(1, jsonVote); err != nil {
			conn.Close()
			fmt.Println("connection closed")
			delete(Connections, pubkey)
		}
		// fmt.Println("message sent: hello")

	}
}

func SendRegistration(registration database.Registration) {

	payload := Payload{
		Type:    "USER_CONFIRM_PERMISSION",
		Payload: registration,
	}

	jsonVote, err := json.Marshal(payload)
	if err != nil {
		logger.Println("broadcast.go", "WebsocketsHandler", err.Error())
	}

	publicKey := registration.Receiver

	if Connections[publicKey] == nil {
		return
	}

	if err := Connections[publicKey].WriteMessage(1, jsonVote); err != nil {
		Connections[publicKey].Close()
		delete(Connections, publicKey)
	}

}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	params := mux.Vars(r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Println("websocket_routes.go", "WebsocketsHandler", err.Error())
	}
	publicKey := database.PublicKey(params["publickey"])
	Connections[publicKey] = conn
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

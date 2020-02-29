package websocket

import (
	"encoding/json"
	"net/http"
	"fmt"
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

func MakeWebSocketMap() {
	Connections = make(map[database.PublicKey]*websocket.Conn)
}
func BroadcastVote(vote database.Vote) {

	payload := Payload{
		Type:    "VOTE_ADD",
		Payload: vote,
	}

	fmt.Printf("%+v\n",vote)

	logger.Println("","BroadcastVote()","Sending Vote...")
	jsonVote, err := json.Marshal(payload)
	if err != nil {
		logger.Println("broadcast.go", "WebsocketsHandler", err.Error())
	}


	for pubkey, conn := range Connections {
		if err := conn.WriteMessage(1, jsonVote); err != nil {
			conn.Close()
			delete(Connections, pubkey)
		}

	}
	
}

func SendRegistration(registration database.Registration) {
	logger.Println("","SendRegistration()","Registration is being sent")
	payload := Payload{
		Type:    "USER_CONFIRM_PERMISSION",
		Payload: registration.Election,
	}

	jsonVote, err := json.Marshal(payload)
	if err != nil {
		logger.Println("broadcast.go", "WebsocketsHandler", err.Error())
	}

	logger.Println("","",payload.Payload.(string))
	publicKey := registration.Receiver
	if Connections[publicKey] == nil {
		logger.Println("","","Public key does not exist in map")
		return
	}

	if err := Connections[publicKey].WriteMessage(1, jsonVote); err != nil {
		logger.Println("","SendRegistration()", "Error sending registration transaction" + err.Error())
		Connections[publicKey].Close()
		delete(Connections, publicKey)
	}

	logger.Println("","SendRegistration()","Registration sent successfully")
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	params := mux.Vars(r)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Println("websocket_routes.go", "WebsocketsHandler", err.Error())
	}
	publicKey := database.PublicKey(params["publickey"])
	Connections[publicKey] = conn

}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

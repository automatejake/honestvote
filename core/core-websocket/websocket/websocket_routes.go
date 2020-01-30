package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
)

var Connections []*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Broadcast(vote database.Vote) {

	payload := Payload{
		Type:    "VOTES_ADD",
		Payload: vote,
	}

	jsonVote, err := json.Marshal(payload)
	if err != nil {
		logger.Println("broadcast.go", "WebsocketsHandler", err.Error())
	}

	for i, conn := range Connections {
		if err := conn.WriteMessage(1, jsonVote); err != nil {
			conn.Close()
			fmt.Println("connection closed")
			Connections = append(Connections[:i], Connections[i+1:]...)
		}
		fmt.Println("message sent: hello")

	}
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Println("websocket_routes.go", "WebsocketsHandler", err.Error())
	}
	Connections = append(Connections, conn)
}

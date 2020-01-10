package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// 1
type longLatStruct struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *longLatStruct)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 2
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/longlat", longLatHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler)
	go echo()

	log.Fatal(http.ListenAndServe(":8844", router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func writer(coord *longLatStruct) {
	broadcast <- coord
}

func longLatHandler(w http.ResponseWriter, r *http.Request) {
	var coordinates longLatStruct
	if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&coordinates)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	clients[ws] = true
}

// 3
func echo() {
	for {
		val := <-broadcast
		latlong := fmt.Sprintf("%f %f %s", val.Lat, val.Long, "fre")
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

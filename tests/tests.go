// websockets.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConn(conn *websocket.Conn) {

	for {
		time.Sleep(time.Second * 2)
		if err := conn.WriteMessage(1, []byte("hello")); err != nil {
			fmt.Println("connection closed")
			conn.Close()
			return
		}
		fmt.Println("message sent: hello")
	}

}

func main() {
	// var m map[]string

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		go HandleConn(conn)

		for {

			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			fmt.Println(msgType)
			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				conn.Close()
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}

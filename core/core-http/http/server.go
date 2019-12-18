package http

import (
	"log"
	"net/http"
)

func CreateServer(port string) {
	HandleRoutes() // imported from routes

	log.Println("Listening...")
	http.ListenAndServe(":"+port, Router)
}

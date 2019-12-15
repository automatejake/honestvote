package corehttp

import (
	"log"
	"net/http"
)

func CreateServer() {
	HandleRoutes() // imported from routes

	log.Println("Listening...")
	http.ListenAndServe(":9001", Router)
}

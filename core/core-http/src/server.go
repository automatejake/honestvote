package corehttp

import (
	"log"
	"net/http"
)

func CreateServer() {
	HandleRoutes() // imported from routes

	log.Println("Listening...")
	http.ListenAndServe(":7001", Router)
}

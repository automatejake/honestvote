package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/tests/logger"
)

var Router = mux.NewRouter()

func CreateServer(port string, server_type string) {
	logger.Println("server.go", "main", "HTTP server running on port: "+port)

	if server_type == "producer" {
		HandleProducerRoutes()
		HandleFullRoutes()
	}

	if server_type == "full" {
		HandleFullRoutes() // imported from routes
	}

	http.Handle("/", Router)
	http.ListenAndServe(":"+port, Router)

}

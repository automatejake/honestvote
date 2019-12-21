package http

import (
	"net/http"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CreateServer(port string) {
	HandleRoutes() // imported from routes

	logger.Println("server.go", "main", "HTTP Service Running on port: "+port)
	http.ListenAndServe(":"+port, Router)
}

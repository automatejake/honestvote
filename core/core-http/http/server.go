package http

import (
	"net/http"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CreateServer(port string, server_type string) {
	if server_type == "peer" {
		HandlePeerRoutes()
		http.ListenAndServe(":"+port, PeerRouter)
	}

	if server_type == "full" {
		HandleFullRoutes() // imported from routes
		http.ListenAndServe(":"+port, FullRouter)
	}

	logger.Println("server.go", "main", "HTTP Service Running on port: "+port)

}

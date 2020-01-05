package http

import (
	"net/http"

	"github.com/jneubaum/honestvote/tests/logger"
)

func CreateServer(port string, server_type string) {
	logger.Println("server.go", "main", "HTTP server running on port: "+port)

	if server_type == "peer" {
		HandlePeerRoutes()
		http.ListenAndServe(":"+port, PeerRouter)
	}

	if server_type == "full" {
		HandleFullRoutes() // imported from routes
		http.ListenAndServe(":"+port, FullRouter)
	}

}

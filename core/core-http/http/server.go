package http

import (
	"net/http"
)

func CreateServer(port string) {
	HandleRoutes() // imported from routes

	logger.Println("main.go", "main", "HTTP Service Running on port: "+port)
	http.ListenAndServe(":"+port, Router)
}

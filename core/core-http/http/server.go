package http

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jneubaum/honestvote/tests/logger"
)

var Router = mux.NewRouter()
var HTTP_Port string

func CreateServer(port string, server_type string) {
	logger.Println("server.go", "main", "HTTP server running on port: "+port)

	if server_type == "producer" {
		HandleProducerRoutes()
		HandleFullRoutes()
	}
	if server_type == "full" {
		HandleFullRoutes() // imported from routes
	}

	HTTP_Port = port
	http.Handle("/", Router)
	http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(Router))

}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

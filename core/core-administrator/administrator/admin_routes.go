package administrator

import (
	"net/http"

	"github.com/gorilla/mux"
)

var AdminRouter = mux.NewRouter()

func HandleFullRoutes() {
	AdminRouter.HandleFunc("/become-peer", BecomePeer).Methods("GET")
	http.Handle("/", AdminRouter)
}

func BecomePeer(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

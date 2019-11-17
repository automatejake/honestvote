package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func VoterClient() {
	r := mux.NewRouter()
	r.HandleFunc("/getCandidates", GetCandidatesHandler)
	http.Handle("/", r)
}

func GetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

package main

import (
	"fmt"
	"os"

	"github.com/jneubaum/honestvote/core/core-discovery/discovery"
	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/joho/godotenv"
)

//defaults
var PEER_SERVICE string = ":9000"
var HTTP_SERVICE string = ":9001"

//this file will be responsible for deploying the app
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	// environmental variables override defaults
	PEER_SERVICE = ":" + os.Getenv("PEER_SERVICE")
	HTTP_SERVICE = ":" + os.Getenv("HTTP_SERVICE")

	// accept optional flags that override environmental variables

	// create http server
	go http.CreateServer(HTTP_SERVICE)

	// search for connections
	go discovery.FindPeer(os.Args[1])

	// accept incoming connections and handle p2p
	p2p.ListenConn(os.Args[1])

}

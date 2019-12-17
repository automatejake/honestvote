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
var ROLE string = "PEER" //options peer || full || registry
var DATABASE_PREFIX string = ""

//this file will be responsible for deploying the app
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	// environmental variables override defaults
	PEER_SERVICE = ":" + os.Getenv("PEER_SERVICE")
	HTTP_SERVICE = ":" + os.Getenv("HTTP_SERVICE")
	ROLE = os.Getenv("ROLE")
	DATABASE_PREFIX = os.Getenv("DATABASE_PREFIX")

	// accept optional flags that override environmental variables
	for index, element := range os.Args {
		switch element {
		case "--peer": //Set the default port for peer tcp service
			PEER_SERVICE = ":" + os.Args[index+1]
		case "--http": //Set the default port for http service
			HTTP_SERVICE = ":" + os.Args[index+1]
		case "--role": //Set the role of the node options PEER || FULL || REGISTRY
			ROLE = os.Args[index+1]
		case "--db-prefix": //Collection prefix (useful for starting up multiple nodes with same database)
			DATABASE_PREFIX = os.Args[index+1]
		}

	}

	// create http server for light clients to get information from
	if ROLE == "full" {
		go http.CreateServer(HTTP_SERVICE)
	}

	// search for connections
	if ROLE == "full" || ROLE == "peer" {
		go discovery.FindPeer(PEER_SERVICE)
	}

	p2p.ListenConn(PEER_SERVICE) // accept incoming connections and handle p2p

}

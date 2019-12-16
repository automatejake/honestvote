package main

import (
	"os"

	"github.com/jneubaum/honestvote/core/core-discovery/discovery"
	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

//this file will be responsible for deploying the app

func main() {

	// create http server
	go http.CreateServer()

	// search for connections
	go discovery.FindPeer(os.Args[1])

	// accept incoming connections and handle p2p
	p2p.ListenConn(os.Args[1])

}

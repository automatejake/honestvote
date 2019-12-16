package main

import (
	"os"

	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

//this file will be responsible for deploying the app

func main() {

	// create http server
	go http.CreateServer()

	go p2p.ListenConn(os.Args[1])

	p2p.PeerToPeer(os.Args[1])
}

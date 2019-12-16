package main

import (
	"github.com/jneubaum/honestvote/core/core-discovery/discovery"
	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

//this file will be responsible for deploying the app

func main() {

	// create http server
	go http.CreateServer()

	go p2p.ListenConn()

	discovery.FindPeer()
}

package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jneubaum/honestvote/tests/logger"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-discovery/discovery"
	"github.com/jneubaum/honestvote/core/core-http/http"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/core/core-registry/registry"
	"github.com/joho/godotenv"
)

//defaults
var TCP_SERVICE string = "7000"  //tcp service for peer to peer routes
var UDP_SERVICE string = "7001"  //udp service for node discovery
var HTTP_SERVICE string = "7002" //tcp service for light nodes to http routes

var ROLE string = "peer" //options peer || full || registry
var COLLECTION_PREFIX string = ""
var REGISTRY_IP string
var REGISTRY_PORT string = "7002"
var LOGGING bool = true
var PRIVATE_KEY, PUBLIC_KEY = crypto.KeyGen()

//this file will be responsible for deploying the app
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	// environmental variables override defaults
	if os.Getenv("TCP_SERVICE") != "" {
		TCP_SERVICE = os.Getenv("TCP_SERVICE")
	}
	if os.Getenv("UDP_SERVICE") != "" {
		UDP_SERVICE = os.Getenv("UDP_SERVICE")
	}
	if os.Getenv("HTTP_SERVICE") != "" {
		HTTP_SERVICE = os.Getenv("HTTP_SERVICE")
	}
	if os.Getenv("ROLE") != "" {
		ROLE = os.Getenv("ROLE")
	}
	if os.Getenv("COLLECTION_PREFIX") != "" {
		COLLECTION_PREFIX = os.Getenv("COLLECTION_PREFIX")
	}
	if os.Getenv("REGISTRY_IP") != "" {
		REGISTRY_IP = os.Getenv("REGISTRY_IP")
	}
	if os.Getenv("REGISTRY_PORT") != "" {
		REGISTRY_PORT = os.Getenv("REGISTRY_PORT")
	}

	//this domain is the default host to resolve traffic
	if REGISTRY_IP == "" {
		registry_ip, err := net.LookupIP("registry.honestvote.io")
		if err != nil {
			fmt.Println("Unknown host")
		} else {
			REGISTRY_IP = registry_ip[0].String()
		}
	}

	// accept optional flags that override environmental variables
	for index, element := range os.Args {
		switch element {
		case "--tcp": //Set the default port for peer tcp service
			TCP_SERVICE = os.Args[index+1]
		case "--udp":
			UDP_SERVICE = os.Args[index+1]
		case "--http": //Set the default port for http service
			HTTP_SERVICE = os.Args[index+1]
		case "--role": //Set the role of the node options PEER || FULL || REGISTRY
			ROLE = os.Args[index+1]
		case "--collection-prefix": //Collection prefix (useful for starting up multiple nodes with same database)
			COLLECTION_PREFIX = os.Args[index+1]
		case "--registry-host": //Sets the registry node
			REGISTRY_IP = os.Args[index+1]
		case "--registry-port": //Sets the registry node port
			REGISTRY_PORT = os.Args[index+1]
		}
	}

	database.CollectionPrefix = COLLECTION_PREFIX
	database.MongoDB = database.MongoConnect() // Connect to data store

	// if logging is turned on
	if LOGGING {
		logger.Logs = true
	}

	// create http server for light clients to get information from
	if ROLE == "full" {
		go http.CreateServer(HTTP_SERVICE)

	}

	// udp service that sends connected peers to other peers
	if ROLE == "registry" {
		go registry.ListenConnections(UDP_SERVICE)
	}

	// find peers to talk to from registry node
	if ROLE == "full" || ROLE == "peer" {
		logger.Println("main.go", "main", "Collection Prefix: "+COLLECTION_PREFIX)
		go discovery.FindPeer(REGISTRY_IP, REGISTRY_PORT, TCP_SERVICE)
	}

	// accept incoming connections and handle p2p
	p2p.ListenConn(TCP_SERVICE)

}

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

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
var TCP_PORT string = "7000"  //tcp PORT for peer to peer routes
var UDP_PORT string = "7001"  //udp PORT for node discovery
var HTTP_PORT string = "7002" //tcp PORT for light nodes to http routes

var ROLE string = "peer" //options peer || full || registry
var COLLECTION_PREFIX string = ""
var REGISTRY_IP string
var REGISTRY_PORT string = "7002"
var LOGGING bool = true

//this file will be responsible for deploying the app
func main() {
	p2p.PrivateKey, p2p.PublicKey = crypto.GenerateKeyPair()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	// environmental variables override defaults
	if os.Getenv("TCP_PORT") != "" {
		TCP_PORT = os.Getenv("TCP_PORT")
	}
	if os.Getenv("UDP_PORT") != "" {
		UDP_PORT = os.Getenv("UDP_PORT")
	}
	if os.Getenv("HTTP_PORT") != "" {
		HTTP_PORT = os.Getenv("HTTP_PORT")
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
	if os.Getenv("PRIVATE_KEY") != "" {
		p2p.PrivateKey = os.Getenv("PRIVATE_KEY")
	}
	if os.Getenv("PUBLIC_KEY") != "" {
		p2p.PublicKey = os.Getenv("PUBLIC_KEY")
	}
	if os.Getenv("PUBLIC_IP_ADDRESS") != "" {
		p2p.PublicIP = os.Getenv("PUBLIC_IP_ADDRESS")
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
		case "--tcp": //Set the default port for peer tcp PORT
			TCP_PORT = os.Args[index+1]
		case "--udp":
			UDP_PORT = os.Args[index+1]
		case "--http": //Set the default port for http PORT
			HTTP_PORT = os.Args[index+1]
		case "--role": //Set the role of the node options PEER || FULL || REGISTRY
			ROLE = os.Args[index+1]
		case "--collection-prefix": //Collection prefix (useful for starting up multiple nodes with same database)
			COLLECTION_PREFIX = os.Args[index+1]
		case "--registry-host": //Sets the registry node
			REGISTRY_IP = os.Args[index+1]
		case "--registry-port": //Sets the registry node port
			REGISTRY_PORT = os.Args[index+1]
		case "--private-key": //Sets the private key
			p2p.PrivateKey = os.Args[index+1]
		case "--public-key": //Sets the public key
			p2p.PublicKey = os.Args[index+1]
		case "--public-ip": //sets the public ip address
			p2p.PublicIP = os.Args[index+1]
		}
	}

	database.CollectionPrefix = COLLECTION_PREFIX
	database.MongoDB = database.MongoConnect() // Connect to data store

	port, _ := strconv.Atoi(TCP_PORT)
	identity := database.PublicKey(p2p.PublicKey)
	p2p.Self = database.Node{Port: port, Role: ROLE, Identity: identity}

	// if logging is turned on
	if LOGGING {
		logger.Logs = true
	}

	// udp PORT that sends connected peers to other peers
	if ROLE == "registry" {
		registry.ListenConnections(UDP_PORT)
	}

	// find peers to talk to from registry node
	if ROLE == "full" || ROLE == "peer" {
		logger.Println("main.go", "main", "Collection Prefix: "+COLLECTION_PREFIX)

		go http.CreateServer(HTTP_PORT, ROLE)
		go discovery.FindPeer(REGISTRY_IP, REGISTRY_PORT, TCP_PORT)

		// accept incoming connections and handle p2p
		p2p.ListenConn(TCP_PORT, ROLE)
	}

}

package main

import (
	"fmt"
	"log"
	"net"
	"os"

	coredb "github.com/jneubaum/honestvote.io/core/core-database/src"
	corehttp "github.com/jneubaum/honestvote.io/core/core-http/src"
	"github.com/joho/godotenv"
)

var Peers []coredb.Peer
var Blockchain []coredb.Block

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New TCP Connection")
	var buf [256]byte

	for {
		// conn.Read(buf[0:])
		msgLength, _ := conn.Read(buf[0:])

		// fmt.Println(string(buf[0:]))
		fmt.Println(string(buf[0:msgLength]))

	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	portString := ":" + (os.Getenv("PORT"))
	listener, err := net.Listen("tcp", portString)
	if err != nil {
		log.Fatal(err)
	}

	go corehttp.CreateServer()

	fmt.Println("Starting on port " + portString)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

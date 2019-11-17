package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jneubaum/honestvote.io/database/models"
	"github.com/jneubaum/honestvote.io/web"
	"github.com/joho/godotenv"
	//"bytes"
	//"encoding/gob"
	//"encoding/hex"
	//"io"
	//"github.com/joho/godotenv"
)

type Peer struct {
	Port   int
	Socket net.Conn
}

// type Candidate struct {
// 	Name      string
// 	PublicKey string
// 	Election  string
// }

var Blockchain []models.Block

// var Blockchain []Block

var Nodes = make(map[int]bool)

var Peers []Peer

var Candidates = []models.Candidate{
	models.Candidate{Name: "Jimmy", PublicKey: "0x54khfn4", Election: "Spring 2020"},
	models.Candidate{Name: "Janice", PublicKey: "0xflkh45n", Election: "Spring 2020"},
	models.Candidate{Name: "Larry", PublicKey: "0xij04ng3", Election: "Spring 2020"}}

func listenConn() {
	portString := ":" + (os.Getenv("PORT"))
	listen, err := net.Listen("tcp", portString)
	if err != nil {
		log.Fatal(err)
	}

	web.VoterClient()

	fmt.Println("Starting on port " + portString)

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New TCP Connection")
	var buf [256]byte

	for {
		// conn.Read(buf[0:])
		msgLength, _ := conn.Read(buf[0:])

		// fmt.Println(string(buf[0:]))
		fmt.Println(string(buf[0:msgLength]))

		if string(buf[0:14]) == "get candidates" {
			go getCandidates(conn)
		}

		// else if string(buf[0:18]) == "recieve candidates" {
		// 	fmt.Println("Recieving candidates!")
		// 	buffer := bytes.NewBuffer(buf[19:length])
		// 	tmpArray := new([]Candidate)
		// 	js := json.NewDecoder(buffer)
		// 	err := js.Decode(tmpArray)
		// 	if err == nil {
		// 		Candidates = append(Candidates, *tmpArray...)
		// 		fmt.Println(*tmpArray)
		// 	}
		// }
	}
}

func getCandidates(socket net.Conn) {
	buffer := new(bytes.Buffer)
	tmpArray := Candidates
	js := json.NewEncoder(buffer)
	err := js.Encode(tmpArray)
	if err == nil {
		socket.Write(buffer.Bytes())
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	// ignore, _ := strconv.Atoi(os.Getenv("PORT"))
	// Nodes[ignore] = true

	listenConn()

}

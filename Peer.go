package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

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

type Candidate struct {
	Name      string
	PublicKey string
	Election  string
}

var nodes = make(map[int]bool)

var Peers []Peer

var Candidates = []Candidate{
	Candidate{"Jimmy", "0x54khfn4", "Spring 2020"},
	Candidate{"Janice", "0xflkh45n", "Spring 2020"},
	Candidate{"Larry", "0xij04ng3", "Spring 2020"}}

func listenConn() {
	portString := ":" + (os.Getenv("PORT"))
	listen, err := net.Listen("tcp", portString)
	if err != nil {
		log.Fatal(err)
	}

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

	var buf [256]byte

	for {
		conn.Read(buf[0:])
		fmt.Println(string(buf[0:]))

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
	// nodes[ignore] = true

	listenConn()

	// for {
	// 	for port := 9000; port <= 9001; port++ {
	// 		if !nodes[port] {
	// 			fmt.Println("Checking...")
	// 			sPort := strconv.Itoa(port)
	// 			conn, _ := net.Dial("tcp", "127.0.0.1:"+sPort)
	// 			if conn != nil {
	// 				fmt.Println("Connected!")
	// 				Peers = append(Peers, Peer{port, conn})

	// 				nodes[port] = true

	// 				go handleConn(conn)
	// 			}
	// 		}
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
}

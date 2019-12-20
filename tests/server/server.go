package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// NOTE : Need to specify client address in WriteToUDP() function
	//        otherwise, you will get this error message
	//        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	// write message back to client
	// message := []byte("Hello UDP client!")

	// tmp_peers := database.FindPeers()
	// peer_json, _ := json.Marshal(tmp_peers)

	// var peers []database.Peer
	// _ = json.Unmarshal(peer_json, &peers)

	// fmt.Println(peers[0].IPAddress)

	tmp_peers := database.FindPeers()
	message, _ := json.Marshal(tmp_peers)

	n, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}

	var peers []database.Peer
	_ = json.Unmarshal(message, &peers)
	log.Println("Sent message: ", message, "\nbytes:", n)

}

func main() {
	hostName := "localhost"
	portNum := "6000"
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	database.MongoDB = database.MongoConnect()
	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 6000")

	defer ln.Close()

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln)
	}

}

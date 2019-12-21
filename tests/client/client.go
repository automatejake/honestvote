package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	hostName := "localhost"
	portNum := "6000"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	message := []byte("Hello UDP server!")

	_, err = conn.Write(message)

	if err != nil {
		log.Println(err)
	}

	// receive message from server
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)

	var peers []database.Peer
	_ = json.Unmarshal(buffer[0:n], &peers)

	fmt.Println("UDP Server : ", addr)
	fmt.Println("Received from UDP server : ", peers[0].IPAddress, "\nbytes: ", n)

	// tmp_peers := database.FindPeers()
	// peer_json, _ := json.Marshal(tmp_peers)

	// var peers []database.Peer
	// _ = json.Unmarshal(peer_json, &peers)

	// fmt.Println(peers[0].IPAddress)

}

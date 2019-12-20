package main

import (
	"encoding/json"
	"fmt"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

type P struct {
	X, Y, Z int
	Name    string
}

func main() {
	database.MongoDB = database.MongoConnect()
	// // database.ExistsInTable("127.0.0.1", "7002")
	// var network bytes.Buffer        // Stand-in for a network connection
	// enc := gob.NewEncoder(&network) // Will write to network.
	// dec := gob.NewDecoder(&network) // Will read from network.
	// // Encode (send) the value.
	// err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	// if err != nil {
	// 	log.Fatal("encode error:", err)
	// }

	// // HERE ARE YOUR BYTES!!!!
	// fmt.Println(network.Bytes())

	// // Decode (receive) the value.
	// var p P
	// err = dec.Decode(&p)
	// if err != nil {
	// 	log.Fatal("decode error:", err)
	// }
	// fmt.Printf("%q: {%d,%d,%d}\n", p.Name, p.X, p.Y, p.Z)

	// type PeerJSON struct {
	// 	Peers []database.Peer
	// }

	// type Peer struct {
	// 	IPAddress   string
	// 	Port        int
	// 	Role        string
	// 	Connections []Peer
	// }

	tmp_peers := database.FindPeers()
	peer_json, _ := json.Marshal(tmp_peers)

	var peers []database.Peer
	_ = json.Unmarshal(peer_json, &peers)

	fmt.Println(peers[0].IPAddress)

	// encoder := gob.NewEncoder(&tmp_peers_byte)
	// // Returns to node the list of nodes to speak with, IP Address and Port
	// for _, elem := range tmp_peers {
	// 	encoder.Encode(elem)
	// 	tmp_peers_bytes = append(tmp_peers_bytes, tmp_peer_bytes)
	// 	fmt.Println(elem.IPAddress, elem.Port)
	// }

	// encoder.Encode(tmp_peers_string)

	// var returned_peers_bytes [][]byte
	// decoder := gob.NewDecoder(&returned_peers_bytes)
	// var tmp_peers_peers []database.Peer
	// _ = decoder.Decode(&tmp_peers_string)

	// fmt.Println()

}

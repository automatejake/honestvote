package main

import (
	"log"
	"testing"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

type P struct {
	X, Y, Z int
	Name    string
}

func findpeers(b *testing.B) {
	database.MongoDB = database.MongoConnect()
	log.Println("Connected")
	database.ExistsInTable("127.0.0.1", 7002)
}

func main() {
	database.MongoDB = database.MongoConnect()
	log.Println("Connected")
	exclude_peer := database.Peer{IPAddress: "127.0.0.1", Port: 7004}
	peers := database.FindPeers(exclude_peer)
	for i := range peers {
		log.Println(peers[i])
	}

	// fmt.Println(testing.Benchmark(findpeers))

}

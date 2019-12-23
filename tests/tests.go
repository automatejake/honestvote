package main

import (
	"io"
	"log"
	"os"
)

// import "github.com/jneubaum/honestvote/tests/logging"

type P struct {
	X, Y, Z int
	Name    string
}

// func findpeers(b *testing.B) {
// 	database.MongoDB = database.MongoConnect()
// 	log.Println("Connected")
// 	database.ExistsInTable("127.0.0.1", 7002)
// }

// func main() {
// 	database.MongoDB = database.MongoConnect()
// 	log.Println("Connected")
// 	exclude_peer := database.Peer{IPAddress: "127.0.0.1", Port: 7004}
// 	peers := database.FindPeers(exclude_peer)
// 	for i := range peers {
// 		log.Println(peers[i])
// 	}

func main() {
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	// fmt.Println(testing.Benchmark(findpeers))

}

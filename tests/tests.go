package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"reflect"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
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
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey

	// privateKey, publicKey := crypto.KeyGen()

	encPriv, encPub := crypto.Encode(privateKey, publicKey)

	fmt.Println(encPriv)
	fmt.Println(encPub)

	priv2, pub2 := crypto.Decode(encPriv, encPub)

	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}

	msg := "hello"
	crypto.Sign1(msg, privateKey, *publicKey)
	fmt.Println(msg)

}

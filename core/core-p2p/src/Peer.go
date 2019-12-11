package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Peer struct {
	Port   int
	Socket net.Conn
}

type Candidate struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Election string `json:"election"`
	Votes    int32  `json:"votes"`
}

var nodes = make(map[int]bool)

var Peers []Peer

var mongoDB *mongo.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Loading ENV Failed")
	}

	mongoDB = mongoConnect()

	ignore, _ := strconv.Atoi(os.Getenv("PORT"))
	nodes[ignore] = true

	go listenConn()

	for {
		for port := 7000; port <= 7001; port++ {
			if !nodes[port] {
				fmt.Println("Checking...")
				sPort := strconv.Itoa(port)
				conn, _ := net.Dial("tcp", "127.0.0.1:"+sPort)
				if conn != nil {
					fmt.Println("Dial Successful!")
					tmpPeer := Peer{port, conn}
					Peers = append(Peers, tmpPeer)
					nodes[port] = true

					conn.Write([]byte("connect " + strconv.Itoa(ignore)))
					go handleConn(conn)
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func listenConn() {
	portString := ":" + os.Getenv("PORT")
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
		length, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {
			port, err := strconv.Atoi(string(buf[8:length]))

			if err == nil {
				nodes[port] = true
				tmpPeer := Peer{port, conn}
				Peers = append(Peers, tmpPeer)
			}
		} else if string(buf[0:8]) == "get data" {
			moveDocuments()
		}
	}
}

//Connect to MongoDB
func mongoConnect() *mongo.Client {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

//Get all the mongoDB data to send over to a full node or peer node that asked for it
func gatherMongoData(client *mongo.Client, filter bson.M) []Candidate {
	var Candidates []Candidate
	collection := client.Database("test_database").Collection("test_collection")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var candidate Candidate
		err = cur.Decode(&candidate)
		if err != nil {
			log.Fatal(err)
		}

		Candidates = append(Candidates, candidate)
	}

	return Candidates
}

//Send the data to the full/peer node
func moveDocuments() {
	MongoData := gatherMongoData(mongoDB, bson.M{})
	buffer := new(bytes.Buffer)
	tmpArray := MongoData
	js := json.NewEncoder(buffer)
	err := js.Encode(tmpArray)
	if err == nil {
		for _, socket := range Peers {
			fmt.Println("Sending documents.")
			socket.Socket.Write(append([]byte("recieve data "), buffer.Bytes()...))
		}
	}
}

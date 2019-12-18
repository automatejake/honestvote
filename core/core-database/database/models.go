package database

import "net"

// Multiple nodes can work on the same host using different collection prefixes
var CollectionPrefix string = ""

// Database is the same for all nodes even for a test net
var DatabaseName string = "honestvote"

// Elections
var ElectionHistory string = "election"

type Block struct {
	Index       int
	Timestamp   string
	Transaction Transaction
	Hash        string
	PrevHash    string
	Validator   string
}

type Transaction struct {
	Sender   string
	Vote     int
	Receiver string
}

type Election struct {
	Name             string `json:"name"`
	RegisteredVoters string `json:"registeredVoters"`
}

type Candidate struct {
	Name      string `json:"name"`
	PublicKey string `json:"key"`
	Election  string `json:"election"`
	Votes     int32  `json:"votes"`
}

type Peer struct {
	IPAddress         string
	Port              int
	Socket            net.Conn
	Role              string
	NumberConnections int
}

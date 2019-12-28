package database

import "net"

// Multiple nodes can work on the same host using different collection prefixes
var CollectionPrefix string = ""

// Database is the same for all nodes even for a test net
var DatabaseName string = "honestvote"

// Elections
var ElectionHistory string = "election"

//Peers on network
var Connections string = "connections"

type Block struct {
	Index       int
	Timestamp   string
	Transaction Transaction
	Hash        string
	PrevHash    string
	Validator   string
	Signature   string
	Valid       bool
	Port        int //TODO: Make Validator part of Peer struct to identify who proposed block
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

type Node struct {
	IPAddress   string
	Port        int
	Role        string
	Connections []Node
}

type TempNode struct {
	IPAddress string
	Port      int
	Role      string
	Socket    net.Conn
	//Validator -- used to identify who sent proposition block
}

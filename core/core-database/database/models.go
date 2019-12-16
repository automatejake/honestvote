package database

import "net"

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
	IPAddress string
	Port      int
	Socket    net.Conn
}

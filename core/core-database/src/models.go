package coredb

import "net"

type Block struct {
	Index     int
	Timestamp string
	Message   string
	Validator string
	PrevHash  string
	Hash      string
}

type Election struct {
	Name             string `json:"name"`
	RegisteredVoters string `json:"registeredVoters"`
}
type Candidate struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Election string `json:"election"`
	Votes    int32  `json:"votes"`
}
type Peer struct {
	Port   int
	Socket net.Conn
}

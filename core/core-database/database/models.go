package database

// Multiple nodes can work on the same host using different collection prefixes
var CollectionPrefix string = ""

// Database is the same for all nodes even for a test net
var DatabaseName string = "honestvote"

// Elections
var ElectionHistory string = "election"

// Nodes on network
var Connections string = "connections"

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int               `json:"index"`
	Timestamp   string            `json:"timestamp"`
	Transaction Transaction       `json:"transaction"`
	Hash        string            `json:"hash"`
	PrevHash    string            `json:"prevhash"`
	Signatures  map[string]string `json:"signatures"`
}

type Transaction struct {
	Sender   string
	Vote     int
	Type     string
	Election string
	Receiver string
}

type Election struct {
	Name             string     `json:"name"`
	RegisteredVoters string     `json:"registeredVoters"`
	Start            string     `json:"start"`
	End              string     `json:"end"`
	Positions        []Position `json:"positions"`
}

type Position struct {
	Name       string      `json:"name"`
	Candidates []Candidate `json:"candidates"`
}

type AwaitingRegistration struct {
	Election  string
	Code      string
	PublicKey string
	Timestamp string
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
	PublicKey   string
	Connections []Node
}

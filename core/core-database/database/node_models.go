package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "connections"  // Nodes on network

type PublicKey string

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int               `json:"index"`
	Timestamp   string            `json:"timestamp"`
	Type        string            `json:"type"`
	Transaction interface{}       `json:"transaction"` // instead of interface, should be transaction
	Hash        string            `json:"hash"`
	PrevHash    string            `json:"prevhash"`
	Signatures  map[string]string `json:"signatures"`
}

/*
*  types of transactions:
*	- becoming a consensus node
*	- declaring an election
*	- registering a student to vote
*	- casting a vote
 */

type Vote struct {
	Value        int            `json:"vote"`
	Registration string         `json:"registration"`
	Election     string         `json:"election"`
	Receiver     map[int]string `json:"receiver"`
	Sender       PublicKey      `json:"sender"`
	Signature    string         `json:"signature"`
}

type Election struct {
	Name           string     `json:"name"`
	Start          string     `json:"start"`
	End            string     `json:"end"`
	EligibleVoters int        `json:"registeredVoters"`
	Positions      []Position `json:"positions"`
	Sender         PublicKey  `json:"sender"`
	Signature      string     `json:"signature"`
}

type Node struct {
	Institution string
	IPAddress   string
	Port        int
	Role        string // peer | full | registry
	Identity    PublicKey
	Signature   string
}

func (node Node) VerifySignature() bool {
	if true {
		return true
	}
	return false
}

type Position struct {
	Name       string      `json:"name"`
	ID         int         `json:"id"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name      string `json:"name"`
	PublicKey string `json:"key"`
	Election  string `json:"election"`
	Votes     int    `json:"votes"`
}

type AwaitingRegistration struct {
	Email     string `json:"email"`
	Election  string `json:"election"`
	Code      string
	PublicKey string `json:"publicKey"`
	Timestamp string
}

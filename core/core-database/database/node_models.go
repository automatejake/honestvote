package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "node_list"    // Nodes on network

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
*  three types of transactions:
*	- declaring an election
*	- registering a student to vote
*	- casting a vote
 */

type Election struct {
	Type           string     `json:"type"`
	ElectionId     string     `json:"id"` //Data Start
	ElectionName   string     `json:"electionName"`
	Institution    string     `json:"institutionName"`
	Description    string     `json:"description"`
	Start          string     `json:"start"`
	End            string     `json:"end"`
	EmailDomain    string     `json:"emailDomain"`
	EligibleVoters int        `json:"registeredVoters"`
	Positions      []Position `json:"positions"` //Data End
	Sender         PublicKey  `json:"sender"`
	Signature      string     `json:"signature"`
}

type Registration struct {
	Type      string    `json:"type"`
	Election  string    `json:"election"` //Data Start
	Receiver  string    `json:"receiver"` //Data End
	Sender    PublicKey `json:"sender"`
	Signature string    `json:"signature"`
}

// valid votes are tied to registration transaction
type Vote struct {
	Type         string         `json:"type"`
	Registration string         `json:"registration"` //Data Start
	Election     string         `json:"election"`
	Receiver     map[int]string `json:"receiver"` //Data End
	Sender       PublicKey      `json:"sender"`
	Signature    string         `json:"signature"`
}

type Node struct {
	Institution string
	IPAddress   string
	Port        int
	Role        string // peer | full | registry
	PublicKey   PublicKey
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
	// Votes     int    `json:"votes"`
}

type AwaitingRegistration struct {
	Email     string `json:"email"`
	Election  string `json:"election"`
	Code      string
	PublicKey string `json:"publicKey"`
	Timestamp string
}

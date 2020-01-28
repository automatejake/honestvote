package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "node_list"    // Nodes on network

type PublicKey string

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int         `json:"index"`
	Timestamp   string      `json:"timestamp"`
	Transaction interface{} `json:"transaction"` // not  included in the hash
	MerkleRoot  string      `json:"merkleRoot"`
	Validator   string      `json:"validator"`
	PrevHash    string      `json:"prevhash"`
	Hash        string      `json:"hash"`
}

/*
*  three types of transactions:
*	- declaring an election
*	- registering a student to vote
*	- casting a vote
 */

type Registration struct {
	Type      string    `json:"type"`
	Election  string    `json:"election"` //Data Start
	Receiver  string    `json:"receiver"` //Data End
	Sender    PublicKey `json:"sender"`
	Signature string    `json:"signature"`
}

// valid votes have a corresponding registration transaction with the public key
type Vote struct {
	Type      string            `json:"type"`
	Election  string            `json:"electionId"` //Data Start
	Receiver  map[string]string `json:"recievers"`  //Data End
	Sender    PublicKey         `json:"sender"`
	Signature string            `json:"signature"`
}

type Election struct {
	Type         string     `json:"type"`
	ElectionName string     `json:"electionName"` //Data Start
	Institution  string     `json:"institutionName"`
	Description  string     `json:"description"`
	Start        string     `json:"startDate"`
	End          string     `json:"endDate"`
	EmailDomain  string     `json:"emailDomain"`
	Positions    []Position `json:"positions"` //Data End
	Sender       PublicKey  `json:"sender"`
	Signature    string     `json:"signature"`
}

type Position struct {
	PositionId string      `json:"id"`
	Name       string      `json:"positionName"`
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Name      string `json:"name"`
	Recipient string `json:"key"`
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

type AwaitingRegistration struct {
	Email     string `json:"email"`
	Election  string `json:"election"`
	Code      string
	PublicKey string `json:"publicKey"`
	Timestamp string
}

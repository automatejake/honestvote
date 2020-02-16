package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "node_list"    // Nodes on network

type PublicKey string

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int         `json:"index" bson:"index"`
	Timestamp   string      `json:"timestamp" bson:"timestamp"`
	Transaction interface{} `json:"transaction" bson:"transaction"` // not  included in the hash
	MerkleRoot  string      `json:"merkleRoot" bson:"merkleRoot"`
	Validator   string      `json:"validator" bson:"validator"`
	Signature   string      `json:"signature" bson:"signature"`
	PrevHash    string      `json:"prevhash" bson:"prevhash"`
	Hash        string      `json:"hash" bson:"hash"`
}

/*
*  three types of transactions:
*	- declaring an election
*	- registering a student to vote
*	- casting a vote
 */

type Registration struct {
	Type        string    `json:"type" bson:"type"`
	Election    string    `json:"electionId" bson:"electionId"` //Data Start
	Receiver    PublicKey `json:"receiver" bson:"receiver"`     //Data End
	RecieverSig string    `json:"recieverSig" bson:"senderSig"`
	Sender      PublicKey `json:"sender" bson:"sender"`
	Signature   string    `json:"signature" bson:"signature"`
}

type AwaitingRegistration struct {
	Email         string    `json:"emailAddress"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	DateOfBirth   string    `json:"dateOfBirth"`
	ElectionName  string    `json:"electionName"`
	ElectionAdmin string    `json:"electionAdmin"`
	Sender        PublicKey `json:"publicKey"`
	SenderSig     string    `json:"senderSig"`
	Code          string    `json:"code"`
	Timestamp     string    `json:"timestamp"`
}

// valid votes have a corresponding registration transaction with the public key
type Vote struct {
	Type      string              `json:"type" bson:"type"`
	Election  string              `json:"electionId" bson:"electionId"` //Data Start
	Receiver  []SelectedCandidate `json:"receivers" bson:"receivers"`   //Data End
	Sender    PublicKey           `json:"sender" bson:"sender"`
	Signature string              `json:"signature" bson:"signature"`
}

type SelectedCandidate struct {
	PositionId string `json:"id" bson:"id"`
	Recipient  string `json:"key" bson:"key"`
}

type Election struct {
	Type         string     `json:"type" bson:"type"`
	ElectionName string     `json:"electionName" bson:"electionName"` //Data Start
	Institution  string     `json:"institutionName" bson:"institutionName"`
	Description  string     `json:"description" bson:"description"`
	Start        string     `json:"startDate" bson:"startDate"`
	End          string     `json:"endDate" bson:"endDate"`
	EmailDomain  string     `json:"emailDomain" bson:"emailDomain"`
	Positions    []Position `json:"positions" bson:"positions"` //Data End
	Sender       PublicKey  `json:"sender" bson:"sender"`
	Signature    string     `json:"signature" bson:"signature"`
}

type Position struct {
	PositionId string      `json:"id" bson:"id"`
	Name       string      `json:"displayName" bson:"displayName"`
	Candidates []Candidate `json:"candidates" bson:"candidates"`
}

type Candidate struct {
	Name      string `json:"name" bson:"name"`
	Recipient string `json:"key" bson:"key"`
}

type Node struct {
	Institution string    `json:"institution" bson:"institution"`
	IPAddress   string    `json:"ipaddress" bson:"ipaddress"`
	Port        int       `json:"port" bson:"port"`
	Role        string    `json:"role" bson:"role"` // peer | full | registry
	PublicKey   PublicKey `json:"publickey" bson:"publickey"`
	Timestamp   string    `json:"timestamp" bson:"timestamp"`
	Signature   string    `json:"signature" bson:"signature"`
}

package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "node_list"    // Nodes on network

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index       int         `json:"index" bson:"index"`
	Timestamp   string      `json:"timestamp" bson:"timestamp"`
	Transaction interface{} `json:"transaction" bson:"transaction"` // DO WE WANT TO HOLD A REFERENCE OF ALL TRASNACTION HASH
	MerkleRoot  *MerkleTree `json:"merkleRoot" bson:"merkleRoot"`
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
	Type        string `json:"type" bson:"type"`
	Election    string `json:"electionId" bson:"electionId"` //Data Start
	Receiver    string `json:"receiver" bson:"receiver"`     //Data End
	RecieverSig string `json:"recieverSig" bson:"senderSig"`
	Sender      string `json:"sender" bson:"sender"`
	Signature   string `json:"signature" bson:"signature"`
}

type AdminSettings struct {
	NodeSettings      Node                      `json:"nodeSettings" bson:"nodeSettings"`
	WhiteListElection WhiteListElectionSettings `json:"whiteListElectionSettings" bson:"whiteListElectionSettings"`
}

type WhiteListElectionSettings struct {
	DatabaseDriver     string `json:"databaseDriver" bson:"databaseDriver"`
	DatabaseUser       string `json:"databaseUser" bson:"databaseUser"`
	DatabasePassword   string `json:"databasePassword" bson:"databasePassword"`
	DatabaseHost       string `json:"databaseHost" bson:"databaseHost"`
	DatabasePort       string `json:"databasePort" bson:"databasePort"`
	DatabaseName       string `json:"databaseName" bson:"databaseName"`
	TableName          string `json:"tableName" bson:"tableName"`
	EligibleVoterField string `json:"eligibleVoterField" bson:"eligibleVoterField"`
}

type AwaitingRegistration struct {
	Email         string `json:"emailAddress"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	DateOfBirth   string `json:"dateOfBirth"`
	ElectionName  string `json:"electionName"`
	ElectionAdmin string `json:"electionAdmin"`
	Sender        string `json:"publicKey"`
	SenderSig     string `json:"senderSig"`
	Code          string `json:"code"`
	Timestamp     string `json:"timestamp"`
}

// valid votes have a corresponding registration transaction with the public key
type Vote struct {
	Type      string              `json:"type" bson:"type"`
	Election  string              `json:"electionId" bson:"electionId"` //Data Start
	Receiver  []SelectedCandidate `json:"receivers" bson:"receivers"`   //Data End
	Sender    string              `json:"sender" bson:"sender"`
	Signature string              `json:"signature" bson:"signature"`
}

type SelectedCandidate struct {
	PositionId string `json:"positionId" bson:"positionId"`
	Recipient  string `json:"candidateName" bson:"candidateName"`
}

type Election struct {
	Type            string          `json:"type" bson:"type"`
	ElectionName    string          `json:"electionName" bson:"electionName"` //Data Start
	Institution     string          `json:"institutionName" bson:"institutionName"`
	Description     string          `json:"description" bson:"description"`
	Start           string          `json:"startDate" bson:"startDate"`
	End             string          `json:"endDate" bson:"endDate"`
	EmailDomain     string          `json:"emailDomain" bson:"emailDomain"`
	ElectionOptions ElectionOptions `json:"electionOptions" bson:"electionOptions"`
	Positions       []Position      `json:"positions" bson:"positions"` //Data End
	Sender          string          `json:"sender" bson:"sender"`
	Signature       string          `json:"signature" bson:"signature"`
}

type ElectionOptions struct {
	ElectionType             string `json:"electionType" bson:"electionType"`                         // (producer nomination | default), producer nomination election is a special election declared to elect a node
	ShowDataDuringElection   string `json:"showDataDuringElection" bson:"showDataDuringElection"`     // (during | after voting | after election end), logic handled by client
	AllowedVotesPerVoter     int    `json:"allowedVotesPerVoter" bson:"allowedVotesPerVoter"`         // amount of votes allowed to be cast by each voter, default 1
	MultipleVotesPerPosition bool   `json:"multipleVotesPerPosition" bson:"multipleVotesPerPosition"` // default false
}

type Position struct {
	PositionId string      `json:"id" bson:"id"`
	Name       string      `json:"displayName" bson:"displayName"`
	Candidates []Candidate `json:"candidates" bson:"candidates"`
}

type Candidate struct {
	Recipient string `json:"key" bson:"key"`
	Name      string `json:"name" bson:"name"`
}

type Node struct {
	Institution  string `json:"institution" bson:"institution"`
	IPAddress    string `json:"ipaddress" bson:"ipaddress"`
	Port         int    `json:"port" bson:"port"`
	Role         string `json:"role" bson:"role"` // peer | full | registry
	ConsensusPos int    `json:"consensuspos" bson:"consensuspos"`
	PublicKey    string `json:"publickey" bson:"publickey"`
	Timestamp    string `json:"timestamp" bson:"timestamp"`
	Signature    string `json:"signature" bson:"signature"`
}

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  string
}

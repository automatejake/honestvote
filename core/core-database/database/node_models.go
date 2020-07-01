package database

var CollectionPrefix string = ""        // Multiple nodes can work on the same host using different collection prefixes
var DatabaseName string = "honestvote"  // Database is the same for all nodes even for a test net
var ElectionHistory string = "election" // Elections
var Connections string = "node_list"    // Nodes on network

// Email registrants
var EmailRegistrants string = "email_registrants"

type Block struct {
	Index      int         `json:"index" bson:"index"`
	Timestamp  string      `json:"timestamp" bson:"timestamp"`
	MerkleRoot *MerkleTree `json:"merkleRoot" bson:"merkleRoot"`
	Validator  string      `json:"validator" bson:"validator"`
	Signature  string      `json:"signature" bson:"signature"`
	PrevHash   string      `json:"prevhash" bson:"prevhash"`
	Hash       string      `json:"hash" bson:"hash"`
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
	RecieverSig string `json:"recieverSig" bson:"recieverSig"`
	Sender      string `json:"sender" bson:"sender"`
	Signature   string `json:"signature" bson:"signature"`
	BlockIndex  int    `json:"blockIndex" bson:"blockIndex"`
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

// This is to prove the authenticity of the sender when requesting admin privelliges
type RequestAdminPrivileges struct {
	Message     []byte `json:"message" bson: "message"`
	Institution string `json:"institution" bson: "institution"`
	Domain      string `json:"domain" bson: "domain"`
	PublicKey   string `json:"publickey" bson: "publickey"`
	Signature   string `json:"signature" bson: "signature"`
}

// These are voters that are waiting to be registered
type AwaitingRegistration struct {
	Email         string `json:"emailAddress" bson: "emailAddress"`
	FirstName     string `json:"firstName" bson: "firstName"`
	LastName      string `json:"lastName" bson: "lastName"`
	DateOfBirth   string `json:"dateOfBirth" bson: "DateOfBirth"`
	ElectionName  string `json:"electionName" bson: "ElectionName"`
	ElectionAdmin string `json:"electionAdmin" bson: "ElectionAdmin"`
	Sender        string `json:"publicKey" bson: "Sender"`
	SenderSig     string `json:"senderSig" bson: "SenderSig"`
	Code          string `json:"code" bson: "code"`
	Timestamp     string `json:"timestamp" bson: "Timestamp"`
	Verified      string `json:"verified" bson: "verified"`
}

// valid votes have a corresponding registration transaction with the public key
type Vote struct {
	Type       string              `json:"type" bson:"type"`
	Election   string              `json:"electionId" bson:"electionId"` //Data Start
	Receiver   []SelectedCandidate `json:"receivers" bson:"receivers"`   //Data End
	Sender     string              `json:"sender" bson:"sender"`
	Signature  string              `json:"signature" bson:"signature"`
	BlockIndex int                 `json:"blockIndex" bson:"blockIndex"`
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
	BlockIndex      int             `json:"blockIndex" bson:"blockIndex"`
}

type ElectionOptions struct {
	ElectionType             string `json:"electionType" bson:"electionType"`                          // (producer nomination | default), producer nomination election is a special election declared to elect a node
	ShowDataDuringElection   string `json:"showDataDuringElection" bson:"showDataDuringElection"`      // (during | after voting | after election end), logic handled by client
	RequireVoteEveryPosition bool   `json:"requireVoteEveryPosition" bson: "requireVoteEveryPosition"` // should a voter be requred to cast a vote for every position, boolean
	AllowedVotesPerVoter     int    `json:"allowedVotesPerVoter" bson:"allowedVotesPerVoter"`          // amount of votes allowed to be cast by each voter, default is 1
	MultipleVotesPerPosition bool   `json:"multipleVotesPerPosition" bson:"multipleVotesPerPosition"`  // default false
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
	Hierarchy int //Higher number, closer to the root
	Left      *MerkleNode
	Right     *MerkleNode
	Hash      string
}

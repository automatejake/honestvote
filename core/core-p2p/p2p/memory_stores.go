package p2p

import (
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

// var Nodes = make(map[int]bool)

//Port is used for sending blocks so the Peers verifying know who do send it back to
// var TCP_PORT int

var PublicKey string
var PrivateKey string
var HTTP_Port string
var Email_Address string
var Email_Password string

var Nodes []net.Conn
var PreviousBlock database.Block
var ProposedBlock database.Block

var Self database.Node
var Whitelist database.WhiteListElectionSettings
var REGISTRATION_TYPE string

// Transactions quene holds transactions that are not yet ready to broadcast to the chain
var TransactionQueue []interface{}

//Transactions that will be added to a Block, clear when Block created
var TransactionsInBlock []string

// These are used to determine who is responsible for generating a block at any given time.  Found in broadcaster.go
var GenesisBlockTime int64 = 0
var Step int64 = 1
var ConsensusNodes int64 = 0

type Message struct {
	Message string `json:"message"`
	Data    []byte `json:"data"`
	Type    string `json:"type"`
}

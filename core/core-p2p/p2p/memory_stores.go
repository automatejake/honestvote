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

var Nodes []net.Conn
var PreviousBlock database.Block
var ProposedBlock database.Block

var Self database.Node

//Block queue starts to fill if ProposedBlock is not nil
var BlockQueue []database.Block

//Nested to pair answer with public_key/signature
var SignatureMap map[string]map[string]bool

type Message struct {
	Message string `json:"message"`
	Data    []byte `json:"data"`
	Type    string `json:"type"`
}

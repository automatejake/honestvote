package p2p

import (
	"net"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

// var Nodes = make(map[int]bool)

//Port is used for sending blocks so the Peers verifying know who do send it back to
var TCP_PORT int

//PrevHash is used to create a new block
var PrevHash = ""

//PrevIndex is used to create a new block
var PrevIndex = 0

var PublicKey string
var PrivateKey string

var Nodes []net.Conn
var ProposedBlock database.Block

var Self database.Node

//Block queue starts to fill if ProposedBlock is not nil
var BlockQueue []database.Block

//Nested to pair answer with public_key/signature
var SignatureMap map[string]map[string]bool
var PublicIP string

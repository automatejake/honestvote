package p2p

import "github.com/jneubaum/honestvote/core/core-database/database"

var Nodes = make(map[int]bool)
var Peers []database.TempPeer
var ProposedBlock database.Block

//Block queue starts to fill if ProposedBlock is not nil
var BlockQueue []database.Block

package p2p

import "github.com/jneubaum/honestvote/core/core-database/database"

// var Nodes = make(map[int]bool)

//Port is used for sending blocks so the Peers verifying know who do send it back to
var Port int
var Nodes []database.TempNode
var ProposedBlock database.Block

//Block queue starts to fill if ProposedBlock is not nil
var BlockQueue []database.Block

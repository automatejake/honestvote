package p2p

import "github.com/jneubaum/honestvote/core/core-database/database"

var Nodes = make(map[int]bool)
var Peers []database.TempPeer

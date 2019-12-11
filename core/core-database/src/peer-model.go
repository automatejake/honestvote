package coredb

import "net"

type Peer struct {
	Port   int
	Socket net.Conn
}

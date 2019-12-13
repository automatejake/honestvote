package main

import (
	"bytes"
	"encoding/json"
	"net"
	"strconv"

	coredb "github.com/jneubaum/honestvote.io/core/core-database/src"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte

	for {
		length, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		if string(buf[0:7]) == "connect" {
			port, err := strconv.Atoi(string(buf[8:length]))

			if err == nil {
				nodes[port] = true
				tmpPeer := coredb.Peer{
					IPAddress: "127.0.0.1",
					Port:      port,
					Socket:    conn,
				}
				Peers = append(Peers, tmpPeer)
			}
		} else if string(buf[0:12]) == "recieve data" {
			buffer := bytes.NewBuffer(buf[13:length])
			tmpArray := new([]coredb.Candidate)
			js := json.NewDecoder(buffer)
			err := js.Decode(tmpArray)
			if err == nil {
				coredb.UpdateMongo(coredb.MongoDB, *tmpArray)
			}
		} else if string(buf[0:8]) == "get data" {
			coredb.MoveDocuments(Peers)
		}
	}
}

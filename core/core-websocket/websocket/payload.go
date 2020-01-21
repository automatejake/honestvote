package websocket

import "github.com/jneubaum/honestvote/core/core-database/database"

type Payload struct {
	Type    string        `json:"type"`
	Payload database.Vote `json:"payload"`
}

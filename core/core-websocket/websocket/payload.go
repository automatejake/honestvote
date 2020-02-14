package websocket

type Payload struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

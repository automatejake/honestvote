package coredb

type Block struct {
	Index     int
	Timestamp string
	Message   string
	Validator string
	PrevHash  string
	Hash      string
}

type Candidate struct {
	Name      string `json:"name"`
	PublicKey string `json:"publickey"`
	Election  string `json:"election"`
}

type Election struct {
	Name             string `json:"name"`
	RegisteredVoters string `json:"registeredVoters"`
}

type Peer struct {
	IP     string
	Socket string
}

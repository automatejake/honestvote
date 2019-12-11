package coredb

type Candidate struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Election string `json:"election"`
	Votes    int32  `json:"votes"`
}

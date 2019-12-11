package coredb

type Block struct {
	Index     int
	Timestamp string
	Message   string
	Validator string
	PrevHash  string
	Hash      string
}

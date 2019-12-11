package coredb

type Election struct {
	Name             string `json:"name"`
	RegisteredVoters string `json:"registeredVoters"`
}

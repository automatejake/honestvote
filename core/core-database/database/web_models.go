package database

import (
	"fmt"
	"time"
)

type API_Type interface {
	ConvertInfo()
}

type CustomError struct {
	Time    time.Time
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.Time, e.Message)
}

type ElectionInfo struct {
	ElectionName string `json:"electionName"` //Data Start
	Institution  string `json:"institutionName"`
	Description  string `json:"description"`
	Start        string `json:"startDate"`
	End          string `json:"endDate"`
	Signature    string `json:"id"`
}

func (e Election) ConvertInfo() ElectionInfo {
	return ElectionInfo{
		ElectionName: e.ElectionName,
		Institution:  e.Institution,
		Description:  e.Description,
		Start:        e.Start,
		End:          e.End,
		Signature:    e.Signature,
	}
}

type VoteInfo struct {
	Sender    PublicKey         `json:"voterId"`
	Election  string            `json:"electionId"`
	Candidate map[string]string `json:"candidateId"`
	Signature string            `json:"signature"`
}

func (v Vote) ConvertInfo() VoteInfo {
	return VoteInfo{
		Sender:    v.Sender,
		Election:  v.Election,
		Candidate: v.Receiver,
		Signature: v.Signature,
	}
}

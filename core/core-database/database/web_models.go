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

package database

type API_Type interface {
	ConvertInfo()
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
	Sender    PublicKey `json:"voterId"`
	Election  string    `json:"electionId"`
	Candidate string    `json:"candidateId"`
	Signature string    `json:"signature"`
}

func (v Vote) ConvertInfo() VoteInfo {
	return VoteInfo{
		Sender:   v.Sender,
		Election: v.Election,
		// Candidate: v
		Signature: v.Signature,
	}
}

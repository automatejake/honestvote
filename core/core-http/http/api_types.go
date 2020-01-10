package http

type AppID string             // AppID is a generic string identifier used throughout the program
type VoterID AppID            // VoterID is an identifier for a Voter
type CandidateID AppID        // CandidateID is an identifier for a Candidate
type ElectionID AppID         // ElectionID is an identifier for an Election
type TicketID AppID           // TicketID is an identifier for a Ticket
type TicketEntryID AppID      // TicketEntryID is an identifier for a TicketEntry
type ElectionPositionID AppID // ElectionPositionID is an identifier for an ElectionPosition
type VotePriority int         // VotePriority is the priority for a vote. This is only used in non FPTP systems

type JSONData struct {
	Data [][]byte `json:"data"`
}

// Election is a given election
type Election struct {
	ID                  ElectionID      `json:"id"`
	ElectionName        string          `json:"displayName"`
	ElectionDescription string          `json:"term"`
	InstitutionName     string          `json:"institutionName"`
	StartDate           string          `json:"startDate"`
	EndDate             string          `json:"endDate"`
	EmailDomain         string          `json:"emailDomain"`
	Type                ElectionType    `json:"type"`
	TicketEntries       []TicketEntry   `json:"ballotEntries"`
	Options             ElectionOptions `json:"options,omitempty"`
}

// ElectionOptions are options that apply to a given Election
type ElectionOptions struct {
	CanHaveMultiTicket         bool `json:"canHaveMultiTicket,omitempty"`
	CandidateCanRunForMultiple bool `json:"candidateCanRunForMultiple,omitempty"`
}

// The ElectionType is an enumeration of different possible election types
type ElectionType int

// These constants represent different election types
const (
	FirstPastThePost ElectionType = iota
	InstantRunoff
	MultiRunoff
)

// TicketEntry is an entry in an election.
// For instance, this would define an entry such as "President", or even
// "President and Vice-President", running on the same ticket.
type TicketEntry struct {
	ID                       TicketEntryID        `json:"id"`
	DisplayName              string               `json:"displayName"`
	AllowedElectionPositions []ElectionPositionID `json:"allowedElectionPositions"`
	Tickets                  []TicketID           `json:"tickets"`
}

// Ticket is a specific ticket associated with a list of candidates
// running for this TicketEntry. For instance, you may have one ElectionPositionEntry
// for a President, and another ElectionPositionEntry for a Vice-President
type Ticket struct {
	ID                      TicketID                `json:"id"`
	ElectionPositionEntries []ElectionPositionEntry `json:"electionPositionEntries"`
	Votes                   []Vote                  `json:"votes"`
}

// ElectionPositionEntry is a single candidate and the office that they are running for.
type ElectionPositionEntry struct {
	CandidateID        CandidateID        `json:"candidateId"`
	ElectionPositionID ElectionPositionID `json:"electionPositionId"`
}

// Vote is a vote that can go toward a particular candidate
type Vote struct {
	VoterID      VoterID      `json:"voterId"`
	TicketID     TicketID     `json:"ticketId"`
	VotePriority VotePriority `json:"votePriority"` // used in rank based voting. for now always 1
}

// ElectionPosition is a particular position
type ElectionPosition struct {
	ID          ElectionPositionID `json:"id"`
	DisplayName string             `json:"displayName"`
}

// Voter is a user that is able to vote
type Voter struct {
	ID          VoterID          `json:"id"`
	Permissions VoterPermissions `json:"permissions"`
}

// VoterPermissions are the permissions granted to a given Voter user
type VoterPermissions struct {
	CanVote []AppID `json:"canVote"`
}

// Candidate is a candidate user that is publicly identified, and is able to run in an election
type Candidate struct {
	ID          CandidateID          `json:"id"`
	FullName    string               `json:"fullName"`
	Permissions CandidatePermissions `json:"permissions"`
}

//CandidatePermissions are the permissions granted to a given Candidate user
type CandidatePermissions struct {
	CanRun []AppID `json:"canRun"`
}

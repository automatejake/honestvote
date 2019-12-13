package corehttp

//Elections is example data for elections
var Elections []Election = []Election{
	{
		ID:          "0",
		DisplayName: "West Chester University Executive Board",
		Term:        "Spring 2020",
		Type:        FirstPastThePost,
		TicketEntries: []TicketEntry{
			{
				ID:                       "1",
				DisplayName:              "Presidential",
				AllowedElectionPositions: []ElectionPositionID{"11"},
				Tickets: []Ticket{
					{
						ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "13", ElectionPositionID: "11"}},
						Votes:                   []Vote{{VoterID: "5", VotePriority: 1}, {VoterID: "6", VotePriority: 1}, {VoterID: "7", VotePriority: 1}},
					}, {
						ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "14", ElectionPositionID: "11"}},
						Votes:                   []Vote{{VoterID: "8", VotePriority: 1}, {VoterID: "9", VotePriority: 1}, {VoterID: "10", VotePriority: 1}},
					},
				},
			}, {
				ID:                       "2",
				DisplayName:              "Secretorial",
				AllowedElectionPositions: []ElectionPositionID{"12"},
				Tickets: []Ticket{
					{
						ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "15", ElectionPositionID: "12"}},
						Votes:                   []Vote{{VoterID: "3", VotePriority: 1}, {VoterID: "4", VotePriority: 1}, {VoterID: "7", VotePriority: 1}},
					}, {
						ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "16", ElectionPositionID: "12"}},
						Votes:                   []Vote{{VoterID: "8", VotePriority: 1}, {VoterID: "9", VotePriority: 1}, {VoterID: "10", VotePriority: 1}},
					},
				},
			},
		},
		Options: ElectionOptions{
			CanHaveMultiTicket:         false,
			CandidateCanRunForMultiple: false,
			CandidateCanVote:           true,
			CandidateCanVoteForSelf:    false,
		},
	},
}

// Positions is example data for positions
var Positions []ElectionPosition = []ElectionPosition{
	{ID: "11", DisplayName: "President"},
	{ID: "12", DisplayName: "Secretary"},
}

// Voters is example data for voters
var Voters []Voter = []Voter{
	{ID: "3", Permissions: VoterPermissions{CanCreateElection: true, CanManageElection: []ElectionID{"0"}, CanVote: []AppID{"0"}}},
	{ID: "4", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "5", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "6", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "7", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "8", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "9", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
	{ID: "10", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"0"}}},
}

// Candidates is example data for candidates
var Candidates []Candidate = []Candidate{
	{ID: "13", DisplayName: "James Brennen", Permissions: CandidatePermissions{CanRun: []AppID{"0"}}},
	{ID: "14", DisplayName: "Mike Grimson", Permissions: CandidatePermissions{CanRun: []AppID{"0"}}},
	{ID: "15", DisplayName: "Alicia Michaels", Permissions: CandidatePermissions{CanRun: []AppID{"0"}}},
	{ID: "16", DisplayName: "Kelly Zimmerman", Permissions: CandidatePermissions{CanRun: []AppID{"0"}}},
}

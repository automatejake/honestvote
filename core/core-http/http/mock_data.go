package http

var Can JSONData

// <Full Node IP Address>:<Full Node Port>/candidates
// <Full Node IP Address>:<Full Node Port>/election?id=<ElectionId>
// <Full Node IP Address>:<Full Node Port>/voters
// <Full Node IP Address>:<Full Node Port>/positions
// <Full Node IP Address>:<Full Node Port>/tickets
//
// // Elections is example data for elections
// var Elections []ElectionInfo = []ElectionInfo{
// 	{
// 		ID:          "election1",
// 		DisplayName: "West Chester University Executive Board",
// 		Term:        "Spring 2020",
// 		Type:        FirstPastThePost,
// 	},
// }

// // Positions is example data for positions
// var Positions []ElectionPosition = []ElectionPosition{
// 	{ID: "position1", DisplayName: "President"},
// 	{ID: "position2", DisplayName: "Secretary"},
// }

// // Voters is example data for voters
// var Voters []Voter = []Voter{
// 	{ID: "voter1", Permissions: VoterPermissions{CanCreateElection: true, CanManageElection: []ElectionID{"election1"}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter2", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter3", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter4", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter5", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter6", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter7", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// 	{ID: "voter8", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
// }

// // Tickets is example data for tickets
// var Tickets []Ticket = []Ticket{
// 	{
// 		ID:                      "ticket1",
// 		ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate1", ElectionPositionID: "position1"}},
// 		Votes: []Vote{
// 			{VoterID: "voter3", TicketID: "ticket1", VotePriority: 1},
// 			{VoterID: "voter4", TicketID: "ticket1", VotePriority: 1},
// 			{VoterID: "voter5", TicketID: "ticket1", VotePriority: 1},
// 		},
// 	}, {
// 		ID:                      "ticket2",
// 		ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate2", ElectionPositionID: "position1"}},
// 		Votes: []Vote{
// 			{VoterID: "voter6", TicketID: "ticket2", VotePriority: 1},
// 			{VoterID: "voter7", TicketID: "ticket2", VotePriority: 1},
// 			{VoterID: "voter8", TicketID: "ticket2", VotePriority: 1},
// 		},
// 	}, {
// 		ID:                      "ticket3",
// 		ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate3", ElectionPositionID: "position2"}},
// 		Votes: []Vote{
// 			{VoterID: "voter1", TicketID: "ticket3", VotePriority: 1},
// 			{VoterID: "voter2", TicketID: "ticket3", VotePriority: 1},
// 			{VoterID: "voter5", TicketID: "ticket3", VotePriority: 1},
// 		},
// 	}, {
// 		ID:                      "ticket4",
// 		ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate4", ElectionPositionID: "position2"}},
// 		Votes: []Vote{
// 			{VoterID: "voter6", TicketID: "ticket4", VotePriority: 1},
// 			{VoterID: "voter7", TicketID: "ticket4", VotePriority: 1},
// 			{VoterID: "voter8", TicketID: "ticket4", VotePriority: 1},
// 		},
// 	},
// }

// // Candidates is example data for candidates
// var Candidates []Candidate = []Candidate{
// 	{ID: "candidate1", DisplayName: "James Brennen", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
// 	{ID: "candidate2", DisplayName: "Mike Grimson", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
// 	{ID: "candidate3", DisplayName: "Alicia Michaels", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
// 	{ID: "candidate4", DisplayName: "Kelly Zimmerman", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
// }

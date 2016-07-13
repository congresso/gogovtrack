package gogovtrack

// VotesResponse is
type VotesResponse struct {
	Meta  Meta    `json:"meta"`
	Votes []*Vote `json:"objects"`
}

// VoteResponse is
type VoteResponse Vote

// Vote is
type Vote struct {
	VoteID int `json:"id"`
}

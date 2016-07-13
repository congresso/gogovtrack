package gogovtrack

// VoteVotersResponse is
type VoteVotersResponse struct {
	Meta       Meta         `json:"meta"`
	VoteVoters []*VoteVoter `json:"objects"`
}

// VoteVoterResponse is
type VoteVoterResponse VoteVoter

// VoteVoter is
type VoteVoter struct {
	VoteVoterID int `json:"id"`
}

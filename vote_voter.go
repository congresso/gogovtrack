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
	Created           string       `json:"created"`
	VoteVoterID       int          `json:"id"`
	Option            Option       `json:"option"`
	Person            Person       `json:"person"`
	PersonRole        EmbeddedRole `json:"person_role"`
	Vote              Vote         `json:"vote"`
	VoterType         string       `json:"voter_type"`
	VoterTypeLabel    string       `json:"voter_type_label"`
	VoteviewExtraCode string       `json:"voteview_extra_code"`
}

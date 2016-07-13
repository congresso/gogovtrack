package gogovtrack

// CommitteesResponse is
type CommitteesResponse struct {
	Meta       Meta        `json:"meta"`
	Committees []Committee `json:"objects"`
}

// CommitteeResponse is
type CommitteeResponse Committee

// Committee is
type Committee struct {
	CommitteeID int `json:"id"`
}

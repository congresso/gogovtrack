package gogovtrack

// CommitteeMembersResponse is
type CommitteeMembersResponse struct {
	Meta             Meta               `json:"meta"`
	CommitteeMembers []*CommitteeMember `json:"objects"`
}

// CommitteeMemberResponse is
type CommitteeMemberResponse CommitteeMember

// CommitteeMember is
type CommitteeMember struct {
	CommitteeMemberID int `json:"id"`
}

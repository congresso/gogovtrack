package gogovtrack

import "testing"

func cm() *CommitteeMemberResource {
	return testAPI.Cm()
}

func TestCommitteeMemberFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := cm().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestCommitteeMemberAll(t *testing.T) {
	resp, err := cm().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.CommitteeMembers) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestCommitteeMemberOne(t *testing.T) {
	c, err := cm().One("207975")
	if err != nil {
		t.Error(err)
	}

	switch {
	case c.CommitteeMemberID != 207975:
		t.Errorf("Expected committee id to equal %d but got %d", 207975, c.CommitteeMemberID)
	}
}

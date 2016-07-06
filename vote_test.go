package gogovtrack

import "testing"

func vr() *VoteResource {
	return testAPI.Vr()
}

func TestVoteFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := vr().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestVoteAll(t *testing.T) {
	resp, err := vr().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Votes) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestVoteOne(t *testing.T) {
	p, err := vr().One("67617")
	if err != nil {
		t.Error(err)
	}

	switch {
	case p.VoteID != 67617:
		t.Errorf("Expected committee id to equal %d but got %d", 67617, p.VoteID)
	}
}

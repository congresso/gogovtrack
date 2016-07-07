package gogovtrack

import "testing"

func vv() *VoteVoterResource {
	return testAPI.Vvr()
}

func TestVoteVoterFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := vv().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestVoteVoterAll(t *testing.T) {
	resp, err := vv().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.VoteVoters) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestVoteVoterOne(t *testing.T) {
	p, err := vv().One("31369127")
	if err != nil {
		t.Error(err)
	}

	switch {
	case p.VoteVoterID != 31369127:
		t.Errorf("Expected committee id to equal %d but got %d", 31369127, p.VoteVoterID)
	}
}

package gogovtrack

import "testing"

func c() *CommitteeResource {
	return testAPI.Cr()
}

func TestCommitteeFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := c().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestCommitteeAll(t *testing.T) {
	resp, err := c().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Committees) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestCommitteeOne(t *testing.T) {
	committee, err := c().One("2650")
	if err != nil {
		t.Error(err)
	}

	switch {
	case committee.CommitteeID != 2650:
		t.Errorf("Expected committee id to equal %d but got %d", 2650, committee.CommitteeID)
	}
}

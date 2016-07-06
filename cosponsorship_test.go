package gogovtrack

import "testing"

func cs() *CosponsorshipResource {
	return testAPI.Cs()
}

func TestCosponsorFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := cs().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestCosponsorAll(t *testing.T) {
	resp, err := cs().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Cosponsorships) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestCosponsorOne(t *testing.T) {
	c, err := cs().One("1")
	if err != nil {
		t.Error(err)
	}

	switch {
	case c.CosponrshipID != 1:
		t.Errorf("Expected committee id to equal %d but got %d", 1, c.CosponrshipID)
	}
}

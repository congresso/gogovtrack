package gogovtrack

import "testing"

func pr() *PersonResource {
	return testAPI.Pr()
}

func TestPersonFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := pr().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestPersonAll(t *testing.T) {
	resp, err := pr().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Persons) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestPersonOne(t *testing.T) {
	p, err := pr().One("400326")
	if err != nil {
		t.Error(err)
	}

	switch {
	case p.PersonID != 400326:
		t.Errorf("Expected committee id to equal %d but got %d", 400326, p.PersonID)
	}
}

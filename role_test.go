package gogovtrack

import "testing"

func rr() *RoleResource {
	return testAPI.Rr()
}

func TestRoleFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	cr := rr().Filter(query)

	if cr.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestRoleAll(t *testing.T) {
	resp, err := rr().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Roles) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestRoleOne(t *testing.T) {
	p, err := rr().One("1")
	if err != nil {
		t.Error(err)
	}

	switch {
	case p.RoleID != 1:
		t.Errorf("Expected committee id to equal %d but got %d", 1, p.RoleID)
	}
}

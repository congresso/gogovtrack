package gogovtrack

import "testing"

func b() *BillResource {
	return testAPI.Br()
}

func TestBillFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	br := b().Filter(query)

	if br.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestBillll(t *testing.T) {
	resp, err := b().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Bills) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestBillOne(t *testing.T) {
	bill, err := b().One("127129")
	if err != nil {
		t.Error(err)
	}

	switch {
	case bill.BillResolutionType != "bill":
		t.Errorf("Expecting bill type to equal %s but got %s", "bill", bill.BillResolutionType)
	}
}

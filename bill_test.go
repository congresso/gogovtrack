package gogovtrack

import (
	"net/http"
	"testing"
)

func b() *BillResource {
	a := NewAPI(http.DefaultClient, baseURL)

	return a.Br()
}

func TestFilter(t *testing.T) {
	query := Q{"id": "1", "limit": "200"}

	br := b().Filter(query)

	if br.Filters == nil {
		t.Error("Expected bill resource filters to not be nil")
	}
}

func TestAll(t *testing.T) {
	resp, err := b().All()
	if err != nil {
		t.Error(err)
	}

	if len(resp.Bills) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

func TestOne(t *testing.T) {
	bill, err := b().One("127129")
	if err != nil {
		t.Error(err)
	}

	switch {
	case bill.BillResolutionType != "bill":
		t.Errorf("Expecting bill type to equal %s but got %s", "bill", bill.BillResolutionType)
	}
}

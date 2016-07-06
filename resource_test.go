package gogovtrack

import (
	"net/http"
	"reflect"
	"testing"
)

func r() *Resource {
	a := &API{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}

	return a.R("bill")
}

func TestFindResource(t *testing.T) {
	r := r().Filter(Q{"one": "two", "three": "fourt"})

	if reflect.TypeOf(r) != reflect.TypeOf(new(Resource)) {
		t.Errorf("Expecting r to be type of %v but got %v", reflect.TypeOf(new(Resource)), reflect.TypeOf(r))
	}
}

func TestBuildFilterQuery(t *testing.T) {
	expected := "?id=1&limit=200"

	query := Q{"id": "1", "limit": "200"}

	result := r().Filter(query).buildFilterQuery()

	if result != expected {
		t.Errorf("Expected query to equal %s but got %s", expected, result)
	}
}

func TestGetResource(t *testing.T) {
	r, err := r().All()
	if err != nil {
		t.Error(err)
	}

	if len(r.Objects) < 1 {
		t.Error("Expected at least 1 bill to be returned")
	}
}

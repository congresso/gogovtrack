package gogovtrack

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewAPI(t *testing.T) {
	a := NewAPI(http.DefaultClient, baseURL)

	if reflect.TypeOf(a) != reflect.TypeOf(new(API)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(API)), reflect.TypeOf(a))
	}
}

func TestBuildFilterQuery(t *testing.T) {
	a := NewAPI(http.DefaultClient, baseURL)

	expected := "?id=1&limit=200"

	query := Q{"id": "1", "limit": "200"}

	result := a.buildFilterQuery(query)

	if result != expected {
		t.Errorf("Expected query to equal %s but got %s", expected, result)
	}
}

func TestBr(t *testing.T) {
	a := NewAPI(http.DefaultClient, baseURL)

	br := a.Br()

	if reflect.TypeOf(br) != reflect.TypeOf(new(BillResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(BillResource)), reflect.TypeOf(br))
	}
}

func TestCo(t *testing.T) {
	a := NewAPI(http.DefaultClient, baseURL)

	co := a.Co()

	if reflect.TypeOf(co) != reflect.TypeOf(new(CommitteeResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(CommitteeResource)), reflect.TypeOf(co))
	}
}

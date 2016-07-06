package gogovtrack

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

var (
	testAPI *API
)

func TestMain(m *testing.M) {
	testAPI = NewAPI(http.DefaultClient, BaseURL)

	os.Exit(m.Run())
}

func TestNewAPI(t *testing.T) {
	a := NewAPI(http.DefaultClient, BaseURL)

	if reflect.TypeOf(a) != reflect.TypeOf(new(API)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(API)), reflect.TypeOf(a))
	}
}

func TestBuildFilterQuery(t *testing.T) {
	expected1 := "?id=1&limit=200"
	expected2 := "?limit=200&id=1"

	query := Q{"id": "1", "limit": "200"}

	result := testAPI.buildFilterQuery(query)

	if result != expected1 {
		if result != expected2 {
			t.Errorf("Expected query to equal %s or %sbut got %s", expected1, expected2, result)
		}
	}
}

func TestBr(t *testing.T) {
	br := testAPI.Br()

	if reflect.TypeOf(br) != reflect.TypeOf(new(BillResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(BillResource)), reflect.TypeOf(br))
	}
}

func TestCo(t *testing.T) {
	co := testAPI.Co()

	if reflect.TypeOf(co) != reflect.TypeOf(new(CommitteeResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(CommitteeResource)), reflect.TypeOf(co))
	}
}

func TestCs(t *testing.T) {
	co := testAPI.Cs()

	if reflect.TypeOf(co) != reflect.TypeOf(new(CosponsorshipResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(CommitteeResource)), reflect.TypeOf(co))
	}
}

func TestPr(t *testing.T) {
	co := testAPI.Pr()

	if reflect.TypeOf(co) != reflect.TypeOf(new(PersonResource)) {
		t.Errorf("Expecting api to be type of %v but got %v", reflect.TypeOf(new(PersonResource)), reflect.TypeOf(co))
	}
}

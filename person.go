package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// PersonsResponse is the json response for getting all persons
type PersonsResponse struct {
	Meta    Meta      `json:"meta"`
	Persons []*Person `json:"objects"`
}

// PersonResponse is the json response for getting a person
type PersonResponse Person

// Person holds the properties of a person
type Person struct {
	PersonID int `json:"id"`
}

// PersonResource is the resource for the person api
type PersonResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (p *PersonResource) Filter(query Q) *PersonResource {
	p.Filters = query

	return p
}

// All prepares and sends the http request to get all persons
func (p *PersonResource) All() (*PersonsResponse, error) {
	filters := p.api.buildFilterQuery(p.Filters)

	url := fmt.Sprintf("%s/%s%s", p.api.baseURL, p.Name, filters)

	resp, err := p.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var pe *PersonsResponse
	if err := json.NewDecoder(resp.Body).Decode(&pe); err != nil {
		return nil, err
	}

	return pe, nil
}

// One prepares and sends the http request to get a person
func (p *PersonResource) One(id string) (*PersonResponse, error) {
	filters := p.api.buildFilterQuery(p.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", p.api.baseURL, p.Name, id, filters)

	resp, err := p.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var pe *PersonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pe); err != nil {
		return nil, err
	}

	return pe, nil
}

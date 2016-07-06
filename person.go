package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// PersonsResponse is
type PersonsResponse struct {
	Meta    Meta      `json:"meta"`
	Persons []*Person `json:"objects"`
}

// PersonResponse is
type PersonResponse Person

// Person is
type Person struct {
	PersonID int `json:"id"`
}

// PersonResource is
type PersonResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (p *PersonResource) Filter(query Q) *PersonResource {
	p.Filters = query

	return p
}

// All is
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

// One is
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

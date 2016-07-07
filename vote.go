package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// VotesResponse is the json response for getting all votes
type VotesResponse struct {
	Meta  Meta    `json:"meta"`
	Votes []*Vote `json:"objects"`
}

// VoteResponse is the json response for getting a vote
type VoteResponse Vote

// Vote holds the properties of a vote
type Vote struct {
	VoteID int `json:"id"`
}

// VoteResource is the resource for the vote api
type VoteResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (v *VoteResource) Filter(query Q) *VoteResource {
	v.Filters = query

	return v
}

// All prepares and sends the http request to get all votes
func (v *VoteResource) All() (*VotesResponse, error) {
	filters := v.api.buildFilterQuery(v.Filters)

	url := fmt.Sprintf("%s/%s%s", v.api.baseURL, v.Name, filters)

	resp, err := v.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var ve *VotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&ve); err != nil {
		return nil, err
	}

	return ve, nil
}

// One prepares and sends the http request to get a vote
func (v *VoteResource) One(id string) (*VoteResponse, error) {
	filters := v.api.buildFilterQuery(v.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", v.api.baseURL, v.Name, id, filters)

	resp, err := v.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var ve *VoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&ve); err != nil {
		return nil, err
	}

	return ve, nil
}

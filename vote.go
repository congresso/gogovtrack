package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// VotesResponse is
type VotesResponse struct {
	Meta  Meta    `json:"meta"`
	Votes []*Vote `json:"objects"`
}

// VoteResponse is
type VoteResponse Vote

// Vote is
type Vote struct {
	VoteID int `json:"id"`
}

// VoteResource is
type VoteResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (v *VoteResource) Filter(query Q) *VoteResource {
	v.Filters = query

	return v
}

// All is
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

// One is
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

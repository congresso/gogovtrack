package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// VoteVotersResponse is
type VoteVotersResponse struct {
	Meta       Meta         `json:"meta"`
	VoteVoters []*VoteVoter `json:"objects"`
}

// VoteVoterResponse is
type VoteVoterResponse VoteVoter

// VoteVoter is
type VoteVoter struct {
	VoteVoterID int `json:"id"`
}

// VoteVoterResource is
type VoteVoterResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (v *VoteVoterResource) Filter(query Q) *VoteVoterResource {
	v.Filters = query

	return v
}

// All is
func (v *VoteVoterResource) All() (*VoteVotersResponse, error) {
	filters := v.api.buildFilterQuery(v.Filters)

	url := fmt.Sprintf("%s/%s%s", v.api.baseURL, v.Name, filters)

	resp, err := v.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var ve *VoteVotersResponse
	if err := json.NewDecoder(resp.Body).Decode(&ve); err != nil {
		return nil, err
	}

	return ve, nil
}

// One is
func (v *VoteVoterResource) One(id string) (*VoteVoterResponse, error) {
	filters := v.api.buildFilterQuery(v.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", v.api.baseURL, v.Name, id, filters)

	resp, err := v.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var ve *VoteVoterResponse
	if err := json.NewDecoder(resp.Body).Decode(&ve); err != nil {
		return nil, err
	}

	return ve, nil
}

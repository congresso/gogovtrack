package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// VoteVotersResponse is the json response for getting all vote voters
type VoteVotersResponse struct {
	Meta       Meta         `json:"meta"`
	VoteVoters []*VoteVoter `json:"objects"`
}

// VoteVoterResponse is the json response for getting a vote voter
type VoteVoterResponse VoteVoter

// VoteVoter holds the properties of a vote voter
type VoteVoter struct {
	VoteVoterID int `json:"id"`
}

// VoteVoterResource is the resource for the vote voter api
type VoteVoterResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (v *VoteVoterResource) Filter(query Q) *VoteVoterResource {
	v.Filters = query

	return v
}

// All prepares and sends the http request to get all vote voters
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

// One prepares and sends the http request to get a vote voter
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

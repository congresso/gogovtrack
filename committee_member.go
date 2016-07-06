package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CommitteeMembersResponse is
type CommitteeMembersResponse struct {
	Meta             Meta               `json:"meta"`
	CommitteeMembers []*CommitteeMember `json:"objects"`
}

// CommitteeMemberResponse is
type CommitteeMemberResponse CommitteeMember

// CommitteeMember is
type CommitteeMember struct {
	CommitteeMemberID int `json:"id"`
}

// CommitteeMemberResource is
type CommitteeMemberResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (c *CommitteeMemberResource) Filter(query Q) *CommitteeMemberResource {
	c.Filters = query

	return c
}

// All is
func (c *CommitteeMemberResource) All() (*CommitteeMembersResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s%s", c.api.baseURL, c.Name, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var cm *CommitteeMembersResponse
	if err := json.NewDecoder(resp.Body).Decode(&cm); err != nil {
		return nil, err
	}

	return cm, nil
}

// One is
func (c *CommitteeMemberResource) One(id string) (*CommitteeMemberResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", c.api.baseURL, c.Name, id, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var cm *CommitteeMemberResponse
	if err := json.NewDecoder(resp.Body).Decode(&cm); err != nil {
		return nil, err
	}

	return cm, nil
}

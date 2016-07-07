package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CommitteeMembersResponse is the json response for getting all committee members
type CommitteeMembersResponse struct {
	Meta             Meta               `json:"meta"`
	CommitteeMembers []*CommitteeMember `json:"objects"`
}

// CommitteeMemberResponse is the json response for getting a committee member
type CommitteeMemberResponse CommitteeMember

// CommitteeMember is struct that holds the properties of a committee member
type CommitteeMember struct {
	CommitteeMemberID int `json:"id"`
}

// CommitteeMemberResource is the resource for the committee member api
type CommitteeMemberResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (c *CommitteeMemberResource) Filter(query Q) *CommitteeMemberResource {
	c.Filters = query

	return c
}

// All prepares and sends the http request to get all the committee members
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

// One prepares and sends the http request to get a committee member by id
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

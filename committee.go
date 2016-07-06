package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CommitteesResponse is
type CommitteesResponse struct {
	Meta       Meta        `json:"meta"`
	Committees []Committee `json:"objects"`
}

// CommitteeResponse is
type CommitteeResponse Bill

// Committee is
type Committee struct {
	Abrev string `json:"abrev"`
	Code  string `json:"code"`
	ID    int    `json:"id"`
}

// CommitteeResource is
type CommitteeResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (c *CommitteeResource) Filter(query Q) *CommitteeResource {
	c.Filters = query

	return c
}

// All is
func (c *CommitteeResource) All() (*CommitteesResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s%s", c.api.baseURL, c.Name, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var cr *CommitteesResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return nil, err
	}

	return cr, nil
}

// One is
func (c *CommitteeResource) One(id string) (*CommitteeResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", c.api.baseURL, c.Name, id, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var cr *CommitteeResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return nil, err
	}

	return cr, nil
}

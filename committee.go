package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CommitteesResponse is the json response for getting all the committees
type CommitteesResponse struct {
	Meta       Meta        `json:"meta"`
	Committees []Committee `json:"objects"`
}

// CommitteeResponse is the json response for getting a committee
type CommitteeResponse Committee

// Committee holds the properties of a committee
type Committee struct {
	Abrev       string `json:"abrev"`
	Code        string `json:"code"`
	CommitteeID int    `json:"id"`
}

// CommitteeResource is the resource for the committee api
type CommitteeResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (c *CommitteeResource) Filter(query Q) *CommitteeResource {
	c.Filters = query

	return c
}

// All prepares and sends the http request to get all committees
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

// One prepares and sends the http request to get a committee
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

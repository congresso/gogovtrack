package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CosponsorshipsResponse is the json response for getting all cosponsorships
type CosponsorshipsResponse struct {
	Meta           Meta             `json:"meta"`
	Cosponsorships []*Cosponsorship `json:"objects"`
}

// CosponsorshipResponse is the json response for getting a cosponsorship
type CosponsorshipResponse Cosponsorship

// Cosponsorship is struct that holds the properties of a cosponsorship
type Cosponsorship struct {
	CosponrshipID int `json:"id"`
}

// CosponsorshipResource is the resource for the cosponsorship api
type CosponsorshipResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (c *CosponsorshipResource) Filter(query Q) *CosponsorshipResource {
	c.Filters = query

	return c
}

// All prepares and sends the http request to get all the cosponsorships
func (c *CosponsorshipResource) All() (*CosponsorshipsResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s%s", c.api.baseURL, c.Name, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var co *CosponsorshipsResponse
	if err := json.NewDecoder(resp.Body).Decode(&co); err != nil {
		return nil, err
	}

	return co, nil
}

// One prepares and sends the http request to get a cosponsorship
func (c *CosponsorshipResource) One(id string) (*CosponsorshipResponse, error) {
	filters := c.api.buildFilterQuery(c.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", c.api.baseURL, c.Name, id, filters)

	resp, err := c.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var co *CosponsorshipResponse
	if err := json.NewDecoder(resp.Body).Decode(&co); err != nil {
		return nil, err
	}

	return co, nil
}

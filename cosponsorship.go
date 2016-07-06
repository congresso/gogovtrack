package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// CosponsorshipsResponse is
type CosponsorshipsResponse struct {
	Meta           Meta             `json:"meta"`
	Cosponsorships []*Cosponsorship `json:"objects"`
}

// CosponsorshipResponse is
type CosponsorshipResponse Cosponsorship

// Cosponsorship is
type Cosponsorship struct {
	CosponrshipID int `json:"id"`
}

// CosponsorshipResource is
type CosponsorshipResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (c *CosponsorshipResource) Filter(query Q) *CosponsorshipResource {
	c.Filters = query

	return c
}

// All is
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

// One is
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

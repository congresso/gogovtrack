package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// BillsResponse is the json response for getting all bills
type BillsResponse struct {
	Meta  Meta    `json:"meta"`
	Bills []*Bill `json:"objects"`
}

// BillResponse is the json response for getting a bill
type BillResponse Bill

// Bill is holds the properties of what makes up a bill
type Bill struct {
	BillResolutionType string   `json:"bill_resolution_type"`
	BillType           string   `json:"bill_type"`
	Committess         []string `json:"committess"`
	Congress           int      `json:"congress"`
	BillID             int      `json:"id"`
}

// BillResource is the bill's resource
type BillResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (b *BillResource) Filter(query Q) *BillResource {
	b.Filters = query

	return b
}

// All gets all the bills
func (b *BillResource) All() (*BillsResponse, error) {
	filters := b.api.buildFilterQuery(b.Filters)

	url := fmt.Sprintf("%s/%s%s", b.api.baseURL, b.Name, filters)

	resp, err := b.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var br *BillsResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return nil, err
	}

	return br, nil
}

// One gets one bill by id
func (b *BillResource) One(id string) (*BillResponse, error) {
	filters := b.api.buildFilterQuery(b.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", b.api.baseURL, b.Name, id, filters)

	resp, err := b.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var br *BillResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return nil, err
	}

	return br, nil
}

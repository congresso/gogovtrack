package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// BillsResponse is
type BillsResponse struct {
	Meta  Meta   `json:"meta"`
	Bills []Bill `json:"objects"`
}

// BillResponse is
type BillResponse Bill

// Bill is
type Bill struct {
	BillResolutionType string   `json:"bill_resolution_type"`
	BillType           string   `json:"bill_type"`
	Committess         []string `json:"committess"`
	Congress           int      `json:"congress"`
	ID                 int      `json:"id"`
}

// BillResource is
type BillResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (b *BillResource) Filter(query Q) *BillResource {
	b.Filters = query

	return b
}

// All is
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

// One is
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

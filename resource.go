package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// Q is
type Q map[string]string

// Resource is
type Resource struct {
	api     *API
	Name    string
	Filters Q
}

// ResourceResponse is
type ResourceResponse struct {
	Objects []Object `json:"objects"`
}

// Object is
type Object struct {
	BillResolutionType string   `json:"bill_resolution_type"`
	BillType           string   `json:"bill_type"`
	Committess         []string `json:"committess"`
	Congress           int      `json:"congress"`
	ID                 int      `json:"id"`
}

// Filter is
func (r *Resource) Filter(query Q) *Resource {
	r.Filters = query

	return r
}

func (r *Resource) buildFilterQuery() string {
	// Check to see if we even have filters to build
	if len(r.Filters) < 1 {
		return ""
	}

	isFirst := true
	query := "?"

	for k, v := range r.Filters {
		if !isFirst {
			query += "&" + k + "=" + v
			continue
		}

		isFirst = false
		query += k + "=" + v
	}

	return query
}

// All is
func (r *Resource) All() (*ResourceResponse, error) {
	filters := r.buildFilterQuery()

	url := fmt.Sprintf("%s/%s%s", r.api.baseURL, r.Name, filters)

	resp, err := r.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var rr *ResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&rr); err != nil {
		return nil, err
	}

	return rr, nil
}

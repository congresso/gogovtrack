package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// RolesResponse is
type RolesResponse struct {
	Meta  Meta    `json:"meta"`
	Roles []*Role `json:"objects"`
}

// RoleResponse is
type RoleResponse Role

// Role is
type Role struct {
	RoleID int `json:"id"`
}

// RoleResource is
type RoleResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is
func (r *RoleResource) Filter(query Q) *RoleResource {
	r.Filters = query

	return r
}

// All is
func (r *RoleResource) All() (*RolesResponse, error) {
	filters := r.api.buildFilterQuery(r.Filters)

	url := fmt.Sprintf("%s/%s%s", r.api.baseURL, r.Name, filters)

	resp, err := r.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var re *RolesResponse
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		return nil, err
	}

	return re, nil
}

// One is
func (r *RoleResource) One(id string) (*RoleResponse, error) {
	filters := r.api.buildFilterQuery(r.Filters)

	url := fmt.Sprintf("%s/%s/%s%s", r.api.baseURL, r.Name, id, filters)

	resp, err := r.api.client.Get(url)
	if err != nil {
		return nil, err
	}

	var re *RoleResponse
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		return nil, err
	}

	return re, nil
}

package gogovtrack

import (
	"encoding/json"
	"fmt"
)

// RolesResponse is the json response for getting all roles
type RolesResponse struct {
	Meta  Meta    `json:"meta"`
	Roles []*Role `json:"objects"`
}

// RoleResponse is the json response for getting a role
type RoleResponse Role

// Role holds the properties of a role
type Role struct {
	RoleID int `json:"id"`
}

// RoleResource is the resource for the role api
type RoleResource struct {
	api     *API
	Name    string
	Filters Q
}

// Filter is a handy method for adding query filters to the request
func (r *RoleResource) Filter(query Q) *RoleResource {
	r.Filters = query

	return r
}

// All prepares and sends the http request to get all roles
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

// One prepares and sends the http request to get a role
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

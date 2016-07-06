package gogovtrack

import "net/http"

const (
	baseURL = "https://www.govtrack.us/api/v2"
)

// Service is a interface that
type Service interface {
}

// API is a struct that
type API struct {
	client  *http.Client
	baseURL string
}

// R is
func (a *API) R(name string) *Resource {
	return &Resource{Name: name, api: a}
}

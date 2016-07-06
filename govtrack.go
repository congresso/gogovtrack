package gogovtrack

import "net/http"

const (
	baseURL = "https://www.govtrack.us/api/v2/"
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
	return &Resource{Name: name}
}

// Helper function to make a get reuqest to the server
func (a *API) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", baseURL+url, nil)
	if err != nil {
		return nil, err
	}

	// Make the request
	return a.client.Do(req)
}

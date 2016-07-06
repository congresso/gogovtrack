package gogovtrack

import "net/http"

const (
	// BaseURL is the base url for the gov track v2 api
	BaseURL = "https://www.govtrack.us/api/v2"
)

// Service is a interface that
type Service interface {
}

// API is a struct that
type API struct {
	client  *http.Client
	baseURL string
}

// NewAPI is
func NewAPI(client *http.Client, baseURL string) *API {
	return &API{
		client:  client,
		baseURL: baseURL,
	}
}

// Br returns a BillResource
func (a *API) Br() *BillResource {
	return &BillResource{Name: "bill", api: a}
}

// Co returns a CommitteeResource
func (a *API) Co() *CommitteeResource {
	return &CommitteeResource{Name: "committee", api: a}
}

func (a *API) buildFilterQuery(filters Q) string {
	// Check to see if we even have filters to build
	if len(filters) < 1 {
		return ""
	}

	isFirst := true
	query := "?"

	for k, v := range filters {
		if !isFirst {
			query += "&" + k + "=" + v
			continue
		}

		isFirst = false
		query += k + "=" + v
	}

	return query
}

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
	BillID                      int           `json:"id"`
	BillResolutionType          string        `json:"bill_resolution_type"`
	BillType                    string        `json:"bill_type"`
	Committess                  []*Committee  `json:"committess"`
	Congress                    int           `json:"congress"`
	Cosponsors                  []*Person     `json:"cosponsor"`
	CurrentStatus               string        `json:"current_status"`
	CurrentStatusDate           string        `json:"current_status_date"`
	CurrentStatusDescription    string        `json:"current_status_description"`
	CurrentStatusLabel          string        `json:"current_status_label"`
	DisplayNumber               string        `json:"display_number"`
	DocsHouseGovPostdate        string        `json:"docs_house_gov_postdate"`
	IntroducedDate              string        `json:"introduced_date"`
	IsAlive                     bool          `json:"is_alive"`
	IsCurrent                   bool          `json:"is_current"`
	Link                        string        `json:"link"`
	LockTitle                   bool          `json:"lock_title"`
	MajorActions                []interface{} `json:"major_actions"`
	Noun                        string        `json:"noun"`
	Number                      int           `json:"number"`
	SenateFloorSchedulePostDate *string       `json:"senate_floor_schedule_postdate"`
	Sliplawnum                  *string       `json:"sliplawnum"`
	Sliplawpubpriv              *string       `json:"sliplawpubpriv"`
	Source                      string        `json:"source"`
	SourceLink                  *string       `json:"source_link"`
	Sponsor                     *Person       `json:"sponsor"`
	SponsorRole                 []Role        `json:"sponsor_role"`
	Terms                       []Term        `json:"terms"`
	Title                       string        `json:"title"`
	TitleWithoutNumber          string        `json:"title_without_number"`
	Titles                      []interface{} `json:"titles"`
}

// Term holds the properties of a term
type Term struct {
	TermID        int    `json:"id"`
	Name          string `json:"name"`
	TermType      string `json:"term_type"`
	TermTypeLabel string `json:"term_type_label"`
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

// All prepares and sends the http request to get all the bills
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

// One prepares and sends the http request to get a bill by id
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

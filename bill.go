package gogovtrack

// BillsResponse is
type BillsResponse struct {
	Meta  Meta    `json:"meta"`
	Bills []*Bill `json:"objects"`
}

// BillResponse is
type BillResponse Bill

// Bill is holds the properties of what makes up a bill
type Bill struct {
	BillID                      int           `json:"id"`
	BillResolutionType          string        `json:"bill_resolution_type"`
	BillType                    string        `json:"bill_type"`
	Committess                  []Committee   `json:"committess"`
	Congress                    int           `json:"congress"`
	Cosponsors                  []Person      `json:"cosponsor"`
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
	Number                      *int          `json:"number"`
	SenateFloorSchedulePostDate *string       `json:"senate_floor_schedule_postdate"`
	Sliplawnum                  *int          `json:"sliplawnum"`
	Sliplawpubpriv              *string       `json:"sliplawpubpriv"`
	Source                      string        `json:"source"`
	SourceLink                  *string       `json:"source_link"`
	Sponsor                     Person        `json:"sponsor"`
	SponsorRole                 EmbeddedRole  `json:"sponsor_role"`
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

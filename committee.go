package gogovtrack

// CommitteesResponse is
type CommitteesResponse struct {
	Meta       Meta        `json:"meta"`
	Committees []Committee `json:"objects"`
}

// CommitteeResponse is
type CommitteeResponse struct {
	Abbrev           string      `json:"abbrev"`
	Code             string      `json:"code"`
	Committee        Committee   `json:"committee"`
	CommitteeType    string      `json:"committee_type"`
	CommitteeID      int         `json:"id"`
	Jurisdiction     interface{} `json:"jurisdiction"`
	JurisdictionLink interface{} `json:"jurisdiction_link"`
	Name             string      `json:"name"`
	Obsolete         bool        `json:"obsolete"`
	URL              string      `json:"url"`
}

// Committee is
type Committee struct {
	Abbrev             string      `json:"abbrev"`
	Code               string      `json:"code"`
	Committee          interface{} `json:"committee"`
	CommitteeType      string      `json:"committee_type"`
	CommitteeTypeLabel string      `json:"committee_type_label"`
	ID                 int         `json:"id"`
	Jurisdiction       interface{} `json:"jurisdiction"`
	JurisdictionLink   interface{} `json:"jurisdiction_link"`
	Name               string      `json:"name"`
	Obsolete           bool        `json:"obsolete"`
	URL                string      `json:"url"`
}

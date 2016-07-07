package gogovtrack

import "net/http"

const (
	// BaseURL is the base url for the gov track v2 api
	BaseURL = "https://www.govtrack.us/api/v2"
	// BillResourceStr is a
	BillResourceStr = "bill"
	// CommitteeMemberResourceStr is
	CommitteeMemberResourceStr = "committee_member"
)

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

// Br returns a new BillResource
func (a *API) Br() *BillResource {
	return &BillResource{Name: "bill", api: a}
}

// Cr returns a new CommitteeResource
func (a *API) Cr() *CommitteeResource {
	return &CommitteeResource{Name: "committee", api: a}
}

// Cmr returns a new CommitteeMemberResource
func (a *API) Cmr() *CommitteeMemberResource {
	return &CommitteeMemberResource{Name: "committee_member", api: a}
}

// Csr returns a new CosponsorshipResource
func (a *API) Csr() *CosponsorshipResource {
	return &CosponsorshipResource{Name: "cosponsorship", api: a}
}

// Pr returns a new PersonResource
func (a *API) Pr() *PersonResource {
	return &PersonResource{Name: "person", api: a}
}

// Rr returns a new RoleResource
func (a *API) Rr() *RoleResource {
	return &RoleResource{Name: "role", api: a}
}

// Vvr returns a new VoteVoterResource
func (a *API) Vvr() *VoteVoterResource {
	return &VoteVoterResource{Name: "vote_voter", api: a}
}

// Vr returns a new VoteResource
func (a *API) Vr() *VoteResource {
	return &VoteResource{Name: "vote", api: a}
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

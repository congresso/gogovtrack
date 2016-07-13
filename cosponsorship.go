package gogovtrack

// CosponsorshipsResponse is
type CosponsorshipsResponse struct {
	Meta           Meta             `json:"meta"`
	Cosponsorships []*Cosponsorship `json:"objects"`
}

// CosponsorshipResponse is
type CosponsorshipResponse Cosponsorship

// Cosponsorship is
type Cosponsorship struct {
	CosponshorshipID int         `json:"id"`
	Bill             int         `json:"bill"`
	Joined           string      `json:"joined"`
	Person           int         `json:"person"`
	Role             int         `json:"role"`
	Withdrawn        interface{} `json:"withdrawn"`
}

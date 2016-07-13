package gogovtrack

// PersonsResponse is
type PersonsResponse struct {
	Meta    Meta      `json:"meta"`
	Persons []*Person `json:"objects"`
}

// PersonResponse is
type PersonResponse Person

// Person is
type Person struct {
	PersonID int `json:"id"`
}

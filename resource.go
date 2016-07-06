package gogovtrack

// Q is
type Q map[string]string

// Resource is
type Resource interface{}

// Meta is
type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

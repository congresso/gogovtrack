package gogovtrack

// Q is
type Q map[string]string

// Resource is
type Resource interface {
	Filter(query Q) interface{}
	One(id string) interface{}
	All() interface{}
}

// Meta is
type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

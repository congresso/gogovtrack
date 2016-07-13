package gogovtrack

// RolesResponse is
type RolesResponse struct {
	Meta  Meta    `json:"meta"`
	Roles []*Role `json:"objects"`
}

// RoleResponse is
type RoleResponse Role

// Role is
type Role struct {
	RoleID int `json:"id"`
}

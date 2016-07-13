package gogovtrack

// RolesResponse is
type RolesResponse struct {
	Meta  Meta   `json:"meta"`
	Roles []Role `json:"objects"`
}

// RoleResponse is
type RoleResponse Role

// Role is
type Role struct {
	Caucus          interface{} `json:"caucus"`
	CongressNumbers []int       `json:"congress_numbers"`
	Current         bool        `json:"current"`
	Description     string      `json:"description"`
	District        int         `json:"district"`
	Enddate         string      `json:"enddate"`
	Extra           interface{} `json:"extra"`
	RoleID          int         `json:"id"`
	LeadershipTitle interface{} `json:"leadership_title"`
	Party           string      `json:"party"`
	Person          Person      `json:"person"`
	Phone           interface{} `json:"phone"`
	RoleType        string      `json:"role_type"`
	RoleTypeLabel   string      `json:"role_type_label"`
	SenatorClass    interface{} `json:"senator_class"`
	SenatorRank     interface{} `json:"senator_rank"`
	Startdate       string      `json:"startdate"`
	State           string      `json:"state"`
	Title           string      `json:"title"`
	TitleLong       string      `json:"title_long"`
	Website         string      `json:"website"`
}

// EmbeddedRole is
type EmbeddedRole struct {
	Caucus          interface{} `json:"caucus"`
	CongressNumbers []int       `json:"congress_numbers"`
	Current         bool        `json:"current"`
	Description     string      `json:"description"`
	District        int         `json:"district"`
	Enddate         string      `json:"enddate"`
	Extra           interface{} `json:"extra"`
	RoleID          int         `json:"id"`
	LeadershipTitle interface{} `json:"leadership_title"`
	Party           string      `json:"party"`
	Person          int         `json:"person"`
	Phone           interface{} `json:"phone"`
	RoleType        string      `json:"role_type"`
	RoleTypeLabel   string      `json:"role_type_label"`
	SenatorClass    interface{} `json:"senator_class"`
	SenatorRank     interface{} `json:"senator_rank"`
	Startdate       string      `json:"startdate"`
	State           string      `json:"state"`
	Title           string      `json:"title"`
	TitleLong       string      `json:"title_long"`
	Website         string      `json:"website"`
}

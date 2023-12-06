package testrail

type rolesResponse struct {
	Offset int `json:"id"`
	Limit  int `json:"limit"`
	Size   int `json:"size"`
	Links  struct {
		Next string
		Prev string
	}
	Roles []Role
}

type Role struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	IsDefault      bool   `json:"is_default"`
	IsProjectAdmin bool   `json:"is_project_admin,omitempty"`
}

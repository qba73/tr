package testrail

// User holds user data.
type User struct {
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	EmailNotifications bool   `json:"email_notifications"`
	IsActive           bool   `json:"is_active"`
	IsAdmin            bool   `json:"is_admin"`
	Name               string `json:"name"`
	RoleID             int    `json:"role_id"`
	Role               string `json:"role"`
	Groups             []int  `json:"group_ids"`
	MFARequired        bool   `json:"mfa_required"`

	// A list of project IDs. Each ID is a project to which the user is assigned.
	AssignedProjects []int `json:"assigned_projects,omitempty"`
	SSOEnabled       bool  `json:"sso_enabled,omitempty"`
}

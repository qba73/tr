package testrail

import (
	"net/http"
	"time"
)

// userResp represents user data received from the server.
type userResp struct {
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

// Client holds data required to communicate with the API.
type Client struct {
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client
}

const libVersion = "0.0.1"

// NewClient knows how to construct a new default TestRail client.
func NewClient(baseURL string) *Client {
	return &Client{
		UserAgent: "testrail/" + libVersion,
		BaseURL:   baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

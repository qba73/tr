package testrail

import (
	"net/http"
	"time"
)

const (
	libVersion = "0.0.1"
	userAgent  = "testrail/" + libVersion
	mimeType   = "application/json"
)

// Client holds data required to communicate with the API.
type Client struct {
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient knows how to construct a new default TestRail client.
func NewClient(baseURL string) *Client {
	return &Client{
		UserAgent: "TestRailClient/" + libVersion,
		BaseURL:   baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

package golangzerotierserviceapi

// Client represents a client for interacting with the ZeroTier API.
type Client struct {
	BaseURL  string
	Token    string
	Insecure bool
}

// NewClient creates a new instance of Client with the provided base URL and API token.
func NewClient(baseURL string, token string, insecure bool) *Client {
	return &Client{
		BaseURL:  baseURL,
		Token:    token,
		Insecure: insecure,
	}
}

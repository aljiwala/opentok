package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/aljiwala/tokbox/config"
)

// Client ...
type Client struct {
	// HTTP client used to communicate with API
	Client *http.Client

	// User agent used when communicating with Twilio API
	UserAgent string

	// The Twilio API base URL
	BaseURL *url.URL

	// APIKey and APISecret which is used for authentication during API request.
	APIKey, APISecret string
}

// NewClient returns a new TokBox API client.
// Will load default http.Client if client is nil.
func NewClient(apiKey, apiSecret string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	c := &Client{
		Client:    client,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	return c
}

// Endpoint ...
func (c *Client) Endpoint(parts ...string) *url.URL {
	up := []string{config.APIVersion, "/", ""}
	up = append(up, parts...)
	u, _ := url.Parse(strings.Join(up, "/"))
	u.Path = fmt.Sprintf("/%s.%s", u.Path, "")
	return u
}

// NewRequest ...
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(parsed)
	req, _ := http.NewRequest(method, u.String(), body)
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// req.SetBasicAuth("", "")
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")

	return req, nil
}

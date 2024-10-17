// Package marksman implements a fully capable client for the SnipeIT API.
package marksman

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	// Version is the SemVer-compliant version string used by marksman for the User-Agent header in HTTP requests.
	Version = "0.0.1"

	headerContentType = "application/json"
	headerAccept      = "application/json"

	base = "/api/v1"
)

var (
	userAgent = fmt.Sprintf("Snipe/v%s", Version)

	defaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
)

// Client is an HTTP client responsible for interfacing with the SnipeIT API.
// It provides methods for every possible route found within the API.
type Client struct {
	root  string
	token string

	client *http.Client
}

// New creates a new client with the given options.
// See [Option] for more information about configuration.
func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, o := range opts {
		o(c)
	}

	if c.root == "" {
		return nil, errors.New("marksman: client: missing root URL")
	}

	if c.token == "" {
		return nil, errors.New("marksman: client: missing token")
	}

	if c.client == nil {
		c.client = defaultClient
	}

	return c, nil
}

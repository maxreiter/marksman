// package marksman implements an API client for the Snipe-IT API.
package marksman

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
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

type Client struct {
	root  string
	token string

	client *http.Client
}

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

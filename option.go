package marksman

import (
	"net/http"
)

type Option func(*Client)

func WithAddress(address string) Option {
	return func(c *Client) {
		c.address = address
	}
}

func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

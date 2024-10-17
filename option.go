package marksman

import "net/http"

type Option func(*Client)

func WithRoot(root string) Option {
	return func(c *Client) {
		c.root = root
	}
}

func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

package marksman

import "net/http"

// Option is used to configure a [Client].
type Option func(*Client)

// Root sets the FQDN a [Client] uses for interfacing with the API.
func Root(root string) Option {
	return func(c *Client) {
		c.root = root
	}
}

// Token sets the bearer token a [Client] uses for authorization with the API.
func Token(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// HTTPClient sets the underlying [http.Client] a [Client] uses for interfacing with the API.
func HTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// UserAgent sets the user agent to use in requests made to the API.
func UserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

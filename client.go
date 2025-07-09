package marksman

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/maxreiter/marksman/pkg/httpx"
)

type Client struct {
	address   string
	token     string
	userAgent string

	client *http.Client
}

func New(options ...Option) (*Client, error) {
	client := new(Client)
	for _, option := range options {
		option(client)
	}

	if client.address == "" {
		return nil, errors.New("marksman: a valid address is required")
	}

	if client.token == "" {
		return nil, errors.New("marksman: a valid token is required")
	}

	if client.userAgent == "" {
		client.userAgent = DefaultUserAgent
	}

	if client.client == nil {
		client.client = httpx.DefaultClient
	}

	return client, nil
}

func (c *Client) buildFullURL(parts ...any) (string, error) {
	fullUrl := []string{"api", "v1"}
	for _, part := range parts {
		switch asserted := part.(type) {
		case string:
			fullUrl = append(fullUrl, asserted)
		case int32:
			fullUrl = append(fullUrl, strconv.Itoa(int(asserted)))
		default:
			return "", fmt.Errorf("marksman: unknown path part in URL")
		}
	}

	return url.JoinPath(c.address, fullUrl...)
}

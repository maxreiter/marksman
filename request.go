package marksman

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/maxreiter/marksman/pkg/httpx"
)

var (
	ErrInvalidRequest = errors.New("marksman: invalid request parameter")

	jsonMethods = []string{http.MethodPost, http.MethodPatch, http.MethodPut}
)

func (c *Client) do(ctx context.Context, input, output any, method string, parts ...any) error {
	method = strings.ToUpper(method)
	if method == "" {
		return fmt.Errorf("%w: an HTTP method must be specified", ErrInvalidRequest)
	}

	if len(parts) == 0 {
		return fmt.Errorf("%w: a valid API URL path must be specified", ErrInvalidRequest)
	}

	fullURL, err := c.buildFullURL(parts...)
	if err != nil {
		return fmt.Errorf("building full address: %w", err)
	}

	var body io.Reader = http.NoBody
	if slices.Contains(jsonMethods, method) && input != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(input); err != nil {
			return fmt.Errorf("encoding request body: %w", err)
		}

		body = &buf
	}

	request, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return fmt.Errorf("creating new request: %w", err)
	}

	request.Header.Set("Content-Type", httpx.ContentTypeJSON)
	request.Header.Set("Accept", httpx.ContentTypeJSON)
	request.Header.Set("User-Agent", c.userAgent)
	request.Header.Set("Authorization", "Bearer "+c.token)

	if method == http.MethodGet && input != nil {
		values, err := query.Values(input)
		if err != nil {
			return fmt.Errorf("encoding request queries: %w", err)
		}

		request.URL.RawQuery = values.Encode()
	}

	response, err := c.client.Do(request)
	if err != nil {
		return fmt.Errorf("performing request: %w", err)
	}

	defer response.Body.Close()

	if !httpx.IsOK(response.StatusCode) {
		return createError(response)
	}

	if err := json.NewDecoder(response.Body).Decode(output); err != nil {
		return fmt.Errorf("decoding response body: %w", err)
	}

	return nil
}

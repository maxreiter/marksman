package marksman

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type request struct {
	method  string
	url     string
	body    io.Reader
	headers http.Header
	query   url.Values
}

func (c *Client) do(ctx context.Context, req request, out any) error {
	route := c.root + base + req.url

	request, err := http.NewRequestWithContext(ctx, req.method, route, req.body)
	if err != nil {
		return err
	}

	if req.headers != nil {
		request.Header = req.headers.Clone()
	}

	request.Header.Set("User-Agent", defaultUserAgent)
	request.Header.Set("Accept", headerAccept)
	request.Header.Set("Content-Type", headerContentType)
	request.Header.Set("Authorization", "Bearer "+c.token)

	if req.query != nil {
		request.URL.RawQuery = req.query.Encode()
	}

	res, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return HTTPError{
			Code: res.StatusCode,
		}
	}

	if out == nil {
		var ae *Error
		if err := json.NewDecoder(res.Body).Decode(&ae); err != nil {
			return err
		}

		if ae != nil {
			return ae
		}
	} else {
		var buf bytes.Buffer
		tee := io.TeeReader(res.Body, &buf)

		if err := json.NewDecoder(tee).Decode(&out); err != nil {
			var ae *Error
			if err := json.NewDecoder(&buf).Decode(&ae); err != nil {
				return err
			}

			return ae
		}
	}

	return nil
}

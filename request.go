package marksman

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type request struct {
	method  string
	url     string
	body    any
	headers http.Header
	query   any
}

func (c *Client) do(ctx context.Context, req request, out any) error {
	route := c.root + base + req.url

	var body io.Reader
	if req.body != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(req.body); err != nil {
			return err
		}

		body = &buf
	}

	request, err := http.NewRequestWithContext(ctx, req.method, route, body)
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
		values, err := query.Values(req.query)
		if err != nil {
			return err
		}

		request.URL.RawQuery = values.Encode()
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
		var buf bytes.Buffer
		tee := io.TeeReader(res.Body, &buf)

		body, _ := io.ReadAll(tee)

		fmt.Println(string(body))

		var response Response
		if err := json.NewDecoder(&buf).Decode(&response); err != nil {
			return err
		}

		if response.Status == ResponseError {
			return Error{response}
		}

		return nil
	}

	var buf bytes.Buffer
	copy := io.TeeReader(res.Body, &buf)

	var response Response
	if err := json.NewDecoder(copy).Decode(&response); err != nil {
		return err
	}

	if response.Status == ResponseError {
		return Error{response}
	}

	if err := json.NewDecoder(&buf).Decode(out); err != nil {
		return err
	}

	return nil
}

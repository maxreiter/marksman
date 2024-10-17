package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/statuslabel"
	"github.com/maxreiter/marksman/snipeit"
)

// StatusLabels is the expected response from endpoints that list [snipeit.StatusLabel].
type StatusLabels struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.StatusLabel `json:"rows"`
}

// StatusLabels fetches a list of [snipeit.StatusLabel].
//
// The following query parameters are accepted:
//   - [statuslabel.Name]
//   - [statuslabel.Limit]: defaults to 50
//   - [statuslabel.Offset]: defaults to 0
//   - [statuslabel.Search]
//   - [statuslabel.Sort]: defaults to "created_at"
//   - [statuslabel.Order]: defaults to "asc"
//   - [statuslabel.StatusType]
func (c *Client) StatusLabels(ctx context.Context, opts ...statuslabel.RequestOption) (*StatusLabels, error) {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/stauslabels",
		query:  values,
	}

	var response *StatusLabels
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateStatusLabel creates a new [snipeit.StatusLabel].
//
// The following body parameters are accepted:
//   - [statuslabel.Name]: required
//   - [statuslabel.Type]: required
//   - [statuslabel.Notes]
//   - [statuslabel.Color]
//   - [statuslabel.ShowInNav]: defaults to false
//   - [statuslabel.DefaultLabel]: defaults to false
func (c *Client) CreateStatusLabel(ctx context.Context, opts ...statuslabel.RequestOption) error {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/statuslabels",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// StatusLabel fetches a single [snipeit.StatusLabel].
func (c *Client) StatusLabel(ctx context.Context, id snipeit.StatusLabelID) (*snipeit.StatusLabel, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/statuslabels/%d", id),
	}

	var response *snipeit.StatusLabel
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateStatusLabel updates a [snipeit.StatusLabel].
//
// The following body parameters are accepted:
//   - [statuslabel.Name]: required
//   - [statuslabel.Type]: required
//   - [statuslabel.Notes]
//   - [statuslabel.Color]
//   - [statuslabel.ShowInNav]: defaults to false
//   - [statuslabel.DefaultLabel]: defaults to false
func (c *Client) UpdateStatusLabel(ctx context.Context, id snipeit.StatusLabelID, opts ...statuslabel.RequestOption) error {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/statuslabels/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// DeleteStatusLabel deletes a [snipeit.StatusLabel].
func (c *Client) DeleteStatusLabel(ctx context.Context, id snipeit.StatusLabelID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/statuslabels/%d", id),
	}

	return c.do(ctx, req, nil)
}

// StatusLabelAssets fetches a list of [snipeit.Asset] with the given [snipeit.StatusLabel].
func (c *Client) StatusLabelAssets(ctx context.Context, id snipeit.StatusLabelID) (*Assets, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/statuslabels/%d/assetlist", id),
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

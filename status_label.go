package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/statuslabel"
	"github.com/maxreiter/marksman/snipeit"
)

type StatusLabels struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.StatusLabel `json:"rows"`
}

func (c *Client) StatusLabels(ctx context.Context, opts ...statuslabel.RequestOption) (*StatusLabels, error) {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) CreateStatusLabel(ctx context.Context, opts ...statuslabel.RequestOption) error {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) UpdateStatusLabel(ctx context.Context, id snipeit.StatusLabelID, opts ...statuslabel.RequestOption) error {
	ro := &statuslabel.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) DeleteStatusLabel(ctx context.Context, id snipeit.StatusLabelID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/statuslabels/%d", id),
	}

	return c.do(ctx, req, nil)
}

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

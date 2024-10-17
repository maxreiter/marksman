package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/license"
	"github.com/maxreiter/marksman/snipeit"
)

type Licenses struct {
	Total int                `json:"total"`
	Rows  []*snipeit.License `json:"rows"`
}

func (c *Client) Licenses(ctx context.Context, opts ...license.RequestOption) (*Licenses, error) {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/licenses",
		query:  values,
	}

	var response *Licenses
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateLicense(ctx context.Context, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return nil
	}

	req := request{
		method: http.MethodPost,
		url:    "/licenses",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) UpdateLicense(ctx context.Context, id snipeit.LicenseID, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/licenses/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

func (c *Client) DeleteLicense(ctx context.Context, id snipeit.LicenseID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/licenses/%d", id),
	}

	return c.do(ctx, req, nil)
}

type LicenseSeats struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.LicenseSeat `json:"rows"`
}

func (c *Client) LicenseSeats(ctx context.Context, id snipeit.LicenseID) (*LicenseSeats, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/licenses/%d/seats", id),
	}

	var response *LicenseSeats
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) LicenseSeat(ctx context.Context, licenseID snipeit.LicenseID, seatID snipeit.LicenseSeatID) (*snipeit.LicenseSeat, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/licenses/%d/seats/%d", licenseID, seatID),
	}

	var response *snipeit.LicenseSeat
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateLicenseSeat(ctx context.Context, licenseID snipeit.LicenseID, seatID snipeit.LicenseSeatID, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/licenses/%d/seats/%d", licenseID, seatID),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

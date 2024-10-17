package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/license"
	"github.com/maxreiter/marksman/snipeit"
)

// Licenses is the expected response from endpoints that list [snipeit.License].
type Licenses struct {
	Total int                `json:"total"`
	Rows  []*snipeit.License `json:"rows"`
}

// Licenses fetches a list of [snipeit.License].
//
// The following query parameters are accepted:
//   - [license.Name]
//   - [license.ProductKey]
//   - [license.Limit]: defaults to 50
//   - [license.Offset]: defaults to 0
//   - [license.Search]
//   - [license.OrderNumber]
//   - [license.Sort]: defaults to "created_at"
//   - [license.Order]: defaults to "desc"
//   - [license.Expand]: defaults to false
//   - [license.PurchaseOrder]
//   - [license.LicenseName]
//   - [license.LicenseEmail]
//   - [license.ManufacturerID]
//   - [license.SupplierID]
//   - [license.CategoryID]
//   - [license.DepreciationID]
//   - [license.Maintained]
//   - [license.Deleted]
func (c *Client) Licenses(ctx context.Context, opts ...license.RequestOption) (*Licenses, error) {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/licenses",
		query:  ro,
	}

	var response *Licenses
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateLicense creates a new [snipeit.License].
//
// The following body parameters are accepted:
//   - [license.Name]: required
//   - [license.Seats]: required
//   - [license.CategoryID]: required
//   - [license.CompanyID]
//   - [license.ExpirationDate]
//   - [license.LicenseEmail]
//   - [license.LicenseName]
//   - [license.Serial]
//   - [license.Maintained]
//   - [license.ManufacturerID]
//   - [license.Notes]
//   - [license.OrderNumber]
//   - [license.PurchaseCost]
//   - [license.PurchaseDate]
//   - [license.PurchaseOrder]
//   - [license.Reassignable]
//   - [license.SupplierID]
//   - [license.TerminationDate]
func (c *Client) CreateLicense(ctx context.Context, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    "/licenses",
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// License fetches a single [snipeit.License].
func (c *Client) License(ctx context.Context, id snipeit.LicenseID) (*snipeit.License, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/licenses/%d", id),
	}

	var response *snipeit.License
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateLicense updates a [snipeit.License].
//
// The following body parameters are accepted:
//   - [license.Name]
//   - [license.CategoryID]: required
//   - [license.CompanyID]
//   - [license.ExpirationDate]
//   - [license.LicenseEmail]
//   - [license.LicenseName]
//   - [license.Serial]
//   - [license.Maintained]
//   - [license.ManufacturerID]
//   - [license.Notes]
//   - [license.OrderNumber]
//   - [license.PurchaseCost]
//   - [license.PurchaseDate]
//   - [license.PurchaseOrder]
//   - [license.Reassignable]
//   - [license.Seats]
//   - [license.SupplierID]
//   - [license.TerminationDate]
func (c *Client) UpdateLicense(ctx context.Context, id snipeit.LicenseID, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/licenses/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteLicense deletes a [snipeit.License].
func (c *Client) DeleteLicense(ctx context.Context, id snipeit.LicenseID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/licenses/%d", id),
	}

	return c.do(ctx, req, nil)
}

// LicenseSeats is the expected response from endpoints that list [snipeit.LicenseSeat].
type LicenseSeats struct {
	Total int                    `json:"total"`
	Rows  []*snipeit.LicenseSeat `json:"rows"`
}

// LicenseSeats fetches a list of [snipeit.LicenseSeat] for a [snipeit.License].
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

// LicenseSeat fetches a single [snipeit.LicenseSeat] for a [snipeit.License].
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

// UpdateLicenseSeat updates a [snipeit.LicenseSeat].
//
// The following body parameters are accepted:
//   - [license.AssignedTo]: should be null if checking in seat
//   - [license.AssetID]: should be null if checkin in seat
//   - [license.Note]
func (c *Client) UpdateLicenseSeat(ctx context.Context, licenseID snipeit.LicenseID, seatID snipeit.LicenseSeatID, opts ...license.RequestOption) error {
	ro := &license.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/licenses/%d/seats/%d", licenseID, seatID),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

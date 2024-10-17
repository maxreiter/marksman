package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/asset"
	"github.com/maxreiter/marksman/snipeit"
)

// Assets is the expected response from endpoints that list assets.
type Assets struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Asset `json:"rows"`
}

// Assets fetches the list of [snipeit.Asset].
//
// The following query parameters are accepted:
//   - [asset.Limit]: defaults to 2
//   - [asset.Offset]: defaults to 0
//   - [asset.Search]
//   - [asset.OrderNumber]
//   - [asset.Sort]: defaults to "created_at"
//   - [asset.Order]: defaults to "desc"
//   - [asset.ModelID]
//   - [asset.CategoryID]
//   - [asset.ManufacturerID]
//   - [asset.CompanyID]
//   - [asset.LocationID]
//   - [asset.Status]
//   - [asset.StatusID]
func (c *Client) Assets(ctx context.Context, opts ...asset.RequestOption) (*Assets, error) {
	ro := &asset.RequestOptions{
		Limit:  50,
		Offset: 0,
		Sort:   "created_at",
		Order:  "desc",
	}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    "/hardware",
		query:  values,
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateAsset creates a new [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [asset.AssetTag]: required
//   - [asset.StatusID]: required
//   - [asset.ModelID]: required
//   - [asset.Name]
//   - [asset.Image]
//   - [asset.Serial]
//   - [asset.PurchaseDate]
//   - [asset.PurchaseCost]
//   - [asset.OrderNumber]
//   - [asset.Notes]
//   - [asset.Archived]: defaults to false
//   - [asset.WarrantyMonths]: defaults to null
//   - [asset.Depreciate]: defaults to false
//   - [asset.SupplierID]: defaults to null
//   - [asset.Requestable]: defaults to false
//   - [asset.RTDLocationID]: defaults to null
//   - [asset.LastAuditDate]
//   - [asset.LocationID]: defaults to null
//   - [asset.BYOD]
func (c *Client) CreateAsset(ctx context.Context, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/hardware",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// Asset fetches a single [snipeit.Asset].
func (c *Client) Asset(ctx context.Context, id snipeit.AssetID) (*snipeit.Asset, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/hardware/%d", id),
	}

	var response *snipeit.Asset
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// AssetByTag fetches a single [snipeit.Asset] by its associated asset tag.
//
// The following query parameters are accepted:
//   - [asset.Deleted]: defaults to false
func (c *Client) AssetByTag(ctx context.Context, tag string, opts ...asset.RequestOption) (*snipeit.Asset, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/hardware/bytag/%s", tag),
		query:  values,
	}

	var response *snipeit.Asset
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// AssetBySerial fetches a single [snipeit.Asset] by its associated serial.
//
// The following query parameters are accepted.
//   - [asset.Deleted]: defaults to false
func (c *Client) AssetBySerial(ctx context.Context, serial string, opts ...asset.RequestOption) (*snipeit.Asset, error) {
	ro := &asset.RequestOptions{}
	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/hardware/byserial/%s", serial),
		query:  values,
	}

	var response *snipeit.Asset
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAsset updates an [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [asset.AssetTag]: required
//   - [asset.StatusID]: required
//   - [asset.ModelID]: required
//   - [asset.Notes]
//   - [asset.LastCheckout]: defaults to null
//   - [asset.AssignedUser]: defaults to null
//   - [asset.AssignedLocation]: defaults to null
//   - [asset.AssignedAsset]: defaults to null
//   - [asset.CompanyID]: defaults to null
//   - [asset.Serial]: defaults to null
//   - [asset.OrderNumber]
//   - [asset.WarrantyMonths]: defaults to null
//   - [asset.PurchaseCost]: defaults to null
//   - [asset.PurchaseDate]: defaults to null
//   - [asset.Requestable]: defaults to false
//   - [asset.Archived]: defaults to false
//   - [asset.RTDLocationID]: defaults to null
//   - [asset.Name]: defaults to null
//   - [asset.LocationID]: defaults to null
//   - [asset.Image]: defaults to null
//   - [asset.BYOD]
func (c *Client) UpdateAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/hardware/%d", id),
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// DeleteAsset deletes an [snipeit.Asset].
func (c *Client) DeleteAsset(ctx context.Context, id snipeit.AssetID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/hardware/%d", id),
	}

	return c.do(ctx, req, nil)
}

// CheckoutAsset checks an [snipeit.Asset] out to a [snipeit.User], [snipeit.Location], or another [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [asset.StatusID]: required
//   - [asset.CheckoutToType]: required, defaults to user
//   - [asset.AssignedUser]: required if checkout_to_type is set to "user"
//   - [asset.AssignedAsset]: required if checkout_to_type is set to "asset"
//   - [asset.AssignedLocation]: required if checkout_to_type is set to "location"
//   - [asset.ExpectedCheckin]
//   - [asset.CheckoutAt]
//   - [asset.Name]
//   - [asset.Note]
func (c *Client) CheckoutAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) (*Assets, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/hardware/%d/checkout", id),
		body:   bod,
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CheckinAsset checks in an [snipeit.Asset].
//
// The following body parameters are accepted:
//   - [asset.StatusID]: required, defaults to null
//   - [asset.Name]: defaults to null
//   - [asset.Note]
//   - [asset.LocationID]
func (c *Client) CheckinAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) (*Assets, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return nil, err
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/hardware/%d/checkout", id),
		body:   bod,
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// AuditAsset marks an [snipeit.Asset] as audited.
//
// The following body parameters are accepted:
//   - [asset.AssetTag]: required
//   - [asset.LocationID]
//   - [asset.NextAuditDate]
func (c *Client) AuditAsset(ctx context.Context, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/hardware/audit",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// DueAssets fetches a list of [snipeit.Asset] that are due for auditing.
func (c *Client) DueAssets(ctx context.Context) (*Assets, error) {
	req := request{
		method: http.MethodGet,
		url:    "/hardware/audit/due",
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// OverdueAssets fetches a list of [snipeit.Asset] that are overdue for auditing.
func (c *Client) OverdueAssets(ctx context.Context) (*Assets, error) {
	req := request{
		method: http.MethodGet,
		url:    "/hardware/audit/overdue",
	}

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// RestoreAsset retores an [snipeit.Asset] that had been deleted.
func (c *Client) RestoreAsset(ctx context.Context, id snipeit.AssetID) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/hardware/%d/restore", id),
	}

	return c.do(ctx, req, nil)
}

// AssetLicenses fetches a list of [snipeit.License] that are checked out to an [snipeit.Asset].
func (c *Client) AssetLicenses(ctx context.Context, id snipeit.AssetID) (*Licenses, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/hardware/%d/licenses", id),
	}

	var response *Licenses
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

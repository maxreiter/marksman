package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/asset"
	"github.com/maxreiter/marksman/snipeit"
)

type Assets struct {
	Total int              `json:"total"`
	Rows  []*snipeit.Asset `json:"rows"`
}

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

	values, err := ro.Values()
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

func (c *Client) CreateAsset(ctx context.Context, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) AssetByTag(ctx context.Context, tag string, opts ...asset.RequestOption) (*snipeit.Asset, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) AssetBySerial(ctx context.Context, serial string, opts ...asset.RequestOption) (*snipeit.Asset, error) {
	ro := &asset.RequestOptions{}
	for _, o := range opts {
		o(ro)
	}

	values, err := ro.Values()
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

func (c *Client) UpdateAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) DeleteAsset(ctx context.Context, id snipeit.AssetID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/hardware/%d", id),
	}

	return c.do(ctx, req, nil)
}

func (c *Client) CheckoutAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) (*Assets, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) CheckinAsset(ctx context.Context, id snipeit.AssetID, opts ...asset.RequestOption) (*Assets, error) {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) AuditAsset(ctx context.Context, opts ...asset.RequestOption) error {
	ro := &asset.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	bod, err := ro.Marshal()
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

func (c *Client) RestoreAsset(ctx context.Context, id snipeit.AssetID) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/hardware/%d/restore", id),
	}

	return c.do(ctx, req, nil)
}

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

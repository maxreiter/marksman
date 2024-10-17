package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/location"
	"github.com/maxreiter/marksman/snipeit"
)

// Locations is the expected response from endpoints that list [snipeit.Location].
type Locations struct {
	Total int                 `json:"total"`
	Rows  []*snipeit.Location `json:"rows"`
}

// Locations fetches a list of [snipeit.Location].
//
// The following query parameters are accepted:
//   - [location.Limit]: defaults to 50
//   - [location.Offset]: defaults to 0
//   - [location.Search]
//   - [location.Sort]: defaults to "created_at"
//   - [location.Order]
//   - [location.Address]
//   - [location.Address2]
//   - [location.City]
//   - [location.Zip]
//   - [location.Country]
func (c *Client) Locations(ctx context.Context, opts ...location.RequestOption) (*Locations, error) {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    "/locations",
		query:  ro,
	}

	var response *Locations
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateLocation creates a new [snipeit.Location].
//
// The following body parameters are accepted:
//   - [location.Name]: required
//   - [location.Address]
//   - [location.Address2]
//   - [location.City]
//   - [location.State]
//   - [location.Country]
//   - [location.Zip]
//   - [location.LDAPOU]
//   - [location.ParentID]
//   - [location.Currency]
//   - [location.ManagerID]
func (c *Client) CreateLocation(ctx context.Context, id snipeit.LocationID, opts ...location.RequestOption) error {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/locations/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// Location fetches a single [snipeit.Location].
func (c *Client) Location(ctx context.Context, id snipeit.LocationID) (*snipeit.Location, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/locations/%d", id),
	}

	var response *snipeit.Location
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateLocation updates a [snipeit.Location].
//
// The following body parameters are accepted:
//   - [location.Name]: required
//   - [location.Address]
//   - [location.Address2]
//   - [location.City]
//   - [location.State]
//   - [location.Country]
//   - [location.Zip]
//   - [location.LDAPOU]
//   - [location.Currency]
//   - [location.ManagerID]
//   - [location.ParentID]
func (c *Client) UpdateLocation(ctx context.Context, id snipeit.LocationID, opts ...location.RequestOption) error {
	ro := &location.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodPut,
		url:    fmt.Sprintf("/locations/%d", id),
		body:   ro,
	}

	return c.do(ctx, req, nil)
}

// DeleteLocation deletes a [snipeit.Location].
func (c *Client) DeleteLocation(ctx context.Context, id snipeit.LocationID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/locations/%d", id),
	}

	return c.do(ctx, req, nil)
}

package marksman

import (
	"context"
	"fmt"
	"net/http"

	"github.com/maxreiter/marksman/params/user"
	"github.com/maxreiter/marksman/snipeit"
)

// Users is the expected response from endpoints that list [snipeit.User].
type Users struct {
	Total int             `json:"total"`
	Rows  []*snipeit.User `json:"rows"`
}

// Me fetches details on the current [snipeit.User].
func (c *Client) Me(ctx context.Context) (*snipeit.User, error) {
	req := request{
		method: http.MethodGet,
		url:    "/users/me",
	}

	var response *snipeit.User
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// Users fetches a list of [snipeit.User].
//
// The following query parameters are accepted:
//   - [user.Limit]: defaults to 50
//   - [user.Offset]: defaults to 0
//   - [user.Sort]: defaults to "created_at"
//   - [user.Order]: defaults to "desc"
//   - [user.Deleted]: defaults to false
//   - [user.All]: defaults to false
//   - [user.FirstName]
//   - [user.LastName]
//   - [user.Search]
//   - [user.Username]
//   - [user.Email]
//   - [user.EmployeeNum]
//   - [user.State]
//   - [user.Zip]
//   - [user.Country]
//   - [user.GroupID]
//   - [user.DepartmentID]
//   - [user.CompanyID]
//   - [user.LocationID]
//   - [user.LDAPImport]
//   - [user.AssetsCount]
//   - [user.LicensesCount]
//   - [user.AccessoriesCount]
//   - [user.ConsumablesCount]
//   - [user.Remote]
//   - [user.VIP]
//   - [user.StartDate]
//   - [user.EndDate]
func (c *Client) Users(ctx context.Context, opts ...user.RequestOption) (*Users, error) {
	ro := &user.RequestOptions{
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
		query:  values,
		method: http.MethodGet,
		url:    "/users",
	}

	var response *Users
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// User fetches a single [snipeit.User].
func (c *Client) User(ctx context.Context, id snipeit.UserID) (*snipeit.User, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d", id),
	}

	var response *snipeit.User
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateUser creates a new [snipeit.User].
//
// The following body parameters are accepted:
//   - [user.FirstName]: required
//   - [user.Username]: required
//   - [user.Password]: required
//   - [user.LastName]
//   - [user.Email]
//   - [user.Permissions]
//   - [user.Activated]
//   - [user.Phone]
//   - [user.JobTitle]
//   - [user.ManagerID]
//   - [user.EmployeeNum]
//   - [user.Notes]
//   - [user.CompanyID]
//   - [user.TwoFactorEnrolled]
//   - [user.TwoFactorOptin]
//   - [user.DepartmentID]
//   - [user.LocationID]
//   - [user.Remote]
//   - [user.Groups]: defaults to null
//   - [user.VIP]: defaults to false
//   - [user.StartDate]
//   - [user.EndDate]
func (c *Client) CreateUser(ctx context.Context, opts ...user.RequestOption) error {
	ro := &user.RequestOptions{}

	bod, err := ro.JSON()
	if err != nil {
		return err
	}

	req := request{
		method: http.MethodPost,
		url:    "/users",
		body:   bod,
	}

	return c.do(ctx, req, nil)
}

// UpdateUser updates a [snipeit.User].
//
// The following body parameters are accepted:
//   - [user.FirstName]: required
//   - [user.Username]: required
//   - [user.Password]: required
//   - [user.VIP]: defaults to false
//   - [user.LastName]
//   - [user.Email]
//   - [user.Permissions]
//   - [user.Activated]
//   - [user.Phone]
//   - [user.JobTitle]
//   - [user.ManagerID]
//   - [user.EmployeeNum]
//   - [user.Notes]
//   - [user.CompanyID]
//   - [user.TwoFactorEnrolled]
//   - [user.TwoFactorOptin]
//   - [user.DepartmentID]
//   - [user.LocationID]
//   - [user.Remote]
//   - [user.Groups]
//   - [user.StartDate]
//   - [user.EndDate]
func (c *Client) UpdateUser(ctx context.Context, id snipeit.UserID, opts ...user.RequestOption) error {
	req := request{
		method: http.MethodPatch,
		url:    fmt.Sprintf("/users/%d", id),
	}

	reqOpts := &user.RequestOptions{}

	for _, o := range opts {
		o(reqOpts)
	}

	bod, err := reqOpts.JSON()
	if err != nil {
		return err
	}

	req.body = bod

	return c.do(ctx, req, nil)
}

// DeleteUser delete a [snipeit.User].
func (c *Client) DeleteUser(ctx context.Context, id snipeit.UserID) error {
	req := request{
		method: http.MethodDelete,
		url:    fmt.Sprintf("/users/%d", id),
	}

	return c.do(ctx, req, nil)
}

// RestoreUser restores a deleted [snipeit.User].
func (c *Client) RestoreUser(ctx context.Context, id snipeit.UserID) error {
	req := request{
		method: http.MethodPost,
		url:    fmt.Sprintf("/users/%d/restore", id),
	}

	return c.do(ctx, req, nil)
}

// UserAssets fetches a list of [snipeit.Asset] associated with a [snipeit.User].
//
// The following query parameters are accepted:
//   - [user.CategoryID]
//   - [user.ModelID]
func (c *Client) UserAssets(ctx context.Context, id snipeit.UserID, opts ...user.RequestOption) (*Assets, error) {
	ro := &user.RequestOptions{}

	for _, o := range opts {
		o(ro)
	}

	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d/assets", id),
	}

	values, err := ro.Query()
	if err != nil {
		return nil, err
	}

	req.query = values

	var response *Assets
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UserAccessories fetches a list of [snipeit.Accessory] associated with a [snipeit.User].
func (c *Client) UserAccessories(ctx context.Context, id snipeit.UserID) (*Accessories, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/users/%d/accessories", id),
	}

	var response *Accessories
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// UserLicenses fethces a list of [snipeit.License] associated with a [snipeit.User].
func (c *Client) UserLicenses(ctx context.Context, id snipeit.UserID) (*Licenses, error) {
	req := request{
		method: http.MethodGet,
		url:    fmt.Sprintf("/licenses/%d", id),
	}

	var response *Licenses
	if err := c.do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

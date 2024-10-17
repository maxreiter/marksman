// Package supplier provides request configuration for methods of the marksman client.
package supplier

// RequestOptions contains possible options for requests made to the /suppliers endpoints.
type RequestOptions struct {
	Name     string `url:"name,omitempty" json:"-"`
	Address  string `url:"address,omitempty" json:"-"`
	Address2 string `url:"address2,omitempty" json:"-"`
	City     string `url:"city,omitempty" json:"-"`
	Zip      string `url:"zip,omitempty" json:"-"`
	Country  string `url:"country,omitempty" json:"-"`
	Fax      string `url:"fax,omitempty" json:"-"`
	Email    string `url:"email,omitempty" json:"-"`
	URL      string `url:"url,omitempty" json:"-"`
	Notes    string `url:"notes,omitempty" json:"-"`
}

// RequestOption is used to configure a [RequestOptions].
type RequestOption func(*RequestOptions)

// Name sets the name of a [snipeit.Supplier].
func Name(name string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Name = name
	}
}

// Address sets the address of a [snipeit.Supplier].
func Address(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address = address
	}
}

// Address2 sets the alternate address of a [snipeit.Supplier].
func Address2(address string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Address2 = address
	}
}

// City sets the city of a [snipeit.Supplier].
func City(city string) RequestOption {
	return func(ro *RequestOptions) {
		ro.City = city
	}
}

// Zip sets the zip code of a [snipeit.Supplier].
func Zip(zip string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Zip = zip
	}
}

// Country sets the country of a [snipeit.Supplier].
func Country(country string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Country = country
	}
}

// Fax sets the fax number of a [snipeit.Supplier].
func Fax(fax string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Fax = fax
	}
}

// Email sets the email address of a [snipeit.Supplier].
func Email(email string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Email = email
	}
}

// URL sets the URL of a [snipeit.Supplier].
func URL(baseURL string) RequestOption {
	return func(ro *RequestOptions) {
		ro.URL = baseURL
	}
}

// Notes sets the notes of a [snipeit.Supplier].
func Notes(notes string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Notes = notes
	}
}

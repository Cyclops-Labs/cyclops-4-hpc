// Code generated by go-swagger; DO NOT EDIT.

package credit_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// AddConsumptionURL generates an URL for the add consumption operation
type AddConsumptionURL struct {
	ID     string
	Medium string

	Amount float64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddConsumptionURL) WithBasePath(bp string) *AddConsumptionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddConsumptionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AddConsumptionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/{medium}/consume/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on AddConsumptionURL")
	}

	medium := o.Medium
	if medium != "" {
		_path = strings.Replace(_path, "{medium}", medium, -1)
	} else {
		return nil, errors.New("medium is required on AddConsumptionURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	amountQ := swag.FormatFloat64(o.Amount)
	if amountQ != "" {
		qs.Set("amount", amountQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AddConsumptionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AddConsumptionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AddConsumptionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AddConsumptionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AddConsumptionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AddConsumptionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

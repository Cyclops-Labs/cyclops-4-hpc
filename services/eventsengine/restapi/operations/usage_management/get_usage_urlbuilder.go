// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetUsageURL generates an URL for the get usage operation
type GetUsageURL struct {
	ID string

	From     *int64
	Region   *string
	Resource *string
	To       *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsageURL) WithBasePath(bp string) *GetUsageURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsageURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUsageURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/usage/{id}"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetUsageURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var fromQ string
	if o.From != nil {
		fromQ = swag.FormatInt64(*o.From)
	}
	if fromQ != "" {
		qs.Set("from", fromQ)
	}

	var regionQ string
	if o.Region != nil {
		regionQ = *o.Region
	}
	if regionQ != "" {
		qs.Set("region", regionQ)
	}

	var resourceQ string
	if o.Resource != nil {
		resourceQ = *o.Resource
	}
	if resourceQ != "" {
		qs.Set("resource", resourceQ)
	}

	var toQ string
	if o.To != nil {
		toQ = swag.FormatInt64(*o.To)
	}
	if toQ != "" {
		qs.Set("to", toQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUsageURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUsageURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUsageURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUsageURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUsageURL")
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
func (o *GetUsageURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
)

// GetSystemUsageURL generates an URL for the get system usage operation
type GetSystemUsageURL struct {
	From   *strfmt.DateTime
	Idlist *string
	Metric *string
	To     *strfmt.DateTime

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetSystemUsageURL) WithBasePath(bp string) *GetSystemUsageURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetSystemUsageURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetSystemUsageURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/usage"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var fromQ string
	if o.From != nil {
		fromQ = o.From.String()
	}
	if fromQ != "" {
		qs.Set("from", fromQ)
	}

	var idlistQ string
	if o.Idlist != nil {
		idlistQ = *o.Idlist
	}
	if idlistQ != "" {
		qs.Set("idlist", idlistQ)
	}

	var metricQ string
	if o.Metric != nil {
		metricQ = *o.Metric
	}
	if metricQ != "" {
		qs.Set("metric", metricQ)
	}

	var toQ string
	if o.To != nil {
		toQ = o.To.String()
	}
	if toQ != "" {
		qs.Set("to", toQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetSystemUsageURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetSystemUsageURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetSystemUsageURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetSystemUsageURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetSystemUsageURL")
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
func (o *GetSystemUsageURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
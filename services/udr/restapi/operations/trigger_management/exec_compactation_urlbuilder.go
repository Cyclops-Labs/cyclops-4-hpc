// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExecCompactationURL generates an URL for the exec compactation operation
type ExecCompactationURL struct {
	FastMode *bool
	From     *strfmt.DateTime
	To       *strfmt.DateTime

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ExecCompactationURL) WithBasePath(bp string) *ExecCompactationURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ExecCompactationURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ExecCompactationURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/trigger/compact"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var fastModeQ string
	if o.FastMode != nil {
		fastModeQ = swag.FormatBool(*o.FastMode)
	}
	if fastModeQ != "" {
		qs.Set("fast_mode", fastModeQ)
	}

	var fromQ string
	if o.From != nil {
		fromQ = o.From.String()
	}
	if fromQ != "" {
		qs.Set("from", fromQ)
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
func (o *ExecCompactationURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ExecCompactationURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ExecCompactationURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ExecCompactationURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ExecCompactationURL")
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
func (o *ExecCompactationURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
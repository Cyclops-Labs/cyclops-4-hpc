// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSystemUsageParams creates a new GetSystemUsageParams object
// no default values defined in spec.
func NewGetSystemUsageParams() GetSystemUsageParams {

	return GetSystemUsageParams{}
}

// GetSystemUsageParams contains all the bound params for the get system usage operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSystemUsage
type GetSystemUsageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Datetime from which to get the usage report
	  In: query
	*/
	From *int64
	/*Resource region to filter the usage
	  In: query
	*/
	Region *string
	/*Resource type to filter the usage
	  In: query
	*/
	Resource *string
	/*Datetime until which to get the usage report
	  In: query
	*/
	To *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSystemUsageParams() beforehand.
func (o *GetSystemUsageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qRegion, qhkRegion, _ := qs.GetOK("region")
	if err := o.bindRegion(qRegion, qhkRegion, route.Formats); err != nil {
		res = append(res, err)
	}

	qResource, qhkResource, _ := qs.GetOK("resource")
	if err := o.bindResource(qResource, qhkResource, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetSystemUsageParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = &value

	return nil
}

// bindRegion binds and validates parameter Region from query.
func (o *GetSystemUsageParams) bindRegion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Region = &raw

	return nil
}

// bindResource binds and validates parameter Resource from query.
func (o *GetSystemUsageParams) bindResource(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Resource = &raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetSystemUsageParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = &value

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewExecCompactationParams creates a new ExecCompactationParams object
// no default values defined in spec.
func NewExecCompactationParams() ExecCompactationParams {

	return ExecCompactationParams{}
}

// ExecCompactationParams contains all the bound params for the exec compactation operation
// typically these are obtained from a http.Request
//
// swagger:parameters execCompactation
type ExecCompactationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Switch for using 15m boundaries instead of 8h
	  In: query
	*/
	FastMode *bool
	/*Datetime from which to get the usage report
	  In: query
	*/
	From *strfmt.DateTime
	/*Datetime until which to get the usage report
	  In: query
	*/
	To *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewExecCompactationParams() beforehand.
func (o *ExecCompactationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFastMode, qhkFastMode, _ := qs.GetOK("fast_mode")
	if err := o.bindFastMode(qFastMode, qhkFastMode, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
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

// bindFastMode binds and validates parameter FastMode from query.
func (o *ExecCompactationParams) bindFastMode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("fast_mode", "query", "bool", raw)
	}
	o.FastMode = &value

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *ExecCompactationParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: datetime
	value, err := formats.Parse("datetime", raw)
	if err != nil {
		return errors.InvalidType("from", "query", "strfmt.DateTime", raw)
	}
	o.From = (value.(*strfmt.DateTime))

	if err := o.validateFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateFrom carries on validations for parameter From
func (o *ExecCompactationParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "datetime", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindTo binds and validates parameter To from query.
func (o *ExecCompactationParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: datetime
	value, err := formats.Parse("datetime", raw)
	if err != nil {
		return errors.InvalidType("to", "query", "strfmt.DateTime", raw)
	}
	o.To = (value.(*strfmt.DateTime))

	if err := o.validateTo(formats); err != nil {
		return err
	}

	return nil
}

// validateTo carries on validations for parameter To
func (o *ExecCompactationParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "datetime", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}
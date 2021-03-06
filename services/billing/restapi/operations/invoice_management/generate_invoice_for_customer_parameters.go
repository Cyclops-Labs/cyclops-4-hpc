// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGenerateInvoiceForCustomerParams creates a new GenerateInvoiceForCustomerParams object
// no default values defined in spec.
func NewGenerateInvoiceForCustomerParams() GenerateInvoiceForCustomerParams {

	return GenerateInvoiceForCustomerParams{}
}

// GenerateInvoiceForCustomerParams contains all the bound params for the generate invoice for customer operation
// typically these are obtained from a http.Request
//
// swagger:parameters GenerateInvoiceForCustomer
type GenerateInvoiceForCustomerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Datetime from which to generate the invoice
	  In: query
	*/
	From *strfmt.DateTime
	/*Id of the account to be checked
	  Required: true
	  In: path
	*/
	ID string
	/*Datetime until which to generate the invoice
	  In: query
	*/
	To *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGenerateInvoiceForCustomerParams() beforehand.
func (o *GenerateInvoiceForCustomerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
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
func (o *GenerateInvoiceForCustomerParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GenerateInvoiceForCustomerParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "datetime", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GenerateInvoiceForCustomerParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GenerateInvoiceForCustomerParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GenerateInvoiceForCustomerParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "datetime", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}

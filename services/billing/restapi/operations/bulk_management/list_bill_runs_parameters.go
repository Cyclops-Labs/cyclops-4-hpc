// Code generated by go-swagger; DO NOT EDIT.

package bulk_management

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

// NewListBillRunsParams creates a new ListBillRunsParams object
// no default values defined in spec.
func NewListBillRunsParams() ListBillRunsParams {

	return ListBillRunsParams{}
}

// ListBillRunsParams contains all the bound params for the list bill runs operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListBillRuns
type ListBillRunsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Amount of months to have in the report
	  In: query
	*/
	Months *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListBillRunsParams() beforehand.
func (o *ListBillRunsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMonths, qhkMonths, _ := qs.GetOK("months")
	if err := o.bindMonths(qMonths, qhkMonths, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMonths binds and validates parameter Months from query.
func (o *ListBillRunsParams) bindMonths(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("months", "query", "int64", raw)
	}
	o.Months = &value

	return nil
}

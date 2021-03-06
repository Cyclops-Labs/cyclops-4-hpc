// Code generated by go-swagger; DO NOT EDIT.

package price_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewListSkuPricesParams creates a new ListSkuPricesParams object
// no default values defined in spec.
func NewListSkuPricesParams() ListSkuPricesParams {

	return ListSkuPricesParams{}
}

// ListSkuPricesParams contains all the bound params for the list sku prices operation
// typically these are obtained from a http.Request
//
// swagger:parameters listSkuPrices
type ListSkuPricesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListSkuPricesParams() beforehand.
func (o *ListSkuPricesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package price_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListSkuPricesParams creates a new ListSkuPricesParams object
// with the default values initialized.
func NewListSkuPricesParams() *ListSkuPricesParams {

	return &ListSkuPricesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSkuPricesParamsWithTimeout creates a new ListSkuPricesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSkuPricesParamsWithTimeout(timeout time.Duration) *ListSkuPricesParams {

	return &ListSkuPricesParams{

		timeout: timeout,
	}
}

// NewListSkuPricesParamsWithContext creates a new ListSkuPricesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSkuPricesParamsWithContext(ctx context.Context) *ListSkuPricesParams {

	return &ListSkuPricesParams{

		Context: ctx,
	}
}

// NewListSkuPricesParamsWithHTTPClient creates a new ListSkuPricesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSkuPricesParamsWithHTTPClient(client *http.Client) *ListSkuPricesParams {

	return &ListSkuPricesParams{
		HTTPClient: client,
	}
}

/*ListSkuPricesParams contains all the parameters to send to the API endpoint
for the list sku prices operation typically these are written to a http.Request
*/
type ListSkuPricesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sku prices params
func (o *ListSkuPricesParams) WithTimeout(timeout time.Duration) *ListSkuPricesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sku prices params
func (o *ListSkuPricesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sku prices params
func (o *ListSkuPricesParams) WithContext(ctx context.Context) *ListSkuPricesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sku prices params
func (o *ListSkuPricesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sku prices params
func (o *ListSkuPricesParams) WithHTTPClient(client *http.Client) *ListSkuPricesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sku prices params
func (o *ListSkuPricesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSkuPricesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

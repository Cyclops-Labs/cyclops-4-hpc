// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

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

// NewGetInvoiceParams creates a new GetInvoiceParams object
// with the default values initialized.
func NewGetInvoiceParams() *GetInvoiceParams {
	var ()
	return &GetInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceParamsWithTimeout creates a new GetInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceParamsWithTimeout(timeout time.Duration) *GetInvoiceParams {
	var ()
	return &GetInvoiceParams{

		timeout: timeout,
	}
}

// NewGetInvoiceParamsWithContext creates a new GetInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceParamsWithContext(ctx context.Context) *GetInvoiceParams {
	var ()
	return &GetInvoiceParams{

		Context: ctx,
	}
}

// NewGetInvoiceParamsWithHTTPClient creates a new GetInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceParamsWithHTTPClient(client *http.Client) *GetInvoiceParams {
	var ()
	return &GetInvoiceParams{
		HTTPClient: client,
	}
}

/*GetInvoiceParams contains all the parameters to send to the API endpoint
for the get invoice operation typically these are written to a http.Request
*/
type GetInvoiceParams struct {

	/*ID
	  Id of the invoice to be checked

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invoice params
func (o *GetInvoiceParams) WithTimeout(timeout time.Duration) *GetInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice params
func (o *GetInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice params
func (o *GetInvoiceParams) WithContext(ctx context.Context) *GetInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice params
func (o *GetInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice params
func (o *GetInvoiceParams) WithHTTPClient(client *http.Client) *GetInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice params
func (o *GetInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get invoice params
func (o *GetInvoiceParams) WithID(id strfmt.UUID) *GetInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get invoice params
func (o *GetInvoiceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

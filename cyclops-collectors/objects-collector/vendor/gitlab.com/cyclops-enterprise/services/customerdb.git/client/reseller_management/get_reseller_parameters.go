// Code generated by go-swagger; DO NOT EDIT.

package reseller_management

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

// NewGetResellerParams creates a new GetResellerParams object
// with the default values initialized.
func NewGetResellerParams() *GetResellerParams {
	var ()
	return &GetResellerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResellerParamsWithTimeout creates a new GetResellerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResellerParamsWithTimeout(timeout time.Duration) *GetResellerParams {
	var ()
	return &GetResellerParams{

		timeout: timeout,
	}
}

// NewGetResellerParamsWithContext creates a new GetResellerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResellerParamsWithContext(ctx context.Context) *GetResellerParams {
	var ()
	return &GetResellerParams{

		Context: ctx,
	}
}

// NewGetResellerParamsWithHTTPClient creates a new GetResellerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResellerParamsWithHTTPClient(client *http.Client) *GetResellerParams {
	var ()
	return &GetResellerParams{
		HTTPClient: client,
	}
}

/*GetResellerParams contains all the parameters to send to the API endpoint
for the get reseller operation typically these are written to a http.Request
*/
type GetResellerParams struct {

	/*ID
	  Id of the reseller to be retrieved

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reseller params
func (o *GetResellerParams) WithTimeout(timeout time.Duration) *GetResellerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reseller params
func (o *GetResellerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reseller params
func (o *GetResellerParams) WithContext(ctx context.Context) *GetResellerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reseller params
func (o *GetResellerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reseller params
func (o *GetResellerParams) WithHTTPClient(client *http.Client) *GetResellerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reseller params
func (o *GetResellerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get reseller params
func (o *GetResellerParams) WithID(id string) *GetResellerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get reseller params
func (o *GetResellerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetResellerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

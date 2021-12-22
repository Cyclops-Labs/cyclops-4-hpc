// Code generated by go-swagger; DO NOT EDIT.

package account_management

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

// NewGetAccountStatusParams creates a new GetAccountStatusParams object
// with the default values initialized.
func NewGetAccountStatusParams() *GetAccountStatusParams {
	var ()
	return &GetAccountStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountStatusParamsWithTimeout creates a new GetAccountStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountStatusParamsWithTimeout(timeout time.Duration) *GetAccountStatusParams {
	var ()
	return &GetAccountStatusParams{

		timeout: timeout,
	}
}

// NewGetAccountStatusParamsWithContext creates a new GetAccountStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountStatusParamsWithContext(ctx context.Context) *GetAccountStatusParams {
	var ()
	return &GetAccountStatusParams{

		Context: ctx,
	}
}

// NewGetAccountStatusParamsWithHTTPClient creates a new GetAccountStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountStatusParamsWithHTTPClient(client *http.Client) *GetAccountStatusParams {
	var ()
	return &GetAccountStatusParams{
		HTTPClient: client,
	}
}

/*GetAccountStatusParams contains all the parameters to send to the API endpoint
for the get account status operation typically these are written to a http.Request
*/
type GetAccountStatusParams struct {

	/*ID
	  Id of the account to be checked

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account status params
func (o *GetAccountStatusParams) WithTimeout(timeout time.Duration) *GetAccountStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account status params
func (o *GetAccountStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account status params
func (o *GetAccountStatusParams) WithContext(ctx context.Context) *GetAccountStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account status params
func (o *GetAccountStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account status params
func (o *GetAccountStatusParams) WithHTTPClient(client *http.Client) *GetAccountStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account status params
func (o *GetAccountStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get account status params
func (o *GetAccountStatusParams) WithID(id string) *GetAccountStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get account status params
func (o *GetAccountStatusParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
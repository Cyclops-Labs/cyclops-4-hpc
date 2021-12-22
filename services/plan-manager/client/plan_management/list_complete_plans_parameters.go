// Code generated by go-swagger; DO NOT EDIT.

package plan_management

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

// NewListCompletePlansParams creates a new ListCompletePlansParams object
// with the default values initialized.
func NewListCompletePlansParams() *ListCompletePlansParams {

	return &ListCompletePlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCompletePlansParamsWithTimeout creates a new ListCompletePlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCompletePlansParamsWithTimeout(timeout time.Duration) *ListCompletePlansParams {

	return &ListCompletePlansParams{

		timeout: timeout,
	}
}

// NewListCompletePlansParamsWithContext creates a new ListCompletePlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCompletePlansParamsWithContext(ctx context.Context) *ListCompletePlansParams {

	return &ListCompletePlansParams{

		Context: ctx,
	}
}

// NewListCompletePlansParamsWithHTTPClient creates a new ListCompletePlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCompletePlansParamsWithHTTPClient(client *http.Client) *ListCompletePlansParams {

	return &ListCompletePlansParams{
		HTTPClient: client,
	}
}

/*ListCompletePlansParams contains all the parameters to send to the API endpoint
for the list complete plans operation typically these are written to a http.Request
*/
type ListCompletePlansParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list complete plans params
func (o *ListCompletePlansParams) WithTimeout(timeout time.Duration) *ListCompletePlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list complete plans params
func (o *ListCompletePlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list complete plans params
func (o *ListCompletePlansParams) WithContext(ctx context.Context) *ListCompletePlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list complete plans params
func (o *ListCompletePlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list complete plans params
func (o *ListCompletePlansParams) WithHTTPClient(client *http.Client) *ListCompletePlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list complete plans params
func (o *ListCompletePlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCompletePlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
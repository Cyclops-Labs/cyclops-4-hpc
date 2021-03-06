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

// NewListPlansParams creates a new ListPlansParams object
// with the default values initialized.
func NewListPlansParams() *ListPlansParams {

	return &ListPlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPlansParamsWithTimeout creates a new ListPlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPlansParamsWithTimeout(timeout time.Duration) *ListPlansParams {

	return &ListPlansParams{

		timeout: timeout,
	}
}

// NewListPlansParamsWithContext creates a new ListPlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPlansParamsWithContext(ctx context.Context) *ListPlansParams {

	return &ListPlansParams{

		Context: ctx,
	}
}

// NewListPlansParamsWithHTTPClient creates a new ListPlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPlansParamsWithHTTPClient(client *http.Client) *ListPlansParams {

	return &ListPlansParams{
		HTTPClient: client,
	}
}

/*ListPlansParams contains all the parameters to send to the API endpoint
for the list plans operation typically these are written to a http.Request
*/
type ListPlansParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list plans params
func (o *ListPlansParams) WithTimeout(timeout time.Duration) *ListPlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list plans params
func (o *ListPlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list plans params
func (o *ListPlansParams) WithContext(ctx context.Context) *ListPlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list plans params
func (o *ListPlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list plans params
func (o *ListPlansParams) WithHTTPClient(client *http.Client) *ListPlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list plans params
func (o *ListPlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

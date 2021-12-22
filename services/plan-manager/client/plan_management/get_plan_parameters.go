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

// NewGetPlanParams creates a new GetPlanParams object
// with the default values initialized.
func NewGetPlanParams() *GetPlanParams {
	var ()
	return &GetPlanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlanParamsWithTimeout creates a new GetPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlanParamsWithTimeout(timeout time.Duration) *GetPlanParams {
	var ()
	return &GetPlanParams{

		timeout: timeout,
	}
}

// NewGetPlanParamsWithContext creates a new GetPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlanParamsWithContext(ctx context.Context) *GetPlanParams {
	var ()
	return &GetPlanParams{

		Context: ctx,
	}
}

// NewGetPlanParamsWithHTTPClient creates a new GetPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlanParamsWithHTTPClient(client *http.Client) *GetPlanParams {
	var ()
	return &GetPlanParams{
		HTTPClient: client,
	}
}

/*GetPlanParams contains all the parameters to send to the API endpoint
for the get plan operation typically these are written to a http.Request
*/
type GetPlanParams struct {

	/*ID
	  Id of plan to be obtained

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plan params
func (o *GetPlanParams) WithTimeout(timeout time.Duration) *GetPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plan params
func (o *GetPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plan params
func (o *GetPlanParams) WithContext(ctx context.Context) *GetPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plan params
func (o *GetPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plan params
func (o *GetPlanParams) WithHTTPClient(client *http.Client) *GetPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plan params
func (o *GetPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get plan params
func (o *GetPlanParams) WithID(id string) *GetPlanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get plan params
func (o *GetPlanParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
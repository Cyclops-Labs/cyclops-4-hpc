// Code generated by go-swagger; DO NOT EDIT.

package cycle_management

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

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// NewUpdateCycleParams creates a new UpdateCycleParams object
// with the default values initialized.
func NewUpdateCycleParams() *UpdateCycleParams {
	var ()
	return &UpdateCycleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCycleParamsWithTimeout creates a new UpdateCycleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCycleParamsWithTimeout(timeout time.Duration) *UpdateCycleParams {
	var ()
	return &UpdateCycleParams{

		timeout: timeout,
	}
}

// NewUpdateCycleParamsWithContext creates a new UpdateCycleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCycleParamsWithContext(ctx context.Context) *UpdateCycleParams {
	var ()
	return &UpdateCycleParams{

		Context: ctx,
	}
}

// NewUpdateCycleParamsWithHTTPClient creates a new UpdateCycleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCycleParamsWithHTTPClient(client *http.Client) *UpdateCycleParams {
	var ()
	return &UpdateCycleParams{
		HTTPClient: client,
	}
}

/*UpdateCycleParams contains all the parameters to send to the API endpoint
for the update cycle operation typically these are written to a http.Request
*/
type UpdateCycleParams struct {

	/*Cycle
	  updated cycle containing all parameters except id

	*/
	Cycle *models.Cycle
	/*ID
	  Id of cycle to be updated

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cycle params
func (o *UpdateCycleParams) WithTimeout(timeout time.Duration) *UpdateCycleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cycle params
func (o *UpdateCycleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cycle params
func (o *UpdateCycleParams) WithContext(ctx context.Context) *UpdateCycleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cycle params
func (o *UpdateCycleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cycle params
func (o *UpdateCycleParams) WithHTTPClient(client *http.Client) *UpdateCycleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cycle params
func (o *UpdateCycleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCycle adds the cycle to the update cycle params
func (o *UpdateCycleParams) WithCycle(cycle *models.Cycle) *UpdateCycleParams {
	o.SetCycle(cycle)
	return o
}

// SetCycle adds the cycle to the update cycle params
func (o *UpdateCycleParams) SetCycle(cycle *models.Cycle) {
	o.Cycle = cycle
}

// WithID adds the id to the update cycle params
func (o *UpdateCycleParams) WithID(id string) *UpdateCycleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update cycle params
func (o *UpdateCycleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCycleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cycle != nil {
		if err := r.SetBodyParam(o.Cycle); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

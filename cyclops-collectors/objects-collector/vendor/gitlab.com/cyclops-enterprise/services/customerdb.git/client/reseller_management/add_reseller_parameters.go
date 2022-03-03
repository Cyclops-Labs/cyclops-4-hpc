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

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
)

// NewAddResellerParams creates a new AddResellerParams object
// with the default values initialized.
func NewAddResellerParams() *AddResellerParams {
	var ()
	return &AddResellerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddResellerParamsWithTimeout creates a new AddResellerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddResellerParamsWithTimeout(timeout time.Duration) *AddResellerParams {
	var ()
	return &AddResellerParams{

		timeout: timeout,
	}
}

// NewAddResellerParamsWithContext creates a new AddResellerParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddResellerParamsWithContext(ctx context.Context) *AddResellerParams {
	var ()
	return &AddResellerParams{

		Context: ctx,
	}
}

// NewAddResellerParamsWithHTTPClient creates a new AddResellerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddResellerParamsWithHTTPClient(client *http.Client) *AddResellerParams {
	var ()
	return &AddResellerParams{
		HTTPClient: client,
	}
}

/*AddResellerParams contains all the parameters to send to the API endpoint
for the add reseller operation typically these are written to a http.Request
*/
type AddResellerParams struct {

	/*Reseller
	  Reseller to be added

	*/
	Reseller *models.Reseller

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add reseller params
func (o *AddResellerParams) WithTimeout(timeout time.Duration) *AddResellerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add reseller params
func (o *AddResellerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add reseller params
func (o *AddResellerParams) WithContext(ctx context.Context) *AddResellerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add reseller params
func (o *AddResellerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add reseller params
func (o *AddResellerParams) WithHTTPClient(client *http.Client) *AddResellerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add reseller params
func (o *AddResellerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReseller adds the reseller to the add reseller params
func (o *AddResellerParams) WithReseller(reseller *models.Reseller) *AddResellerParams {
	o.SetReseller(reseller)
	return o
}

// SetReseller adds the reseller to the add reseller params
func (o *AddResellerParams) SetReseller(reseller *models.Reseller) {
	o.Reseller = reseller
}

// WriteToRequest writes these params to a swagger request
func (o *AddResellerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Reseller != nil {
		if err := r.SetBodyParam(o.Reseller); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

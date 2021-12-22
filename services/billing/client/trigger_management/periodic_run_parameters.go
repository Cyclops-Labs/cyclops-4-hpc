// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

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

// NewPeriodicRunParams creates a new PeriodicRunParams object
// with the default values initialized.
func NewPeriodicRunParams() *PeriodicRunParams {
	var ()
	return &PeriodicRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeriodicRunParamsWithTimeout creates a new PeriodicRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeriodicRunParamsWithTimeout(timeout time.Duration) *PeriodicRunParams {
	var ()
	return &PeriodicRunParams{

		timeout: timeout,
	}
}

// NewPeriodicRunParamsWithContext creates a new PeriodicRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeriodicRunParamsWithContext(ctx context.Context) *PeriodicRunParams {
	var ()
	return &PeriodicRunParams{

		Context: ctx,
	}
}

// NewPeriodicRunParamsWithHTTPClient creates a new PeriodicRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeriodicRunParamsWithHTTPClient(client *http.Client) *PeriodicRunParams {
	var ()
	return &PeriodicRunParams{
		HTTPClient: client,
	}
}

/*PeriodicRunParams contains all the parameters to send to the API endpoint
for the periodic run operation typically these are written to a http.Request
*/
type PeriodicRunParams struct {

	/*Today
	  Datetime to override the time.now() to simulate other days

	*/
	Today *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the periodic run params
func (o *PeriodicRunParams) WithTimeout(timeout time.Duration) *PeriodicRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the periodic run params
func (o *PeriodicRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the periodic run params
func (o *PeriodicRunParams) WithContext(ctx context.Context) *PeriodicRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the periodic run params
func (o *PeriodicRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the periodic run params
func (o *PeriodicRunParams) WithHTTPClient(client *http.Client) *PeriodicRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the periodic run params
func (o *PeriodicRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToday adds the today to the periodic run params
func (o *PeriodicRunParams) WithToday(today *strfmt.DateTime) *PeriodicRunParams {
	o.SetToday(today)
	return o
}

// SetToday adds the today to the periodic run params
func (o *PeriodicRunParams) SetToday(today *strfmt.DateTime) {
	o.Today = today
}

// WriteToRequest writes these params to a swagger request
func (o *PeriodicRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Today != nil {

		// query param today
		var qrToday strfmt.DateTime
		if o.Today != nil {
			qrToday = *o.Today
		}
		qToday := qrToday.String()
		if qToday != "" {
			if err := r.SetQueryParam("today", qToday); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
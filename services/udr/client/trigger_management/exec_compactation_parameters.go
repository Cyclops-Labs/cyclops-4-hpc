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
	"github.com/go-openapi/swag"
)

// NewExecCompactationParams creates a new ExecCompactationParams object
// with the default values initialized.
func NewExecCompactationParams() *ExecCompactationParams {
	var ()
	return &ExecCompactationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecCompactationParamsWithTimeout creates a new ExecCompactationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecCompactationParamsWithTimeout(timeout time.Duration) *ExecCompactationParams {
	var ()
	return &ExecCompactationParams{

		timeout: timeout,
	}
}

// NewExecCompactationParamsWithContext creates a new ExecCompactationParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecCompactationParamsWithContext(ctx context.Context) *ExecCompactationParams {
	var ()
	return &ExecCompactationParams{

		Context: ctx,
	}
}

// NewExecCompactationParamsWithHTTPClient creates a new ExecCompactationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecCompactationParamsWithHTTPClient(client *http.Client) *ExecCompactationParams {
	var ()
	return &ExecCompactationParams{
		HTTPClient: client,
	}
}

/*ExecCompactationParams contains all the parameters to send to the API endpoint
for the exec compactation operation typically these are written to a http.Request
*/
type ExecCompactationParams struct {

	/*FastMode
	  Switch for using 15m boundaries instead of 8h

	*/
	FastMode *bool
	/*From
	  Datetime from which to get the usage report

	*/
	From *strfmt.DateTime
	/*To
	  Datetime until which to get the usage report

	*/
	To *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the exec compactation params
func (o *ExecCompactationParams) WithTimeout(timeout time.Duration) *ExecCompactationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec compactation params
func (o *ExecCompactationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec compactation params
func (o *ExecCompactationParams) WithContext(ctx context.Context) *ExecCompactationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec compactation params
func (o *ExecCompactationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec compactation params
func (o *ExecCompactationParams) WithHTTPClient(client *http.Client) *ExecCompactationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec compactation params
func (o *ExecCompactationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFastMode adds the fastMode to the exec compactation params
func (o *ExecCompactationParams) WithFastMode(fastMode *bool) *ExecCompactationParams {
	o.SetFastMode(fastMode)
	return o
}

// SetFastMode adds the fastMode to the exec compactation params
func (o *ExecCompactationParams) SetFastMode(fastMode *bool) {
	o.FastMode = fastMode
}

// WithFrom adds the from to the exec compactation params
func (o *ExecCompactationParams) WithFrom(from *strfmt.DateTime) *ExecCompactationParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the exec compactation params
func (o *ExecCompactationParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithTo adds the to to the exec compactation params
func (o *ExecCompactationParams) WithTo(to *strfmt.DateTime) *ExecCompactationParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the exec compactation params
func (o *ExecCompactationParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *ExecCompactationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FastMode != nil {

		// query param fast_mode
		var qrFastMode bool
		if o.FastMode != nil {
			qrFastMode = *o.FastMode
		}
		qFastMode := swag.FormatBool(qrFastMode)
		if qFastMode != "" {
			if err := r.SetQueryParam("fast_mode", qFastMode); err != nil {
				return err
			}
		}

	}

	if o.From != nil {

		// query param from
		var qrFrom strfmt.DateTime
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom.String()
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.To != nil {

		// query param to
		var qrTo strfmt.DateTime
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo.String()
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

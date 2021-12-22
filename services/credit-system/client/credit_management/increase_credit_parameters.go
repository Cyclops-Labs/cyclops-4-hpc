// Code generated by go-swagger; DO NOT EDIT.

package credit_management

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

// NewIncreaseCreditParams creates a new IncreaseCreditParams object
// with the default values initialized.
func NewIncreaseCreditParams() *IncreaseCreditParams {
	var ()
	return &IncreaseCreditParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIncreaseCreditParamsWithTimeout creates a new IncreaseCreditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIncreaseCreditParamsWithTimeout(timeout time.Duration) *IncreaseCreditParams {
	var ()
	return &IncreaseCreditParams{

		timeout: timeout,
	}
}

// NewIncreaseCreditParamsWithContext creates a new IncreaseCreditParams object
// with the default values initialized, and the ability to set a context for a request
func NewIncreaseCreditParamsWithContext(ctx context.Context) *IncreaseCreditParams {
	var ()
	return &IncreaseCreditParams{

		Context: ctx,
	}
}

// NewIncreaseCreditParamsWithHTTPClient creates a new IncreaseCreditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIncreaseCreditParamsWithHTTPClient(client *http.Client) *IncreaseCreditParams {
	var ()
	return &IncreaseCreditParams{
		HTTPClient: client,
	}
}

/*IncreaseCreditParams contains all the parameters to send to the API endpoint
for the increase credit operation typically these are written to a http.Request
*/
type IncreaseCreditParams struct {

	/*Amount
	  Amount to be inccreased

	*/
	Amount float64
	/*ID
	  Id of the account to be checked

	*/
	ID string
	/*Medium
	  Medium (cash/credit) to be used in the accounting

	*/
	Medium string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the increase credit params
func (o *IncreaseCreditParams) WithTimeout(timeout time.Duration) *IncreaseCreditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the increase credit params
func (o *IncreaseCreditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the increase credit params
func (o *IncreaseCreditParams) WithContext(ctx context.Context) *IncreaseCreditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the increase credit params
func (o *IncreaseCreditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the increase credit params
func (o *IncreaseCreditParams) WithHTTPClient(client *http.Client) *IncreaseCreditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the increase credit params
func (o *IncreaseCreditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the increase credit params
func (o *IncreaseCreditParams) WithAmount(amount float64) *IncreaseCreditParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the increase credit params
func (o *IncreaseCreditParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithID adds the id to the increase credit params
func (o *IncreaseCreditParams) WithID(id string) *IncreaseCreditParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the increase credit params
func (o *IncreaseCreditParams) SetID(id string) {
	o.ID = id
}

// WithMedium adds the medium to the increase credit params
func (o *IncreaseCreditParams) WithMedium(medium string) *IncreaseCreditParams {
	o.SetMedium(medium)
	return o
}

// SetMedium adds the medium to the increase credit params
func (o *IncreaseCreditParams) SetMedium(medium string) {
	o.Medium = medium
}

// WriteToRequest writes these params to a swagger request
func (o *IncreaseCreditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatFloat64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param medium
	if err := r.SetPathParam("medium", o.Medium); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

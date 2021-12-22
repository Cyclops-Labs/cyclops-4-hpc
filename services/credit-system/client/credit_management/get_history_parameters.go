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

// NewGetHistoryParams creates a new GetHistoryParams object
// with the default values initialized.
func NewGetHistoryParams() *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryParamsWithTimeout creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryParamsWithTimeout(timeout time.Duration) *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		timeout: timeout,
	}
}

// NewGetHistoryParamsWithContext creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryParamsWithContext(ctx context.Context) *GetHistoryParams {
	var ()
	return &GetHistoryParams{

		Context: ctx,
	}
}

// NewGetHistoryParamsWithHTTPClient creates a new GetHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryParamsWithHTTPClient(client *http.Client) *GetHistoryParams {
	var ()
	return &GetHistoryParams{
		HTTPClient: client,
	}
}

/*GetHistoryParams contains all the parameters to send to the API endpoint
for the get history operation typically these are written to a http.Request
*/
type GetHistoryParams struct {

	/*FilterSystem
	  Boolean variable to control if the system consumptions have to be listed or not

	*/
	FilterSystem *bool
	/*ID
	  Id of the account to get the history

	*/
	ID string
	/*Medium
	  Medium (cash/credit) to be used as filter

	*/
	Medium *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history params
func (o *GetHistoryParams) WithTimeout(timeout time.Duration) *GetHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history params
func (o *GetHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history params
func (o *GetHistoryParams) WithContext(ctx context.Context) *GetHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history params
func (o *GetHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) WithHTTPClient(client *http.Client) *GetHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history params
func (o *GetHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterSystem adds the filterSystem to the get history params
func (o *GetHistoryParams) WithFilterSystem(filterSystem *bool) *GetHistoryParams {
	o.SetFilterSystem(filterSystem)
	return o
}

// SetFilterSystem adds the filterSystem to the get history params
func (o *GetHistoryParams) SetFilterSystem(filterSystem *bool) {
	o.FilterSystem = filterSystem
}

// WithID adds the id to the get history params
func (o *GetHistoryParams) WithID(id string) *GetHistoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get history params
func (o *GetHistoryParams) SetID(id string) {
	o.ID = id
}

// WithMedium adds the medium to the get history params
func (o *GetHistoryParams) WithMedium(medium *string) *GetHistoryParams {
	o.SetMedium(medium)
	return o
}

// SetMedium adds the medium to the get history params
func (o *GetHistoryParams) SetMedium(medium *string) {
	o.Medium = medium
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterSystem != nil {

		// query param filterSystem
		var qrFilterSystem bool
		if o.FilterSystem != nil {
			qrFilterSystem = *o.FilterSystem
		}
		qFilterSystem := swag.FormatBool(qrFilterSystem)
		if qFilterSystem != "" {
			if err := r.SetQueryParam("filterSystem", qFilterSystem); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Medium != nil {

		// query param medium
		var qrMedium string
		if o.Medium != nil {
			qrMedium = *o.Medium
		}
		qMedium := qrMedium
		if qMedium != "" {
			if err := r.SetQueryParam("medium", qMedium); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
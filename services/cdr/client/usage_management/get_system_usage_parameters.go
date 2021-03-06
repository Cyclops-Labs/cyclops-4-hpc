// Code generated by go-swagger; DO NOT EDIT.

package usage_management

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

// NewGetSystemUsageParams creates a new GetSystemUsageParams object
// with the default values initialized.
func NewGetSystemUsageParams() *GetSystemUsageParams {
	var ()
	return &GetSystemUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemUsageParamsWithTimeout creates a new GetSystemUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemUsageParamsWithTimeout(timeout time.Duration) *GetSystemUsageParams {
	var ()
	return &GetSystemUsageParams{

		timeout: timeout,
	}
}

// NewGetSystemUsageParamsWithContext creates a new GetSystemUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemUsageParamsWithContext(ctx context.Context) *GetSystemUsageParams {
	var ()
	return &GetSystemUsageParams{

		Context: ctx,
	}
}

// NewGetSystemUsageParamsWithHTTPClient creates a new GetSystemUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemUsageParamsWithHTTPClient(client *http.Client) *GetSystemUsageParams {
	var ()
	return &GetSystemUsageParams{
		HTTPClient: client,
	}
}

/*GetSystemUsageParams contains all the parameters to send to the API endpoint
for the get system usage operation typically these are written to a http.Request
*/
type GetSystemUsageParams struct {

	/*From
	  Datetime from which to get the usage report

	*/
	From *strfmt.DateTime
	/*Idlist
	  List of ids to be queried

	*/
	Idlist *string
	/*Metric
	  Metric(s) to get the usage report

	*/
	Metric *string
	/*To
	  Datetime until which to get the usage report

	*/
	To *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system usage params
func (o *GetSystemUsageParams) WithTimeout(timeout time.Duration) *GetSystemUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system usage params
func (o *GetSystemUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system usage params
func (o *GetSystemUsageParams) WithContext(ctx context.Context) *GetSystemUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system usage params
func (o *GetSystemUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system usage params
func (o *GetSystemUsageParams) WithHTTPClient(client *http.Client) *GetSystemUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system usage params
func (o *GetSystemUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get system usage params
func (o *GetSystemUsageParams) WithFrom(from *strfmt.DateTime) *GetSystemUsageParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get system usage params
func (o *GetSystemUsageParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithIdlist adds the idlist to the get system usage params
func (o *GetSystemUsageParams) WithIdlist(idlist *string) *GetSystemUsageParams {
	o.SetIdlist(idlist)
	return o
}

// SetIdlist adds the idlist to the get system usage params
func (o *GetSystemUsageParams) SetIdlist(idlist *string) {
	o.Idlist = idlist
}

// WithMetric adds the metric to the get system usage params
func (o *GetSystemUsageParams) WithMetric(metric *string) *GetSystemUsageParams {
	o.SetMetric(metric)
	return o
}

// SetMetric adds the metric to the get system usage params
func (o *GetSystemUsageParams) SetMetric(metric *string) {
	o.Metric = metric
}

// WithTo adds the to to the get system usage params
func (o *GetSystemUsageParams) WithTo(to *strfmt.DateTime) *GetSystemUsageParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get system usage params
func (o *GetSystemUsageParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Idlist != nil {

		// query param idlist
		var qrIdlist string
		if o.Idlist != nil {
			qrIdlist = *o.Idlist
		}
		qIdlist := qrIdlist
		if qIdlist != "" {
			if err := r.SetQueryParam("idlist", qIdlist); err != nil {
				return err
			}
		}

	}

	if o.Metric != nil {

		// query param metric
		var qrMetric string
		if o.Metric != nil {
			qrMetric = *o.Metric
		}
		qMetric := qrMetric
		if qMetric != "" {
			if err := r.SetQueryParam("metric", qMetric); err != nil {
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

// Code generated by go-swagger; DO NOT EDIT.

package sync_management

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

// NewSyncFlavorsParams creates a new SyncFlavorsParams object
// with the default values initialized.
func NewSyncFlavorsParams() *SyncFlavorsParams {

	return &SyncFlavorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncFlavorsParamsWithTimeout creates a new SyncFlavorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncFlavorsParamsWithTimeout(timeout time.Duration) *SyncFlavorsParams {

	return &SyncFlavorsParams{

		timeout: timeout,
	}
}

// NewSyncFlavorsParamsWithContext creates a new SyncFlavorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncFlavorsParamsWithContext(ctx context.Context) *SyncFlavorsParams {

	return &SyncFlavorsParams{

		Context: ctx,
	}
}

// NewSyncFlavorsParamsWithHTTPClient creates a new SyncFlavorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncFlavorsParamsWithHTTPClient(client *http.Client) *SyncFlavorsParams {

	return &SyncFlavorsParams{
		HTTPClient: client,
	}
}

/*SyncFlavorsParams contains all the parameters to send to the API endpoint
for the sync flavors operation typically these are written to a http.Request
*/
type SyncFlavorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync flavors params
func (o *SyncFlavorsParams) WithTimeout(timeout time.Duration) *SyncFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync flavors params
func (o *SyncFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync flavors params
func (o *SyncFlavorsParams) WithContext(ctx context.Context) *SyncFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync flavors params
func (o *SyncFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync flavors params
func (o *SyncFlavorsParams) WithHTTPClient(client *http.Client) *SyncFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync flavors params
func (o *SyncFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SyncFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

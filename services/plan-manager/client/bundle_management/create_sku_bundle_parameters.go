// Code generated by go-swagger; DO NOT EDIT.

package bundle_management

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

// NewCreateSkuBundleParams creates a new CreateSkuBundleParams object
// with the default values initialized.
func NewCreateSkuBundleParams() *CreateSkuBundleParams {
	var ()
	return &CreateSkuBundleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSkuBundleParamsWithTimeout creates a new CreateSkuBundleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSkuBundleParamsWithTimeout(timeout time.Duration) *CreateSkuBundleParams {
	var ()
	return &CreateSkuBundleParams{

		timeout: timeout,
	}
}

// NewCreateSkuBundleParamsWithContext creates a new CreateSkuBundleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSkuBundleParamsWithContext(ctx context.Context) *CreateSkuBundleParams {
	var ()
	return &CreateSkuBundleParams{

		Context: ctx,
	}
}

// NewCreateSkuBundleParamsWithHTTPClient creates a new CreateSkuBundleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSkuBundleParamsWithHTTPClient(client *http.Client) *CreateSkuBundleParams {
	var ()
	return &CreateSkuBundleParams{
		HTTPClient: client,
	}
}

/*CreateSkuBundleParams contains all the parameters to send to the API endpoint
for the create sku bundle operation typically these are written to a http.Request
*/
type CreateSkuBundleParams struct {

	/*Bundle
	  SKU bundle to be added

	*/
	Bundle *models.SkuBundle

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sku bundle params
func (o *CreateSkuBundleParams) WithTimeout(timeout time.Duration) *CreateSkuBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sku bundle params
func (o *CreateSkuBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sku bundle params
func (o *CreateSkuBundleParams) WithContext(ctx context.Context) *CreateSkuBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sku bundle params
func (o *CreateSkuBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sku bundle params
func (o *CreateSkuBundleParams) WithHTTPClient(client *http.Client) *CreateSkuBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sku bundle params
func (o *CreateSkuBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundle adds the bundle to the create sku bundle params
func (o *CreateSkuBundleParams) WithBundle(bundle *models.SkuBundle) *CreateSkuBundleParams {
	o.SetBundle(bundle)
	return o
}

// SetBundle adds the bundle to the create sku bundle params
func (o *CreateSkuBundleParams) SetBundle(bundle *models.SkuBundle) {
	o.Bundle = bundle
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSkuBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bundle != nil {
		if err := r.SetBodyParam(o.Bundle); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

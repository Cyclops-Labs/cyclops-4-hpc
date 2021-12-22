// Code generated by go-swagger; DO NOT EDIT.

package bundle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// NewUpdateSkuBundleParams creates a new UpdateSkuBundleParams object
// no default values defined in spec.
func NewUpdateSkuBundleParams() UpdateSkuBundleParams {

	return UpdateSkuBundleParams{}
}

// UpdateSkuBundleParams contains all the bound params for the update sku bundle operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateSkuBundle
type UpdateSkuBundleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*updated sku bundle containing all parameters except id
	  Required: true
	  In: body
	*/
	Bundle *models.SkuBundle
	/*Id of sku bundle to be obtained
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateSkuBundleParams() beforehand.
func (o *UpdateSkuBundleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.SkuBundle
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("bundle", "body", ""))
			} else {
				res = append(res, errors.NewParseError("bundle", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Bundle = &body
			}
		}
	} else {
		res = append(res, errors.Required("bundle", "body", ""))
	}
	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UpdateSkuBundleParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}
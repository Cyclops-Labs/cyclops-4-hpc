// Code generated by go-swagger; DO NOT EDIT.

package event_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/models"
)

// NewAddEventParams creates a new AddEventParams object
// no default values defined in spec.
func NewAddEventParams() AddEventParams {

	return AddEventParams{}
}

// AddEventParams contains all the bound params for the add event operation
// typically these are obtained from a http.Request
//
// swagger:parameters addEvent
type AddEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Event to be added to the system
	  Required: true
	  In: body
	*/
	Event *models.Event
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddEventParams() beforehand.
func (o *AddEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Event
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("event", "body", ""))
			} else {
				res = append(res, errors.NewParseError("event", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Event = &body
			}
		}
	} else {
		res = append(res, errors.Required("event", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
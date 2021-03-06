// Code generated by go-swagger; DO NOT EDIT.

package event_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddEventHandlerFunc turns a function with the right signature into a add event handler
type AddEventHandlerFunc func(AddEventParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddEventHandlerFunc) Handle(params AddEventParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddEventHandler interface for that can handle valid add event params
type AddEventHandler interface {
	Handle(AddEventParams, interface{}) middleware.Responder
}

// NewAddEvent creates a new http.Handler for the add event operation
func NewAddEvent(ctx *middleware.Context, handler AddEventHandler) *AddEvent {
	return &AddEvent{Context: ctx, Handler: handler}
}

/*AddEvent swagger:route POST /event eventManagement addEvent

Takes into the system the provided event

*/
type AddEvent struct {
	Context *middleware.Context
	Handler AddEventHandler
}

func (o *AddEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddEventParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

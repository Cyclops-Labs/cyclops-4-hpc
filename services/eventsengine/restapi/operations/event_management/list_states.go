// Code generated by go-swagger; DO NOT EDIT.

package event_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStatesHandlerFunc turns a function with the right signature into a list states handler
type ListStatesHandlerFunc func(ListStatesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStatesHandlerFunc) Handle(params ListStatesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListStatesHandler interface for that can handle valid list states params
type ListStatesHandler interface {
	Handle(ListStatesParams, interface{}) middleware.Responder
}

// NewListStates creates a new http.Handler for the list states operation
func NewListStates(ctx *middleware.Context, handler ListStatesHandler) *ListStates {
	return &ListStates{Context: ctx, Handler: handler}
}

/*ListStates swagger:route GET /event/status eventManagement listStates

Provides the list of states in not terminated state

*/
type ListStates struct {
	Context *middleware.Context
	Handler ListStatesHandler
}

func (o *ListStates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStatesParams()

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

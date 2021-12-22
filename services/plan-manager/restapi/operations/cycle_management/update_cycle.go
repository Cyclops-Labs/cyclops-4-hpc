// Code generated by go-swagger; DO NOT EDIT.

package cycle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateCycleHandlerFunc turns a function with the right signature into a update cycle handler
type UpdateCycleHandlerFunc func(UpdateCycleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCycleHandlerFunc) Handle(params UpdateCycleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateCycleHandler interface for that can handle valid update cycle params
type UpdateCycleHandler interface {
	Handle(UpdateCycleParams, interface{}) middleware.Responder
}

// NewUpdateCycle creates a new http.Handler for the update cycle operation
func NewUpdateCycle(ctx *middleware.Context, handler UpdateCycleHandler) *UpdateCycle {
	return &UpdateCycle{Context: ctx, Handler: handler}
}

/*UpdateCycle swagger:route PUT /cycle/{id} cycleManagement updateCycle

Update specific cycle

Update cycle with given id

*/
type UpdateCycle struct {
	Context *middleware.Context
	Handler UpdateCycleHandler
}

func (o *UpdateCycle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateCycleParams()

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

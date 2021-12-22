// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ExecCompactationHandlerFunc turns a function with the right signature into a exec compactation handler
type ExecCompactationHandlerFunc func(ExecCompactationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ExecCompactationHandlerFunc) Handle(params ExecCompactationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ExecCompactationHandler interface for that can handle valid exec compactation params
type ExecCompactationHandler interface {
	Handle(ExecCompactationParams, interface{}) middleware.Responder
}

// NewExecCompactation creates a new http.Handler for the exec compactation operation
func NewExecCompactation(ctx *middleware.Context, handler ExecCompactationHandler) *ExecCompactation {
	return &ExecCompactation{Context: ctx, Handler: handler}
}

/*ExecCompactation swagger:route GET /trigger/compact triggerManagement execCompactation

Compactation task trigger

*/
type ExecCompactation struct {
	Context *middleware.Context
	Handler ExecCompactationHandler
}

func (o *ExecCompactation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExecCompactationParams()

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
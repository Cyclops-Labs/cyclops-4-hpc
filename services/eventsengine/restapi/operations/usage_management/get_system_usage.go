// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSystemUsageHandlerFunc turns a function with the right signature into a get system usage handler
type GetSystemUsageHandlerFunc func(GetSystemUsageParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSystemUsageHandlerFunc) Handle(params GetSystemUsageParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSystemUsageHandler interface for that can handle valid get system usage params
type GetSystemUsageHandler interface {
	Handle(GetSystemUsageParams, interface{}) middleware.Responder
}

// NewGetSystemUsage creates a new http.Handler for the get system usage operation
func NewGetSystemUsage(ctx *middleware.Context, handler GetSystemUsageHandler) *GetSystemUsage {
	return &GetSystemUsage{Context: ctx, Handler: handler}
}

/*GetSystemUsage swagger:route GET /usage usageManagement getSystemUsage

Generates an aggregated response by account of the usage recorded in the system during the time-window specified

*/
type GetSystemUsage struct {
	Context *middleware.Context
	Handler GetSystemUsageHandler
}

func (o *GetSystemUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSystemUsageParams()

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
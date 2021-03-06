// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsageHandlerFunc turns a function with the right signature into a get usage handler
type GetUsageHandlerFunc func(GetUsageParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsageHandlerFunc) Handle(params GetUsageParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUsageHandler interface for that can handle valid get usage params
type GetUsageHandler interface {
	Handle(GetUsageParams, interface{}) middleware.Responder
}

// NewGetUsage creates a new http.Handler for the get usage operation
func NewGetUsage(ctx *middleware.Context, handler GetUsageHandler) *GetUsage {
	return &GetUsage{Context: ctx, Handler: handler}
}

/*GetUsage swagger:route GET /usage/{id} usageManagement getUsage

Generates an aggregated response of the usage recorded in the system during the time-window specified for the selected account

*/
type GetUsage struct {
	Context *middleware.Context
	Handler GetUsageHandler
}

func (o *GetUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsageParams()

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

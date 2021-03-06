// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsageSummaryHandlerFunc turns a function with the right signature into a get usage summary handler
type GetUsageSummaryHandlerFunc func(GetUsageSummaryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsageSummaryHandlerFunc) Handle(params GetUsageSummaryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUsageSummaryHandler interface for that can handle valid get usage summary params
type GetUsageSummaryHandler interface {
	Handle(GetUsageSummaryParams, interface{}) middleware.Responder
}

// NewGetUsageSummary creates a new http.Handler for the get usage summary operation
func NewGetUsageSummary(ctx *middleware.Context, handler GetUsageSummaryHandler) *GetUsageSummary {
	return &GetUsageSummary{Context: ctx, Handler: handler}
}

/*GetUsageSummary swagger:route GET /usage/summary/{id} usageManagement getUsageSummary

Summary report meant for the UI for the resources linked to the ResellerID provided within the specified time window

*/
type GetUsageSummary struct {
	Context *middleware.Context
	Handler GetUsageSummaryHandler
}

func (o *GetUsageSummary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsageSummaryParams()

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

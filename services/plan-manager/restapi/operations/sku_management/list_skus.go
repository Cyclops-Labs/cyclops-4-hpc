// Code generated by go-swagger; DO NOT EDIT.

package sku_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListSkusHandlerFunc turns a function with the right signature into a list skus handler
type ListSkusHandlerFunc func(ListSkusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSkusHandlerFunc) Handle(params ListSkusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListSkusHandler interface for that can handle valid list skus params
type ListSkusHandler interface {
	Handle(ListSkusParams, interface{}) middleware.Responder
}

// NewListSkus creates a new http.Handler for the list skus operation
func NewListSkus(ctx *middleware.Context, handler ListSkusHandler) *ListSkus {
	return &ListSkus{Context: ctx, Handler: handler}
}

/*ListSkus swagger:route GET /sku skuManagement listSkus

list SKUs

lists all skus

*/
type ListSkus struct {
	Context *middleware.Context
	Handler ListSkusHandler
}

func (o *ListSkus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSkusParams()

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

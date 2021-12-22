// Code generated by go-swagger; DO NOT EDIT.

package sku_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateSkuHandlerFunc turns a function with the right signature into a update sku handler
type UpdateSkuHandlerFunc func(UpdateSkuParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateSkuHandlerFunc) Handle(params UpdateSkuParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateSkuHandler interface for that can handle valid update sku params
type UpdateSkuHandler interface {
	Handle(UpdateSkuParams, interface{}) middleware.Responder
}

// NewUpdateSku creates a new http.Handler for the update sku operation
func NewUpdateSku(ctx *middleware.Context, handler UpdateSkuHandler) *UpdateSku {
	return &UpdateSku{Context: ctx, Handler: handler}
}

/*UpdateSku swagger:route PUT /sku/{id} skuManagement updateSku

Update specific sku

Update sku with given skuid

*/
type UpdateSku struct {
	Context *middleware.Context
	Handler UpdateSkuHandler
}

func (o *UpdateSku) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateSkuParams()

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
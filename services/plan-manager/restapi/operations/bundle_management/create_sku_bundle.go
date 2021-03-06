// Code generated by go-swagger; DO NOT EDIT.

package bundle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateSkuBundleHandlerFunc turns a function with the right signature into a create sku bundle handler
type CreateSkuBundleHandlerFunc func(CreateSkuBundleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateSkuBundleHandlerFunc) Handle(params CreateSkuBundleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateSkuBundleHandler interface for that can handle valid create sku bundle params
type CreateSkuBundleHandler interface {
	Handle(CreateSkuBundleParams, interface{}) middleware.Responder
}

// NewCreateSkuBundle creates a new http.Handler for the create sku bundle operation
func NewCreateSkuBundle(ctx *middleware.Context, handler CreateSkuBundleHandler) *CreateSkuBundle {
	return &CreateSkuBundle{Context: ctx, Handler: handler}
}

/*CreateSkuBundle swagger:route POST /sku/bundle bundleManagement createSkuBundle

create SKU bundle

Creates a new sku bundle

*/
type CreateSkuBundle struct {
	Context *middleware.Context
	Handler CreateSkuBundleHandler
}

func (o *CreateSkuBundle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateSkuBundleParams()

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

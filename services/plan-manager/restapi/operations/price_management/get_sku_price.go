// Code generated by go-swagger; DO NOT EDIT.

package price_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSkuPriceHandlerFunc turns a function with the right signature into a get sku price handler
type GetSkuPriceHandlerFunc func(GetSkuPriceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSkuPriceHandlerFunc) Handle(params GetSkuPriceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSkuPriceHandler interface for that can handle valid get sku price params
type GetSkuPriceHandler interface {
	Handle(GetSkuPriceParams, interface{}) middleware.Responder
}

// NewGetSkuPrice creates a new http.Handler for the get sku price operation
func NewGetSkuPrice(ctx *middleware.Context, handler GetSkuPriceHandler) *GetSkuPrice {
	return &GetSkuPrice{Context: ctx, Handler: handler}
}

/*GetSkuPrice swagger:route GET /sku/price/{id} priceManagement getSkuPrice

Get specific sku price

get sku price with given skupriceid

*/
type GetSkuPrice struct {
	Context *middleware.Context
	Handler GetSkuPriceHandler
}

func (o *GetSkuPrice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSkuPriceParams()

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
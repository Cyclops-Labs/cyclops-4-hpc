// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GenerateInvoiceForResellerHandlerFunc turns a function with the right signature into a generate invoice for reseller handler
type GenerateInvoiceForResellerHandlerFunc func(GenerateInvoiceForResellerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GenerateInvoiceForResellerHandlerFunc) Handle(params GenerateInvoiceForResellerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GenerateInvoiceForResellerHandler interface for that can handle valid generate invoice for reseller params
type GenerateInvoiceForResellerHandler interface {
	Handle(GenerateInvoiceForResellerParams, interface{}) middleware.Responder
}

// NewGenerateInvoiceForReseller creates a new http.Handler for the generate invoice for reseller operation
func NewGenerateInvoiceForReseller(ctx *middleware.Context, handler GenerateInvoiceForResellerHandler) *GenerateInvoiceForReseller {
	return &GenerateInvoiceForReseller{Context: ctx, Handler: handler}
}

/*GenerateInvoiceForReseller swagger:route POST /invoice/reseller/{id} invoiceManagement generateInvoiceForReseller

Generate invoice for the provided reseller for the provided time window or last period

*/
type GenerateInvoiceForReseller struct {
	Context *middleware.Context
	Handler GenerateInvoiceForResellerHandler
}

func (o *GenerateInvoiceForReseller) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenerateInvoiceForResellerParams()

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
// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListResellerInvoicesHandlerFunc turns a function with the right signature into a list reseller invoices handler
type ListResellerInvoicesHandlerFunc func(ListResellerInvoicesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListResellerInvoicesHandlerFunc) Handle(params ListResellerInvoicesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListResellerInvoicesHandler interface for that can handle valid list reseller invoices params
type ListResellerInvoicesHandler interface {
	Handle(ListResellerInvoicesParams, interface{}) middleware.Responder
}

// NewListResellerInvoices creates a new http.Handler for the list reseller invoices operation
func NewListResellerInvoices(ctx *middleware.Context, handler ListResellerInvoicesHandler) *ListResellerInvoices {
	return &ListResellerInvoices{Context: ctx, Handler: handler}
}

/*ListResellerInvoices swagger:route GET /invoice/reseller invoiceManagement listResellerInvoices

Retrieve resellers' invoices

*/
type ListResellerInvoices struct {
	Context *middleware.Context
	Handler ListResellerInvoicesHandler
}

func (o *ListResellerInvoices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListResellerInvoicesParams()

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

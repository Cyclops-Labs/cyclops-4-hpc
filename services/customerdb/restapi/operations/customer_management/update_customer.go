// Code generated by go-swagger; DO NOT EDIT.

package customer_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateCustomerHandlerFunc turns a function with the right signature into a update customer handler
type UpdateCustomerHandlerFunc func(UpdateCustomerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCustomerHandlerFunc) Handle(params UpdateCustomerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateCustomerHandler interface for that can handle valid update customer params
type UpdateCustomerHandler interface {
	Handle(UpdateCustomerParams, interface{}) middleware.Responder
}

// NewUpdateCustomer creates a new http.Handler for the update customer operation
func NewUpdateCustomer(ctx *middleware.Context, handler UpdateCustomerHandler) *UpdateCustomer {
	return &UpdateCustomer{Context: ctx, Handler: handler}
}

/*UpdateCustomer swagger:route PUT /customer/{id} customerManagement updateCustomer

Updates the information of the customer with the given id

*/
type UpdateCustomer struct {
	Context *middleware.Context
	Handler UpdateCustomerHandler
}

func (o *UpdateCustomer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateCustomerParams()

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

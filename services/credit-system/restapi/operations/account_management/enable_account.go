// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EnableAccountHandlerFunc turns a function with the right signature into a enable account handler
type EnableAccountHandlerFunc func(EnableAccountParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EnableAccountHandlerFunc) Handle(params EnableAccountParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EnableAccountHandler interface for that can handle valid enable account params
type EnableAccountHandler interface {
	Handle(EnableAccountParams, interface{}) middleware.Responder
}

// NewEnableAccount creates a new http.Handler for the enable account operation
func NewEnableAccount(ctx *middleware.Context, handler EnableAccountHandler) *EnableAccount {
	return &EnableAccount{Context: ctx, Handler: handler}
}

/*EnableAccount swagger:route POST /account/enable/{id} accountManagement enableAccount

Enables the account with the id provided in the system

*/
type EnableAccount struct {
	Context *middleware.Context
	Handler EnableAccountHandler
}

func (o *EnableAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEnableAccountParams()

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

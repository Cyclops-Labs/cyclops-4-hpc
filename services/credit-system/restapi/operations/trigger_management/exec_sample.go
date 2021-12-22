// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ExecSampleHandlerFunc turns a function with the right signature into a exec sample handler
type ExecSampleHandlerFunc func(ExecSampleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ExecSampleHandlerFunc) Handle(params ExecSampleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ExecSampleHandler interface for that can handle valid exec sample params
type ExecSampleHandler interface {
	Handle(ExecSampleParams, interface{}) middleware.Responder
}

// NewExecSample creates a new http.Handler for the exec sample operation
func NewExecSample(ctx *middleware.Context, handler ExecSampleHandler) *ExecSample {
	return &ExecSample{Context: ctx, Handler: handler}
}

/*ExecSample swagger:route GET /trigger/sample triggerManagement execSample

Sample task trigger

*/
type ExecSample struct {
	Context *middleware.Context
	Handler ExecSampleHandler
}

func (o *ExecSample) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewExecSampleParams()

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
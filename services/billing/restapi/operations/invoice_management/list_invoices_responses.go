// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/models"
)

// ListInvoicesOKCode is the HTTP code returned for type ListInvoicesOK
const ListInvoicesOKCode int = 200

/*ListInvoicesOK Description of a successfully operation

swagger:response listInvoicesOK
*/
type ListInvoicesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Invoice `json:"body,omitempty"`
}

// NewListInvoicesOK creates ListInvoicesOK with default headers values
func NewListInvoicesOK() *ListInvoicesOK {

	return &ListInvoicesOK{}
}

// WithPayload adds the payload to the list invoices o k response
func (o *ListInvoicesOK) WithPayload(payload []*models.Invoice) *ListInvoicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list invoices o k response
func (o *ListInvoicesOK) SetPayload(payload []*models.Invoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInvoicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Invoice, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListInvoicesInternalServerErrorCode is the HTTP code returned for type ListInvoicesInternalServerError
const ListInvoicesInternalServerErrorCode int = 500

/*ListInvoicesInternalServerError Something unexpected happend, error raised

swagger:response listInvoicesInternalServerError
*/
type ListInvoicesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListInvoicesInternalServerError creates ListInvoicesInternalServerError with default headers values
func NewListInvoicesInternalServerError() *ListInvoicesInternalServerError {

	return &ListInvoicesInternalServerError{}
}

// WithPayload adds the payload to the list invoices internal server error response
func (o *ListInvoicesInternalServerError) WithPayload(payload *models.ErrorResponse) *ListInvoicesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list invoices internal server error response
func (o *ListInvoicesInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInvoicesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
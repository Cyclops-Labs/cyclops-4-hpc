// Code generated by go-swagger; DO NOT EDIT.

package customer_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
)

// ListCustomersOKCode is the HTTP code returned for type ListCustomersOK
const ListCustomersOKCode int = 200

/*ListCustomersOK List of customers in the system returned

swagger:response listCustomersOK
*/
type ListCustomersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Customer `json:"body,omitempty"`
}

// NewListCustomersOK creates ListCustomersOK with default headers values
func NewListCustomersOK() *ListCustomersOK {

	return &ListCustomersOK{}
}

// WithPayload adds the payload to the list customers o k response
func (o *ListCustomersOK) WithPayload(payload []*models.Customer) *ListCustomersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list customers o k response
func (o *ListCustomersOK) SetPayload(payload []*models.Customer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCustomersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Customer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListCustomersInternalServerErrorCode is the HTTP code returned for type ListCustomersInternalServerError
const ListCustomersInternalServerErrorCode int = 500

/*ListCustomersInternalServerError Something unexpected happend, error raised

swagger:response listCustomersInternalServerError
*/
type ListCustomersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListCustomersInternalServerError creates ListCustomersInternalServerError with default headers values
func NewListCustomersInternalServerError() *ListCustomersInternalServerError {

	return &ListCustomersInternalServerError{}
}

// WithPayload adds the payload to the list customers internal server error response
func (o *ListCustomersInternalServerError) WithPayload(payload *models.ErrorResponse) *ListCustomersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list customers internal server error response
func (o *ListCustomersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCustomersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/models"
)

// ListAccountsOKCode is the HTTP code returned for type ListAccountsOK
const ListAccountsOKCode int = 200

/*ListAccountsOK List of accounts in the system

swagger:response listAccountsOK
*/
type ListAccountsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AccountStatus `json:"body,omitempty"`
}

// NewListAccountsOK creates ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {

	return &ListAccountsOK{}
}

// WithPayload adds the payload to the list accounts o k response
func (o *ListAccountsOK) WithPayload(payload []*models.AccountStatus) *ListAccountsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list accounts o k response
func (o *ListAccountsOK) SetPayload(payload []*models.AccountStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAccountsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AccountStatus, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAccountsInternalServerErrorCode is the HTTP code returned for type ListAccountsInternalServerError
const ListAccountsInternalServerErrorCode int = 500

/*ListAccountsInternalServerError Something unexpected happend, error raised

swagger:response listAccountsInternalServerError
*/
type ListAccountsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListAccountsInternalServerError creates ListAccountsInternalServerError with default headers values
func NewListAccountsInternalServerError() *ListAccountsInternalServerError {

	return &ListAccountsInternalServerError{}
}

// WithPayload adds the payload to the list accounts internal server error response
func (o *ListAccountsInternalServerError) WithPayload(payload *models.ErrorResponse) *ListAccountsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list accounts internal server error response
func (o *ListAccountsInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAccountsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

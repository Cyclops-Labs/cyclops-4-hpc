// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/models"
)

// GetAccountStatusOKCode is the HTTP code returned for type GetAccountStatusOK
const GetAccountStatusOKCode int = 200

/*GetAccountStatusOK Status information of the account with provided id in the system

swagger:response getAccountStatusOK
*/
type GetAccountStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccountStatus `json:"body,omitempty"`
}

// NewGetAccountStatusOK creates GetAccountStatusOK with default headers values
func NewGetAccountStatusOK() *GetAccountStatusOK {

	return &GetAccountStatusOK{}
}

// WithPayload adds the payload to the get account status o k response
func (o *GetAccountStatusOK) WithPayload(payload *models.AccountStatus) *GetAccountStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account status o k response
func (o *GetAccountStatusOK) SetPayload(payload *models.AccountStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountStatusNotFoundCode is the HTTP code returned for type GetAccountStatusNotFound
const GetAccountStatusNotFoundCode int = 404

/*GetAccountStatusNotFound The account with the id provided doesn't exist

swagger:response getAccountStatusNotFound
*/
type GetAccountStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetAccountStatusNotFound creates GetAccountStatusNotFound with default headers values
func NewGetAccountStatusNotFound() *GetAccountStatusNotFound {

	return &GetAccountStatusNotFound{}
}

// WithPayload adds the payload to the get account status not found response
func (o *GetAccountStatusNotFound) WithPayload(payload *models.ErrorResponse) *GetAccountStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account status not found response
func (o *GetAccountStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAccountStatusInternalServerErrorCode is the HTTP code returned for type GetAccountStatusInternalServerError
const GetAccountStatusInternalServerErrorCode int = 500

/*GetAccountStatusInternalServerError Something unexpected happend, error raised

swagger:response getAccountStatusInternalServerError
*/
type GetAccountStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetAccountStatusInternalServerError creates GetAccountStatusInternalServerError with default headers values
func NewGetAccountStatusInternalServerError() *GetAccountStatusInternalServerError {

	return &GetAccountStatusInternalServerError{}
}

// WithPayload adds the payload to the get account status internal server error response
func (o *GetAccountStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *GetAccountStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get account status internal server error response
func (o *GetAccountStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAccountStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

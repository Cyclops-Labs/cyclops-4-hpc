// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/models"
)

// GetUsageOKCode is the HTTP code returned for type GetUsageOK
const GetUsageOKCode int = 200

/*GetUsageOK Description of a successfully operation

swagger:response getUsageOK
*/
type GetUsageOK struct {

	/*
	  In: Body
	*/
	Payload *models.Usage `json:"body,omitempty"`
}

// NewGetUsageOK creates GetUsageOK with default headers values
func NewGetUsageOK() *GetUsageOK {

	return &GetUsageOK{}
}

// WithPayload adds the payload to the get usage o k response
func (o *GetUsageOK) WithPayload(payload *models.Usage) *GetUsageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get usage o k response
func (o *GetUsageOK) SetPayload(payload *models.Usage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsageNotFoundCode is the HTTP code returned for type GetUsageNotFound
const GetUsageNotFoundCode int = 404

/*GetUsageNotFound Item not found in the system

swagger:response getUsageNotFound
*/
type GetUsageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUsageNotFound creates GetUsageNotFound with default headers values
func NewGetUsageNotFound() *GetUsageNotFound {

	return &GetUsageNotFound{}
}

// WithPayload adds the payload to the get usage not found response
func (o *GetUsageNotFound) WithPayload(payload *models.ErrorResponse) *GetUsageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get usage not found response
func (o *GetUsageNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsageInternalServerErrorCode is the HTTP code returned for type GetUsageInternalServerError
const GetUsageInternalServerErrorCode int = 500

/*GetUsageInternalServerError Something unexpected happend, error raised

swagger:response getUsageInternalServerError
*/
type GetUsageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUsageInternalServerError creates GetUsageInternalServerError with default headers values
func NewGetUsageInternalServerError() *GetUsageInternalServerError {

	return &GetUsageInternalServerError{}
}

// WithPayload adds the payload to the get usage internal server error response
func (o *GetUsageInternalServerError) WithPayload(payload *models.ErrorResponse) *GetUsageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get usage internal server error response
func (o *GetUsageInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
// Code generated by go-swagger; DO NOT EDIT.

package status_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
)

// GetStatusOKCode is the HTTP code returned for type GetStatusOK
const GetStatusOKCode int = 200

/*GetStatusOK Status information of the system

swagger:response getStatusOK
*/
type GetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewGetStatusOK creates GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {

	return &GetStatusOK{}
}

// WithPayload adds the payload to the get status o k response
func (o *GetStatusOK) WithPayload(payload *models.Status) *GetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status o k response
func (o *GetStatusOK) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStatusNotFoundCode is the HTTP code returned for type GetStatusNotFound
const GetStatusNotFoundCode int = 404

/*GetStatusNotFound The endpoint provided doesn't exist

swagger:response getStatusNotFound
*/
type GetStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetStatusNotFound creates GetStatusNotFound with default headers values
func NewGetStatusNotFound() *GetStatusNotFound {

	return &GetStatusNotFound{}
}

// WithPayload adds the payload to the get status not found response
func (o *GetStatusNotFound) WithPayload(payload *models.ErrorResponse) *GetStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get status not found response
func (o *GetStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

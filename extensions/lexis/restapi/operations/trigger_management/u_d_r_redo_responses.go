// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
)

// UDRRedoOKCode is the HTTP code returned for type UDRRedoOK
const UDRRedoOKCode int = 200

/*UDRRedoOK Generation task executed successfully.

swagger:response uDRRedoOK
*/
type UDRRedoOK struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewUDRRedoOK creates UDRRedoOK with default headers values
func NewUDRRedoOK() *UDRRedoOK {

	return &UDRRedoOK{}
}

// WithPayload adds the payload to the u d r redo o k response
func (o *UDRRedoOK) WithPayload(payload *models.ItemCreatedResponse) *UDRRedoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the u d r redo o k response
func (o *UDRRedoOK) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UDRRedoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UDRRedoInternalServerErrorCode is the HTTP code returned for type UDRRedoInternalServerError
const UDRRedoInternalServerErrorCode int = 500

/*UDRRedoInternalServerError Something unexpected happend, error raised

swagger:response uDRRedoInternalServerError
*/
type UDRRedoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUDRRedoInternalServerError creates UDRRedoInternalServerError with default headers values
func NewUDRRedoInternalServerError() *UDRRedoInternalServerError {

	return &UDRRedoInternalServerError{}
}

// WithPayload adds the payload to the u d r redo internal server error response
func (o *UDRRedoInternalServerError) WithPayload(payload *models.ErrorResponse) *UDRRedoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the u d r redo internal server error response
func (o *UDRRedoInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UDRRedoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package reseller_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
)

// ListResellersOKCode is the HTTP code returned for type ListResellersOK
const ListResellersOKCode int = 200

/*ListResellersOK List of resellers in the system returned

swagger:response listResellersOK
*/
type ListResellersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Reseller `json:"body,omitempty"`
}

// NewListResellersOK creates ListResellersOK with default headers values
func NewListResellersOK() *ListResellersOK {

	return &ListResellersOK{}
}

// WithPayload adds the payload to the list resellers o k response
func (o *ListResellersOK) WithPayload(payload []*models.Reseller) *ListResellersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list resellers o k response
func (o *ListResellersOK) SetPayload(payload []*models.Reseller) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListResellersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Reseller, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListResellersInternalServerErrorCode is the HTTP code returned for type ListResellersInternalServerError
const ListResellersInternalServerErrorCode int = 500

/*ListResellersInternalServerError Something unexpected happend, error raised

swagger:response listResellersInternalServerError
*/
type ListResellersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListResellersInternalServerError creates ListResellersInternalServerError with default headers values
func NewListResellersInternalServerError() *ListResellersInternalServerError {

	return &ListResellersInternalServerError{}
}

// WithPayload adds the payload to the list resellers internal server error response
func (o *ListResellersInternalServerError) WithPayload(payload *models.ErrorResponse) *ListResellersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list resellers internal server error response
func (o *ListResellersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListResellersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

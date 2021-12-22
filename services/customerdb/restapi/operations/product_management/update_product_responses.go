// Code generated by go-swagger; DO NOT EDIT.

package product_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
)

// UpdateProductOKCode is the HTTP code returned for type UpdateProductOK
const UpdateProductOKCode int = 200

/*UpdateProductOK Product with the given id was updated

swagger:response updateProductOK
*/
type UpdateProductOK struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewUpdateProductOK creates UpdateProductOK with default headers values
func NewUpdateProductOK() *UpdateProductOK {

	return &UpdateProductOK{}
}

// WithPayload adds the payload to the update product o k response
func (o *UpdateProductOK) WithPayload(payload *models.ItemCreatedResponse) *UpdateProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product o k response
func (o *UpdateProductOK) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateProductNotFoundCode is the HTTP code returned for type UpdateProductNotFound
const UpdateProductNotFoundCode int = 404

/*UpdateProductNotFound The product with the given id wasn't found

swagger:response updateProductNotFound
*/
type UpdateProductNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateProductNotFound creates UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {

	return &UpdateProductNotFound{}
}

// WithPayload adds the payload to the update product not found response
func (o *UpdateProductNotFound) WithPayload(payload *models.ErrorResponse) *UpdateProductNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product not found response
func (o *UpdateProductNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateProductInternalServerErrorCode is the HTTP code returned for type UpdateProductInternalServerError
const UpdateProductInternalServerErrorCode int = 500

/*UpdateProductInternalServerError Something unexpected happend, error raised

swagger:response updateProductInternalServerError
*/
type UpdateProductInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateProductInternalServerError creates UpdateProductInternalServerError with default headers values
func NewUpdateProductInternalServerError() *UpdateProductInternalServerError {

	return &UpdateProductInternalServerError{}
}

// WithPayload adds the payload to the update product internal server error response
func (o *UpdateProductInternalServerError) WithPayload(payload *models.ErrorResponse) *UpdateProductInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product internal server error response
func (o *UpdateProductInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

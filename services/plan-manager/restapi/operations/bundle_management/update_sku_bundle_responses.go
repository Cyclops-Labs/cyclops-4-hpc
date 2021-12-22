// Code generated by go-swagger; DO NOT EDIT.

package bundle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// UpdateSkuBundleOKCode is the HTTP code returned for type UpdateSkuBundleOK
const UpdateSkuBundleOKCode int = 200

/*UpdateSkuBundleOK updated sku bundle

swagger:response updateSkuBundleOK
*/
type UpdateSkuBundleOK struct {

	/*
	  In: Body
	*/
	Payload *models.SkuBundle `json:"body,omitempty"`
}

// NewUpdateSkuBundleOK creates UpdateSkuBundleOK with default headers values
func NewUpdateSkuBundleOK() *UpdateSkuBundleOK {

	return &UpdateSkuBundleOK{}
}

// WithPayload adds the payload to the update sku bundle o k response
func (o *UpdateSkuBundleOK) WithPayload(payload *models.SkuBundle) *UpdateSkuBundleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update sku bundle o k response
func (o *UpdateSkuBundleOK) SetPayload(payload *models.SkuBundle) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSkuBundleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSkuBundleNotFoundCode is the HTTP code returned for type UpdateSkuBundleNotFound
const UpdateSkuBundleNotFoundCode int = 404

/*UpdateSkuBundleNotFound sku bundle with id not found

swagger:response updateSkuBundleNotFound
*/
type UpdateSkuBundleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateSkuBundleNotFound creates UpdateSkuBundleNotFound with default headers values
func NewUpdateSkuBundleNotFound() *UpdateSkuBundleNotFound {

	return &UpdateSkuBundleNotFound{}
}

// WithPayload adds the payload to the update sku bundle not found response
func (o *UpdateSkuBundleNotFound) WithPayload(payload *models.ErrorResponse) *UpdateSkuBundleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update sku bundle not found response
func (o *UpdateSkuBundleNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSkuBundleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSkuBundleInternalServerErrorCode is the HTTP code returned for type UpdateSkuBundleInternalServerError
const UpdateSkuBundleInternalServerErrorCode int = 500

/*UpdateSkuBundleInternalServerError unexpected error

swagger:response updateSkuBundleInternalServerError
*/
type UpdateSkuBundleInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateSkuBundleInternalServerError creates UpdateSkuBundleInternalServerError with default headers values
func NewUpdateSkuBundleInternalServerError() *UpdateSkuBundleInternalServerError {

	return &UpdateSkuBundleInternalServerError{}
}

// WithPayload adds the payload to the update sku bundle internal server error response
func (o *UpdateSkuBundleInternalServerError) WithPayload(payload *models.ErrorResponse) *UpdateSkuBundleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update sku bundle internal server error response
func (o *UpdateSkuBundleInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSkuBundleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

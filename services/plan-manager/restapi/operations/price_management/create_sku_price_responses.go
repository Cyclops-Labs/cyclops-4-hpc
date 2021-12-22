// Code generated by go-swagger; DO NOT EDIT.

package price_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// CreateSkuPriceCreatedCode is the HTTP code returned for type CreateSkuPriceCreated
const CreateSkuPriceCreatedCode int = 201

/*CreateSkuPriceCreated item created

swagger:response createSkuPriceCreated
*/
type CreateSkuPriceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewCreateSkuPriceCreated creates CreateSkuPriceCreated with default headers values
func NewCreateSkuPriceCreated() *CreateSkuPriceCreated {

	return &CreateSkuPriceCreated{}
}

// WithPayload adds the payload to the create sku price created response
func (o *CreateSkuPriceCreated) WithPayload(payload *models.ItemCreatedResponse) *CreateSkuPriceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create sku price created response
func (o *CreateSkuPriceCreated) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSkuPriceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSkuPriceBadRequestCode is the HTTP code returned for type CreateSkuPriceBadRequest
const CreateSkuPriceBadRequestCode int = 400

/*CreateSkuPriceBadRequest invalid input, object invalid

swagger:response createSkuPriceBadRequest
*/
type CreateSkuPriceBadRequest struct {
}

// NewCreateSkuPriceBadRequest creates CreateSkuPriceBadRequest with default headers values
func NewCreateSkuPriceBadRequest() *CreateSkuPriceBadRequest {

	return &CreateSkuPriceBadRequest{}
}

// WriteResponse to the client
func (o *CreateSkuPriceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateSkuPriceConflictCode is the HTTP code returned for type CreateSkuPriceConflict
const CreateSkuPriceConflictCode int = 409

/*CreateSkuPriceConflict an existing item already exists

swagger:response createSkuPriceConflict
*/
type CreateSkuPriceConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateSkuPriceConflict creates CreateSkuPriceConflict with default headers values
func NewCreateSkuPriceConflict() *CreateSkuPriceConflict {

	return &CreateSkuPriceConflict{}
}

// WithPayload adds the payload to the create sku price conflict response
func (o *CreateSkuPriceConflict) WithPayload(payload *models.ErrorResponse) *CreateSkuPriceConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create sku price conflict response
func (o *CreateSkuPriceConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSkuPriceConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSkuPriceInternalServerErrorCode is the HTTP code returned for type CreateSkuPriceInternalServerError
const CreateSkuPriceInternalServerErrorCode int = 500

/*CreateSkuPriceInternalServerError unexpected error

swagger:response createSkuPriceInternalServerError
*/
type CreateSkuPriceInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCreateSkuPriceInternalServerError creates CreateSkuPriceInternalServerError with default headers values
func NewCreateSkuPriceInternalServerError() *CreateSkuPriceInternalServerError {

	return &CreateSkuPriceInternalServerError{}
}

// WithPayload adds the payload to the create sku price internal server error response
func (o *CreateSkuPriceInternalServerError) WithPayload(payload *models.ErrorResponse) *CreateSkuPriceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create sku price internal server error response
func (o *CreateSkuPriceInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSkuPriceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
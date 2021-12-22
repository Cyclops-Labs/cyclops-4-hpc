// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/models"
)

// GenerateInvoiceForResellerAcceptedCode is the HTTP code returned for type GenerateInvoiceForResellerAccepted
const GenerateInvoiceForResellerAcceptedCode int = 202

/*GenerateInvoiceForResellerAccepted The request for processing had been added to the queue

swagger:response generateInvoiceForResellerAccepted
*/
type GenerateInvoiceForResellerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewGenerateInvoiceForResellerAccepted creates GenerateInvoiceForResellerAccepted with default headers values
func NewGenerateInvoiceForResellerAccepted() *GenerateInvoiceForResellerAccepted {

	return &GenerateInvoiceForResellerAccepted{}
}

// WithPayload adds the payload to the generate invoice for reseller accepted response
func (o *GenerateInvoiceForResellerAccepted) WithPayload(payload *models.ItemCreatedResponse) *GenerateInvoiceForResellerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate invoice for reseller accepted response
func (o *GenerateInvoiceForResellerAccepted) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateInvoiceForResellerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateInvoiceForResellerBadRequestCode is the HTTP code returned for type GenerateInvoiceForResellerBadRequest
const GenerateInvoiceForResellerBadRequestCode int = 400

/*GenerateInvoiceForResellerBadRequest Invalid input, object invalid

swagger:response generateInvoiceForResellerBadRequest
*/
type GenerateInvoiceForResellerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGenerateInvoiceForResellerBadRequest creates GenerateInvoiceForResellerBadRequest with default headers values
func NewGenerateInvoiceForResellerBadRequest() *GenerateInvoiceForResellerBadRequest {

	return &GenerateInvoiceForResellerBadRequest{}
}

// WithPayload adds the payload to the generate invoice for reseller bad request response
func (o *GenerateInvoiceForResellerBadRequest) WithPayload(payload *models.ErrorResponse) *GenerateInvoiceForResellerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate invoice for reseller bad request response
func (o *GenerateInvoiceForResellerBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateInvoiceForResellerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GenerateInvoiceForResellerInternalServerErrorCode is the HTTP code returned for type GenerateInvoiceForResellerInternalServerError
const GenerateInvoiceForResellerInternalServerErrorCode int = 500

/*GenerateInvoiceForResellerInternalServerError Something unexpected happend, error raised

swagger:response generateInvoiceForResellerInternalServerError
*/
type GenerateInvoiceForResellerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGenerateInvoiceForResellerInternalServerError creates GenerateInvoiceForResellerInternalServerError with default headers values
func NewGenerateInvoiceForResellerInternalServerError() *GenerateInvoiceForResellerInternalServerError {

	return &GenerateInvoiceForResellerInternalServerError{}
}

// WithPayload adds the payload to the generate invoice for reseller internal server error response
func (o *GenerateInvoiceForResellerInternalServerError) WithPayload(payload *models.ErrorResponse) *GenerateInvoiceForResellerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the generate invoice for reseller internal server error response
func (o *GenerateInvoiceForResellerInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GenerateInvoiceForResellerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

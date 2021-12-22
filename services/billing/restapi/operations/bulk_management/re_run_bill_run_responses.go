// Code generated by go-swagger; DO NOT EDIT.

package bulk_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/models"
)

// ReRunBillRunAcceptedCode is the HTTP code returned for type ReRunBillRunAccepted
const ReRunBillRunAcceptedCode int = 202

/*ReRunBillRunAccepted The request for processing had been added to the queue

swagger:response reRunBillRunAccepted
*/
type ReRunBillRunAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewReRunBillRunAccepted creates ReRunBillRunAccepted with default headers values
func NewReRunBillRunAccepted() *ReRunBillRunAccepted {

	return &ReRunBillRunAccepted{}
}

// WithPayload adds the payload to the re run bill run accepted response
func (o *ReRunBillRunAccepted) WithPayload(payload *models.ItemCreatedResponse) *ReRunBillRunAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the re run bill run accepted response
func (o *ReRunBillRunAccepted) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReRunBillRunAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReRunBillRunNotFoundCode is the HTTP code returned for type ReRunBillRunNotFound
const ReRunBillRunNotFoundCode int = 404

/*ReRunBillRunNotFound The invoice id provided doesn't exist

swagger:response reRunBillRunNotFound
*/
type ReRunBillRunNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewReRunBillRunNotFound creates ReRunBillRunNotFound with default headers values
func NewReRunBillRunNotFound() *ReRunBillRunNotFound {

	return &ReRunBillRunNotFound{}
}

// WithPayload adds the payload to the re run bill run not found response
func (o *ReRunBillRunNotFound) WithPayload(payload *models.ErrorResponse) *ReRunBillRunNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the re run bill run not found response
func (o *ReRunBillRunNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReRunBillRunNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReRunBillRunInternalServerErrorCode is the HTTP code returned for type ReRunBillRunInternalServerError
const ReRunBillRunInternalServerErrorCode int = 500

/*ReRunBillRunInternalServerError Something unexpected happend, error raised

swagger:response reRunBillRunInternalServerError
*/
type ReRunBillRunInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewReRunBillRunInternalServerError creates ReRunBillRunInternalServerError with default headers values
func NewReRunBillRunInternalServerError() *ReRunBillRunInternalServerError {

	return &ReRunBillRunInternalServerError{}
}

// WithPayload adds the payload to the re run bill run internal server error response
func (o *ReRunBillRunInternalServerError) WithPayload(payload *models.ErrorResponse) *ReRunBillRunInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the re run bill run internal server error response
func (o *ReRunBillRunInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReRunBillRunInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/models"
)

// PeriodicRunAcceptedCode is the HTTP code returned for type PeriodicRunAccepted
const PeriodicRunAcceptedCode int = 202

/*PeriodicRunAccepted The request for processing the periodic run had been added to the queue

swagger:response periodicRunAccepted
*/
type PeriodicRunAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ItemCreatedResponse `json:"body,omitempty"`
}

// NewPeriodicRunAccepted creates PeriodicRunAccepted with default headers values
func NewPeriodicRunAccepted() *PeriodicRunAccepted {

	return &PeriodicRunAccepted{}
}

// WithPayload adds the payload to the periodic run accepted response
func (o *PeriodicRunAccepted) WithPayload(payload *models.ItemCreatedResponse) *PeriodicRunAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the periodic run accepted response
func (o *PeriodicRunAccepted) SetPayload(payload *models.ItemCreatedResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PeriodicRunAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
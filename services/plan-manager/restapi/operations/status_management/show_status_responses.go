// Code generated by go-swagger; DO NOT EDIT.

package status_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// ShowStatusOKCode is the HTTP code returned for type ShowStatusOK
const ShowStatusOKCode int = 200

/*ShowStatusOK Status information of the system

swagger:response showStatusOK
*/
type ShowStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.Status `json:"body,omitempty"`
}

// NewShowStatusOK creates ShowStatusOK with default headers values
func NewShowStatusOK() *ShowStatusOK {

	return &ShowStatusOK{}
}

// WithPayload adds the payload to the show status o k response
func (o *ShowStatusOK) WithPayload(payload *models.Status) *ShowStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show status o k response
func (o *ShowStatusOK) SetPayload(payload *models.Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

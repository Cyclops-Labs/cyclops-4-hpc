// Code generated by go-swagger; DO NOT EDIT.

package cycle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// CreateCycleReader is a Reader for the CreateCycle structure.
type CreateCycleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCycleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCycleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCycleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCycleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCycleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCycleCreated creates a CreateCycleCreated with default headers values
func NewCreateCycleCreated() *CreateCycleCreated {
	return &CreateCycleCreated{}
}

/*CreateCycleCreated handles this case with default header values.

item created
*/
type CreateCycleCreated struct {
	Payload *models.ItemCreatedResponse
}

func (o *CreateCycleCreated) Error() string {
	return fmt.Sprintf("[POST /cycle][%d] createCycleCreated  %+v", 201, o.Payload)
}

func (o *CreateCycleCreated) GetPayload() *models.ItemCreatedResponse {
	return o.Payload
}

func (o *CreateCycleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ItemCreatedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCycleBadRequest creates a CreateCycleBadRequest with default headers values
func NewCreateCycleBadRequest() *CreateCycleBadRequest {
	return &CreateCycleBadRequest{}
}

/*CreateCycleBadRequest handles this case with default header values.

invalid input, object invalid
*/
type CreateCycleBadRequest struct {
}

func (o *CreateCycleBadRequest) Error() string {
	return fmt.Sprintf("[POST /cycle][%d] createCycleBadRequest ", 400)
}

func (o *CreateCycleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCycleConflict creates a CreateCycleConflict with default headers values
func NewCreateCycleConflict() *CreateCycleConflict {
	return &CreateCycleConflict{}
}

/*CreateCycleConflict handles this case with default header values.

an existing item already exists
*/
type CreateCycleConflict struct {
	Payload *models.ErrorResponse
}

func (o *CreateCycleConflict) Error() string {
	return fmt.Sprintf("[POST /cycle][%d] createCycleConflict  %+v", 409, o.Payload)
}

func (o *CreateCycleConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateCycleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCycleInternalServerError creates a CreateCycleInternalServerError with default headers values
func NewCreateCycleInternalServerError() *CreateCycleInternalServerError {
	return &CreateCycleInternalServerError{}
}

/*CreateCycleInternalServerError handles this case with default header values.

unexpected error
*/
type CreateCycleInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateCycleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cycle][%d] createCycleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCycleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateCycleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

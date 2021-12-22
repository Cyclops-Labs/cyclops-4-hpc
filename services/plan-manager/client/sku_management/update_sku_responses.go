// Code generated by go-swagger; DO NOT EDIT.

package sku_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// UpdateSkuReader is a Reader for the UpdateSku structure.
type UpdateSkuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSkuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSkuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSkuNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSkuInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSkuOK creates a UpdateSkuOK with default headers values
func NewUpdateSkuOK() *UpdateSkuOK {
	return &UpdateSkuOK{}
}

/*UpdateSkuOK handles this case with default header values.

updated sku
*/
type UpdateSkuOK struct {
	Payload *models.Sku
}

func (o *UpdateSkuOK) Error() string {
	return fmt.Sprintf("[PUT /sku/{id}][%d] updateSkuOK  %+v", 200, o.Payload)
}

func (o *UpdateSkuOK) GetPayload() *models.Sku {
	return o.Payload
}

func (o *UpdateSkuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sku)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSkuNotFound creates a UpdateSkuNotFound with default headers values
func NewUpdateSkuNotFound() *UpdateSkuNotFound {
	return &UpdateSkuNotFound{}
}

/*UpdateSkuNotFound handles this case with default header values.

sku with skuid not found
*/
type UpdateSkuNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateSkuNotFound) Error() string {
	return fmt.Sprintf("[PUT /sku/{id}][%d] updateSkuNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSkuNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSkuNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSkuInternalServerError creates a UpdateSkuInternalServerError with default headers values
func NewUpdateSkuInternalServerError() *UpdateSkuInternalServerError {
	return &UpdateSkuInternalServerError{}
}

/*UpdateSkuInternalServerError handles this case with default header values.

unexpected error
*/
type UpdateSkuInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *UpdateSkuInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /sku/{id}][%d] updateSkuInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSkuInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSkuInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

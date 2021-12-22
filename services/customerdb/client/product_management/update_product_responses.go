// Code generated by go-swagger; DO NOT EDIT.

package product_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProductOK creates a UpdateProductOK with default headers values
func NewUpdateProductOK() *UpdateProductOK {
	return &UpdateProductOK{}
}

/*UpdateProductOK handles this case with default header values.

Product with the given id was updated
*/
type UpdateProductOK struct {
	Payload *models.ItemCreatedResponse
}

func (o *UpdateProductOK) Error() string {
	return fmt.Sprintf("[PUT /product/{id}][%d] updateProductOK  %+v", 200, o.Payload)
}

func (o *UpdateProductOK) GetPayload() *models.ItemCreatedResponse {
	return o.Payload
}

func (o *UpdateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ItemCreatedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/*UpdateProductNotFound handles this case with default header values.

The product with the given id wasn't found
*/
type UpdateProductNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PUT /product/{id}][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductInternalServerError creates a UpdateProductInternalServerError with default headers values
func NewUpdateProductInternalServerError() *UpdateProductInternalServerError {
	return &UpdateProductInternalServerError{}
}

/*UpdateProductInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type UpdateProductInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *UpdateProductInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /product/{id}][%d] updateProductInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProductInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

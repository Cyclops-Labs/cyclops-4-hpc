// Code generated by go-swagger; DO NOT EDIT.

package bundle_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// UpdateSkuBundleReader is a Reader for the UpdateSkuBundle structure.
type UpdateSkuBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSkuBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSkuBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSkuBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSkuBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSkuBundleOK creates a UpdateSkuBundleOK with default headers values
func NewUpdateSkuBundleOK() *UpdateSkuBundleOK {
	return &UpdateSkuBundleOK{}
}

/*UpdateSkuBundleOK handles this case with default header values.

updated sku bundle
*/
type UpdateSkuBundleOK struct {
	Payload *models.SkuBundle
}

func (o *UpdateSkuBundleOK) Error() string {
	return fmt.Sprintf("[PUT /sku/bundle/{id}][%d] updateSkuBundleOK  %+v", 200, o.Payload)
}

func (o *UpdateSkuBundleOK) GetPayload() *models.SkuBundle {
	return o.Payload
}

func (o *UpdateSkuBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkuBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSkuBundleNotFound creates a UpdateSkuBundleNotFound with default headers values
func NewUpdateSkuBundleNotFound() *UpdateSkuBundleNotFound {
	return &UpdateSkuBundleNotFound{}
}

/*UpdateSkuBundleNotFound handles this case with default header values.

sku bundle with id not found
*/
type UpdateSkuBundleNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateSkuBundleNotFound) Error() string {
	return fmt.Sprintf("[PUT /sku/bundle/{id}][%d] updateSkuBundleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSkuBundleNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSkuBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSkuBundleInternalServerError creates a UpdateSkuBundleInternalServerError with default headers values
func NewUpdateSkuBundleInternalServerError() *UpdateSkuBundleInternalServerError {
	return &UpdateSkuBundleInternalServerError{}
}

/*UpdateSkuBundleInternalServerError handles this case with default header values.

unexpected error
*/
type UpdateSkuBundleInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *UpdateSkuBundleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /sku/bundle/{id}][%d] updateSkuBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSkuBundleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateSkuBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

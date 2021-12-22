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

// GetSkuBundleReader is a Reader for the GetSkuBundle structure.
type GetSkuBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSkuBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSkuBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSkuBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSkuBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSkuBundleOK creates a GetSkuBundleOK with default headers values
func NewGetSkuBundleOK() *GetSkuBundleOK {
	return &GetSkuBundleOK{}
}

/*GetSkuBundleOK handles this case with default header values.

sku bundle returned
*/
type GetSkuBundleOK struct {
	Payload *models.SkuBundle
}

func (o *GetSkuBundleOK) Error() string {
	return fmt.Sprintf("[GET /sku/bundle/{id}][%d] getSkuBundleOK  %+v", 200, o.Payload)
}

func (o *GetSkuBundleOK) GetPayload() *models.SkuBundle {
	return o.Payload
}

func (o *GetSkuBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkuBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSkuBundleNotFound creates a GetSkuBundleNotFound with default headers values
func NewGetSkuBundleNotFound() *GetSkuBundleNotFound {
	return &GetSkuBundleNotFound{}
}

/*GetSkuBundleNotFound handles this case with default header values.

sku bundle with id not found
*/
type GetSkuBundleNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetSkuBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /sku/bundle/{id}][%d] getSkuBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetSkuBundleNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSkuBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSkuBundleInternalServerError creates a GetSkuBundleInternalServerError with default headers values
func NewGetSkuBundleInternalServerError() *GetSkuBundleInternalServerError {
	return &GetSkuBundleInternalServerError{}
}

/*GetSkuBundleInternalServerError handles this case with default header values.

unexpected error
*/
type GetSkuBundleInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetSkuBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sku/bundle/{id}][%d] getSkuBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSkuBundleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSkuBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
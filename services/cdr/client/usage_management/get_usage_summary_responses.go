// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/cdr/models"
)

// GetUsageSummaryReader is a Reader for the GetUsageSummary structure.
type GetUsageSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsageSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsageSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUsageSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsageSummaryOK creates a GetUsageSummaryOK with default headers values
func NewGetUsageSummaryOK() *GetUsageSummaryOK {
	return &GetUsageSummaryOK{}
}

/*GetUsageSummaryOK handles this case with default header values.

Description of a successfully operation
*/
type GetUsageSummaryOK struct {
	Payload *models.UISummary
}

func (o *GetUsageSummaryOK) Error() string {
	return fmt.Sprintf("[GET /usage/summary/{id}][%d] getUsageSummaryOK  %+v", 200, o.Payload)
}

func (o *GetUsageSummaryOK) GetPayload() *models.UISummary {
	return o.Payload
}

func (o *GetUsageSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UISummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageSummaryInternalServerError creates a GetUsageSummaryInternalServerError with default headers values
func NewGetUsageSummaryInternalServerError() *GetUsageSummaryInternalServerError {
	return &GetUsageSummaryInternalServerError{}
}

/*GetUsageSummaryInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type GetUsageSummaryInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetUsageSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usage/summary/{id}][%d] getUsageSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsageSummaryInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsageSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
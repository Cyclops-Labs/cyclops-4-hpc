// Code generated by go-swagger; DO NOT EDIT.

package credit_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/models"
)

// GetCreditReader is a Reader for the GetCredit structure.
type GetCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCreditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCreditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCreditOK creates a GetCreditOK with default headers values
func NewGetCreditOK() *GetCreditOK {
	return &GetCreditOK{}
}

/*GetCreditOK handles this case with default header values.

Credit status of the account with the provided id
*/
type GetCreditOK struct {
	Payload *models.CreditStatus
}

func (o *GetCreditOK) Error() string {
	return fmt.Sprintf("[GET /account/available/{id}][%d] getCreditOK  %+v", 200, o.Payload)
}

func (o *GetCreditOK) GetPayload() *models.CreditStatus {
	return o.Payload
}

func (o *GetCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreditStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCreditNotFound creates a GetCreditNotFound with default headers values
func NewGetCreditNotFound() *GetCreditNotFound {
	return &GetCreditNotFound{}
}

/*GetCreditNotFound handles this case with default header values.

The account with the provided id doesn't exist
*/
type GetCreditNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetCreditNotFound) Error() string {
	return fmt.Sprintf("[GET /account/available/{id}][%d] getCreditNotFound  %+v", 404, o.Payload)
}

func (o *GetCreditNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCreditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCreditInternalServerError creates a GetCreditInternalServerError with default headers values
func NewGetCreditInternalServerError() *GetCreditInternalServerError {
	return &GetCreditInternalServerError{}
}

/*GetCreditInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type GetCreditInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetCreditInternalServerError) Error() string {
	return fmt.Sprintf("[GET /account/available/{id}][%d] getCreditInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCreditInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCreditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

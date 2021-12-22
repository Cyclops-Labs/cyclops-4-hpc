// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/udr/models"
)

// ExecCompactationReader is a Reader for the ExecCompactation structure.
type ExecCompactationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecCompactationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecCompactationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewExecCompactationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecCompactationOK creates a ExecCompactationOK with default headers values
func NewExecCompactationOK() *ExecCompactationOK {
	return &ExecCompactationOK{}
}

/*ExecCompactationOK handles this case with default header values.

Compactation task executed successfully.
*/
type ExecCompactationOK struct {
}

func (o *ExecCompactationOK) Error() string {
	return fmt.Sprintf("[GET /trigger/compact][%d] execCompactationOK ", 200)
}

func (o *ExecCompactationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecCompactationInternalServerError creates a ExecCompactationInternalServerError with default headers values
func NewExecCompactationInternalServerError() *ExecCompactationInternalServerError {
	return &ExecCompactationInternalServerError{}
}

/*ExecCompactationInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type ExecCompactationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ExecCompactationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /trigger/compact][%d] execCompactationInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecCompactationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecCompactationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

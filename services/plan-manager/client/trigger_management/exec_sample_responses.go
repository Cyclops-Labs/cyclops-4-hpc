// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/plan-manager/models"
)

// ExecSampleReader is a Reader for the ExecSample structure.
type ExecSampleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecSampleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecSampleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewExecSampleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecSampleOK creates a ExecSampleOK with default headers values
func NewExecSampleOK() *ExecSampleOK {
	return &ExecSampleOK{}
}

/*ExecSampleOK handles this case with default header values.

Sample task executed successfully
*/
type ExecSampleOK struct {
}

func (o *ExecSampleOK) Error() string {
	return fmt.Sprintf("[GET /trigger/sample][%d] execSampleOK ", 200)
}

func (o *ExecSampleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecSampleInternalServerError creates a ExecSampleInternalServerError with default headers values
func NewExecSampleInternalServerError() *ExecSampleInternalServerError {
	return &ExecSampleInternalServerError{}
}

/*ExecSampleInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type ExecSampleInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ExecSampleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /trigger/sample][%d] execSampleInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecSampleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecSampleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

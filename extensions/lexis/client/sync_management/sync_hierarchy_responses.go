// Code generated by go-swagger; DO NOT EDIT.

package sync_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/extensions/lexis/models"
)

// SyncHierarchyReader is a Reader for the SyncHierarchy structure.
type SyncHierarchyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncHierarchyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncHierarchyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSyncHierarchyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSyncHierarchyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncHierarchyOK creates a SyncHierarchyOK with default headers values
func NewSyncHierarchyOK() *SyncHierarchyOK {
	return &SyncHierarchyOK{}
}

/*SyncHierarchyOK handles this case with default header values.

The load of data was completely successfully
*/
type SyncHierarchyOK struct {
}

func (o *SyncHierarchyOK) Error() string {
	return fmt.Sprintf("[GET /sync/hierarchy][%d] syncHierarchyOK ", 200)
}

func (o *SyncHierarchyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncHierarchyAccepted creates a SyncHierarchyAccepted with default headers values
func NewSyncHierarchyAccepted() *SyncHierarchyAccepted {
	return &SyncHierarchyAccepted{}
}

/*SyncHierarchyAccepted handles this case with default header values.

Operation done but there might have been some fails when adding part of the data
*/
type SyncHierarchyAccepted struct {
}

func (o *SyncHierarchyAccepted) Error() string {
	return fmt.Sprintf("[GET /sync/hierarchy][%d] syncHierarchyAccepted ", 202)
}

func (o *SyncHierarchyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncHierarchyInternalServerError creates a SyncHierarchyInternalServerError with default headers values
func NewSyncHierarchyInternalServerError() *SyncHierarchyInternalServerError {
	return &SyncHierarchyInternalServerError{}
}

/*SyncHierarchyInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type SyncHierarchyInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SyncHierarchyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sync/hierarchy][%d] syncHierarchyInternalServerError  %+v", 500, o.Payload)
}

func (o *SyncHierarchyInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SyncHierarchyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

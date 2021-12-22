// Code generated by go-swagger; DO NOT EDIT.

package trigger_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the trigger management client
type API interface {
	/*
	   ExecSample samples task trigger*/
	ExecSample(ctx context.Context, params *ExecSampleParams) (*ExecSampleOK, error)
}

// New creates a new trigger management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for trigger management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ExecSample samples task trigger
*/
func (a *Client) ExecSample(ctx context.Context, params *ExecSampleParams) (*ExecSampleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "execSample",
		Method:             "GET",
		PathPattern:        "/trigger/sample",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExecSampleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExecSampleOK), nil

}
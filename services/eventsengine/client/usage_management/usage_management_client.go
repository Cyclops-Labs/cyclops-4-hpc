// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the usage management client
type API interface {
	/*
	   GetSystemUsage generates an aggregated response by account of the usage recorded in the system during the time window specified*/
	GetSystemUsage(ctx context.Context, params *GetSystemUsageParams) (*GetSystemUsageOK, error)
	/*
	   GetUsage generates an aggregated response of the usage recorded in the system during the time window specified for the selected account*/
	GetUsage(ctx context.Context, params *GetUsageParams) (*GetUsageOK, error)
}

// New creates a new usage management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for usage management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetSystemUsage generates an aggregated response by account of the usage recorded in the system during the time window specified
*/
func (a *Client) GetSystemUsage(ctx context.Context, params *GetSystemUsageParams) (*GetSystemUsageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemUsage",
		Method:             "GET",
		PathPattern:        "/usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSystemUsageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSystemUsageOK), nil

}

/*
GetUsage generates an aggregated response of the usage recorded in the system during the time window specified for the selected account
*/
func (a *Client) GetUsage(ctx context.Context, params *GetUsageParams) (*GetUsageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsage",
		Method:             "GET",
		PathPattern:        "/usage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUsageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsageOK), nil

}
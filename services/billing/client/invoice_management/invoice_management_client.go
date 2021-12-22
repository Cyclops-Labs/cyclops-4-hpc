// Code generated by go-swagger; DO NOT EDIT.

package invoice_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the invoice management client
type API interface {
	/*
	   GenerateInvoiceForCustomer generates invoice for the provided customer for the provided time window or last period*/
	GenerateInvoiceForCustomer(ctx context.Context, params *GenerateInvoiceForCustomerParams) (*GenerateInvoiceForCustomerAccepted, error)
	/*
	   GenerateInvoiceForReseller generates invoice for the provided reseller for the provided time window or last period*/
	GenerateInvoiceForReseller(ctx context.Context, params *GenerateInvoiceForResellerParams) (*GenerateInvoiceForResellerAccepted, error)
	/*
	   GetInvoice summaries for this endpoint*/
	GetInvoice(ctx context.Context, params *GetInvoiceParams) (*GetInvoiceOK, error)
	/*
	   GetInvoicesByCustomer retrieves invoices by customer id*/
	GetInvoicesByCustomer(ctx context.Context, params *GetInvoicesByCustomerParams) (*GetInvoicesByCustomerOK, error)
	/*
	   GetInvoicesByReseller retrieves invoices by reseller id*/
	GetInvoicesByReseller(ctx context.Context, params *GetInvoicesByResellerParams) (*GetInvoicesByResellerOK, error)
	/*
	   ListCustomerInvoices retrieves customers invoices*/
	ListCustomerInvoices(ctx context.Context, params *ListCustomerInvoicesParams) (*ListCustomerInvoicesOK, error)
	/*
	   ListInvoices summaries for this endpoint*/
	ListInvoices(ctx context.Context, params *ListInvoicesParams) (*ListInvoicesOK, error)
	/*
	   ListResellerInvoices retrieves resellers invoices*/
	ListResellerInvoices(ctx context.Context, params *ListResellerInvoicesParams) (*ListResellerInvoicesOK, error)
}

// New creates a new invoice management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for invoice management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GenerateInvoiceForCustomer generates invoice for the provided customer for the provided time window or last period
*/
func (a *Client) GenerateInvoiceForCustomer(ctx context.Context, params *GenerateInvoiceForCustomerParams) (*GenerateInvoiceForCustomerAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GenerateInvoiceForCustomer",
		Method:             "POST",
		PathPattern:        "/invoice/customer/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GenerateInvoiceForCustomerReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateInvoiceForCustomerAccepted), nil

}

/*
GenerateInvoiceForReseller generates invoice for the provided reseller for the provided time window or last period
*/
func (a *Client) GenerateInvoiceForReseller(ctx context.Context, params *GenerateInvoiceForResellerParams) (*GenerateInvoiceForResellerAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GenerateInvoiceForReseller",
		Method:             "POST",
		PathPattern:        "/invoice/reseller/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GenerateInvoiceForResellerReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateInvoiceForResellerAccepted), nil

}

/*
GetInvoice summaries for this endpoint
*/
func (a *Client) GetInvoice(ctx context.Context, params *GetInvoiceParams) (*GetInvoiceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInvoice",
		Method:             "GET",
		PathPattern:        "/invoice/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInvoiceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoiceOK), nil

}

/*
GetInvoicesByCustomer retrieves invoices by customer id
*/
func (a *Client) GetInvoicesByCustomer(ctx context.Context, params *GetInvoicesByCustomerParams) (*GetInvoicesByCustomerOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInvoicesByCustomer",
		Method:             "GET",
		PathPattern:        "/invoice/customer/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInvoicesByCustomerReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoicesByCustomerOK), nil

}

/*
GetInvoicesByReseller retrieves invoices by reseller id
*/
func (a *Client) GetInvoicesByReseller(ctx context.Context, params *GetInvoicesByResellerParams) (*GetInvoicesByResellerOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInvoicesByReseller",
		Method:             "GET",
		PathPattern:        "/invoice/reseller/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInvoicesByResellerReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInvoicesByResellerOK), nil

}

/*
ListCustomerInvoices retrieves customers invoices
*/
func (a *Client) ListCustomerInvoices(ctx context.Context, params *ListCustomerInvoicesParams) (*ListCustomerInvoicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCustomerInvoices",
		Method:             "GET",
		PathPattern:        "/invoice/customer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListCustomerInvoicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListCustomerInvoicesOK), nil

}

/*
ListInvoices summaries for this endpoint
*/
func (a *Client) ListInvoices(ctx context.Context, params *ListInvoicesParams) (*ListInvoicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListInvoices",
		Method:             "GET",
		PathPattern:        "/invoice",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListInvoicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInvoicesOK), nil

}

/*
ListResellerInvoices retrieves resellers invoices
*/
func (a *Client) ListResellerInvoices(ctx context.Context, params *ListResellerInvoicesParams) (*ListResellerInvoicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListResellerInvoices",
		Method:             "GET",
		PathPattern:        "/invoice/reseller",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListResellerInvoicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListResellerInvoicesOK), nil

}
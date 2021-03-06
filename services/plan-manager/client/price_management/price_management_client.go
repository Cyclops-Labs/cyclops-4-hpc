// Code generated by go-swagger; DO NOT EDIT.

package price_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the price management client
type API interface {
	/*
	   CreateSkuPrice creates s k u price

	   Creates a new sku price*/
	CreateSkuPrice(ctx context.Context, params *CreateSkuPriceParams) (*CreateSkuPriceCreated, error)
	/*
	   GetSkuPrice gets specific sku price

	   get sku price with given skupriceid*/
	GetSkuPrice(ctx context.Context, params *GetSkuPriceParams) (*GetSkuPriceOK, error)
	/*
	   ListSkuPrices lists s k u prices

	   lists all sku prices*/
	ListSkuPrices(ctx context.Context, params *ListSkuPricesParams) (*ListSkuPricesOK, error)
	/*
	   UpdateSkuPrice updates specific sku price

	   Update sku price with given skupriceid*/
	UpdateSkuPrice(ctx context.Context, params *UpdateSkuPriceParams) (*UpdateSkuPriceOK, error)
}

// New creates a new price management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for price management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateSkuPrice creates s k u price

Creates a new sku price
*/
func (a *Client) CreateSkuPrice(ctx context.Context, params *CreateSkuPriceParams) (*CreateSkuPriceCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSkuPrice",
		Method:             "POST",
		PathPattern:        "/sku/price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSkuPriceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSkuPriceCreated), nil

}

/*
GetSkuPrice gets specific sku price

get sku price with given skupriceid
*/
func (a *Client) GetSkuPrice(ctx context.Context, params *GetSkuPriceParams) (*GetSkuPriceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSkuPrice",
		Method:             "GET",
		PathPattern:        "/sku/price/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSkuPriceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSkuPriceOK), nil

}

/*
ListSkuPrices lists s k u prices

lists all sku prices
*/
func (a *Client) ListSkuPrices(ctx context.Context, params *ListSkuPricesParams) (*ListSkuPricesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSkuPrices",
		Method:             "GET",
		PathPattern:        "/sku/price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSkuPricesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSkuPricesOK), nil

}

/*
UpdateSkuPrice updates specific sku price

Update sku price with given skupriceid
*/
func (a *Client) UpdateSkuPrice(ctx context.Context, params *UpdateSkuPriceParams) (*UpdateSkuPriceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSkuPrice",
		Method:             "PUT",
		PathPattern:        "/sku/price/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateSkuPriceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSkuPriceOK), nil

}

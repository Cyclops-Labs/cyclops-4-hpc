// Code generated by go-swagger; DO NOT EDIT.

package plan_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the plan management client
type API interface {
	/*
	   CreatePlan creates a plan

	   Creates a new plan*/
	CreatePlan(ctx context.Context, params *CreatePlanParams) (*CreatePlanCreated, error)
	/*
	   GetCompletePlan gets complete plan

	   gets complete plan with planid*/
	GetCompletePlan(ctx context.Context, params *GetCompletePlanParams) (*GetCompletePlanOK, error)
	/*
	   GetPlan gets specific plan

	   get plan with given planid*/
	GetPlan(ctx context.Context, params *GetPlanParams) (*GetPlanOK, error)
	/*
	   ListCompletePlans gets full information relating to known plans

	   Obtains full information on all known plans*/
	ListCompletePlans(ctx context.Context, params *ListCompletePlansParams) (*ListCompletePlansOK, error)
	/*
	   ListPlans lists all plans

	   lists all plans (tbd - pagination?)*/
	ListPlans(ctx context.Context, params *ListPlansParams) (*ListPlansOK, error)
	/*
	   UpdatePlan updates specific plan

	   Update plan with given planId*/
	UpdatePlan(ctx context.Context, params *UpdatePlanParams) (*UpdatePlanOK, error)
}

// New creates a new plan management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for plan management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreatePlan creates a plan

Creates a new plan
*/
func (a *Client) CreatePlan(ctx context.Context, params *CreatePlanParams) (*CreatePlanCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPlan",
		Method:             "POST",
		PathPattern:        "/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePlanCreated), nil

}

/*
GetCompletePlan gets complete plan

gets complete plan with planid
*/
func (a *Client) GetCompletePlan(ctx context.Context, params *GetCompletePlanParams) (*GetCompletePlanOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCompletePlan",
		Method:             "GET",
		PathPattern:        "/plan/complete/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCompletePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCompletePlanOK), nil

}

/*
GetPlan gets specific plan

get plan with given planid
*/
func (a *Client) GetPlan(ctx context.Context, params *GetPlanParams) (*GetPlanOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlan",
		Method:             "GET",
		PathPattern:        "/plan/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlanOK), nil

}

/*
ListCompletePlans gets full information relating to known plans

Obtains full information on all known plans
*/
func (a *Client) ListCompletePlans(ctx context.Context, params *ListCompletePlansParams) (*ListCompletePlansOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCompletePlans",
		Method:             "GET",
		PathPattern:        "/plan/complete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListCompletePlansReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListCompletePlansOK), nil

}

/*
ListPlans lists all plans

lists all plans (tbd - pagination?)
*/
func (a *Client) ListPlans(ctx context.Context, params *ListPlansParams) (*ListPlansOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPlans",
		Method:             "GET",
		PathPattern:        "/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPlansReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPlansOK), nil

}

/*
UpdatePlan updates specific plan

Update plan with given planId
*/
func (a *Client) UpdatePlan(ctx context.Context, params *UpdatePlanParams) (*UpdatePlanOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePlan",
		Method:             "PUT",
		PathPattern:        "/plan/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdatePlanOK), nil

}

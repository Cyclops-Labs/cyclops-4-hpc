// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/bulk_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/invoice_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/trigger_management"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name BulkManagementAPI -inpkg

/* BulkManagementAPI  */
type BulkManagementAPI interface {
	/* GetBillRun Get the status report of the billrun requested */
	GetBillRun(ctx context.Context, params bulk_management.GetBillRunParams) middleware.Responder

	/* ListBillRuns Show the status report of the billruns present in the system */
	ListBillRuns(ctx context.Context, params bulk_management.ListBillRunsParams) middleware.Responder

	/* ListBillRunsByOrganization Show the status report of the billruns present in the system */
	ListBillRunsByOrganization(ctx context.Context, params bulk_management.ListBillRunsByOrganizationParams) middleware.Responder

	/* ReRunAllBillRuns Try to re-run the failed invoices in all the billruns. */
	ReRunAllBillRuns(ctx context.Context, params bulk_management.ReRunAllBillRunsParams) middleware.Responder

	/* ReRunBillRun Try to re-run the failed invoices in the billrun. */
	ReRunBillRun(ctx context.Context, params bulk_management.ReRunBillRunParams) middleware.Responder
}

//go:generate mockery -name InvoiceManagementAPI -inpkg

/* InvoiceManagementAPI  */
type InvoiceManagementAPI interface {
	/* GenerateInvoiceForCustomer Generate invoice for the provided customer for the provided time window or last period */
	GenerateInvoiceForCustomer(ctx context.Context, params invoice_management.GenerateInvoiceForCustomerParams) middleware.Responder

	/* GenerateInvoiceForReseller Generate invoice for the provided reseller for the provided time window or last period */
	GenerateInvoiceForReseller(ctx context.Context, params invoice_management.GenerateInvoiceForResellerParams) middleware.Responder

	/* GetInvoice Summary for this endpoint */
	GetInvoice(ctx context.Context, params invoice_management.GetInvoiceParams) middleware.Responder

	/* GetInvoicesByCustomer Retrieve invoices by customer id */
	GetInvoicesByCustomer(ctx context.Context, params invoice_management.GetInvoicesByCustomerParams) middleware.Responder

	/* GetInvoicesByReseller Retrieve invoices by reseller id */
	GetInvoicesByReseller(ctx context.Context, params invoice_management.GetInvoicesByResellerParams) middleware.Responder

	/* ListCustomerInvoices Retrieve customers' invoices */
	ListCustomerInvoices(ctx context.Context, params invoice_management.ListCustomerInvoicesParams) middleware.Responder

	/* ListInvoices Summary for this endpoint */
	ListInvoices(ctx context.Context, params invoice_management.ListInvoicesParams) middleware.Responder

	/* ListResellerInvoices Retrieve resellers' invoices */
	ListResellerInvoices(ctx context.Context, params invoice_management.ListResellerInvoicesParams) middleware.Responder
}

//go:generate mockery -name StatusManagementAPI -inpkg

/* StatusManagementAPI  */
type StatusManagementAPI interface {
	/* GetStatus Basic status of the system */
	GetStatus(ctx context.Context, params status_management.GetStatusParams) middleware.Responder

	/* ShowStatus Basic status of the system */
	ShowStatus(ctx context.Context, params status_management.ShowStatusParams) middleware.Responder
}

//go:generate mockery -name TriggerManagementAPI -inpkg

/* TriggerManagementAPI  */
type TriggerManagementAPI interface {
	/* PeriodicRun Periodic run of the bulk generation of invoices */
	PeriodicRun(ctx context.Context, params trigger_management.PeriodicRunParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	BulkManagementAPI
	InvoiceManagementAPI
	StatusManagementAPI
	TriggerManagementAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error

	// AuthAPIKeyHeader Applies when the "X-API-KEY" header is set
	AuthAPIKeyHeader func(token string) (interface{}, error)

	// AuthAPIKeyParam Applies when the "api_key" query is set
	AuthAPIKeyParam func(token string) (interface{}, error)

	// AuthKeycloak For OAuth2 authentication
	AuthKeycloak func(token string, scopes []string) (interface{}, error)
	// Authenticator to use for all APIKey authentication
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// Authenticator to use for all Bearer authentication
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// Authenticator to use for all Basic authentication
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *BillingManagementAPI instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.BillingManagementAPIAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewBillingManagementAPIAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	if c.APIKeyAuthenticator != nil {
		api.APIKeyAuthenticator = c.APIKeyAuthenticator
	}
	if c.BasicAuthenticator != nil {
		api.BasicAuthenticator = c.BasicAuthenticator
	}
	if c.BearerAuthenticator != nil {
		api.BearerAuthenticator = c.BearerAuthenticator
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	api.APIKeyHeaderAuth = func(token string) (interface{}, error) {
		if c.AuthAPIKeyHeader == nil {
			return token, nil
		}
		return c.AuthAPIKeyHeader(token)
	}

	api.APIKeyParamAuth = func(token string) (interface{}, error) {
		if c.AuthAPIKeyParam == nil {
			return token, nil
		}
		return c.AuthAPIKeyParam(token)
	}

	api.KeycloakAuth = func(token string, scopes []string) (interface{}, error) {
		if c.AuthKeycloak == nil {
			return token, nil
		}
		return c.AuthKeycloak(token, scopes)
	}
	api.APIAuthorizer = authorizer(c.Authorizer)
	api.InvoiceManagementGenerateInvoiceForCustomerHandler = invoice_management.GenerateInvoiceForCustomerHandlerFunc(func(params invoice_management.GenerateInvoiceForCustomerParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.GenerateInvoiceForCustomer(ctx, params)
	})
	api.InvoiceManagementGenerateInvoiceForResellerHandler = invoice_management.GenerateInvoiceForResellerHandlerFunc(func(params invoice_management.GenerateInvoiceForResellerParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.GenerateInvoiceForReseller(ctx, params)
	})
	api.BulkManagementGetBillRunHandler = bulk_management.GetBillRunHandlerFunc(func(params bulk_management.GetBillRunParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.BulkManagementAPI.GetBillRun(ctx, params)
	})
	api.InvoiceManagementGetInvoiceHandler = invoice_management.GetInvoiceHandlerFunc(func(params invoice_management.GetInvoiceParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.GetInvoice(ctx, params)
	})
	api.InvoiceManagementGetInvoicesByCustomerHandler = invoice_management.GetInvoicesByCustomerHandlerFunc(func(params invoice_management.GetInvoicesByCustomerParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.GetInvoicesByCustomer(ctx, params)
	})
	api.InvoiceManagementGetInvoicesByResellerHandler = invoice_management.GetInvoicesByResellerHandlerFunc(func(params invoice_management.GetInvoicesByResellerParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.GetInvoicesByReseller(ctx, params)
	})
	api.BulkManagementListBillRunsHandler = bulk_management.ListBillRunsHandlerFunc(func(params bulk_management.ListBillRunsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.BulkManagementAPI.ListBillRuns(ctx, params)
	})
	api.BulkManagementListBillRunsByOrganizationHandler = bulk_management.ListBillRunsByOrganizationHandlerFunc(func(params bulk_management.ListBillRunsByOrganizationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.BulkManagementAPI.ListBillRunsByOrganization(ctx, params)
	})
	api.InvoiceManagementListCustomerInvoicesHandler = invoice_management.ListCustomerInvoicesHandlerFunc(func(params invoice_management.ListCustomerInvoicesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.ListCustomerInvoices(ctx, params)
	})
	api.InvoiceManagementListInvoicesHandler = invoice_management.ListInvoicesHandlerFunc(func(params invoice_management.ListInvoicesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.ListInvoices(ctx, params)
	})
	api.InvoiceManagementListResellerInvoicesHandler = invoice_management.ListResellerInvoicesHandlerFunc(func(params invoice_management.ListResellerInvoicesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InvoiceManagementAPI.ListResellerInvoices(ctx, params)
	})
	api.BulkManagementReRunAllBillRunsHandler = bulk_management.ReRunAllBillRunsHandlerFunc(func(params bulk_management.ReRunAllBillRunsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.BulkManagementAPI.ReRunAllBillRuns(ctx, params)
	})
	api.BulkManagementReRunBillRunHandler = bulk_management.ReRunBillRunHandlerFunc(func(params bulk_management.ReRunBillRunParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.BulkManagementAPI.ReRunBillRun(ctx, params)
	})
	api.StatusManagementGetStatusHandler = status_management.GetStatusHandlerFunc(func(params status_management.GetStatusParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.StatusManagementAPI.GetStatus(ctx, params)
	})
	api.TriggerManagementPeriodicRunHandler = trigger_management.PeriodicRunHandlerFunc(func(params trigger_management.PeriodicRunParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.TriggerManagementAPI.PeriodicRun(ctx, params)
	})
	api.StatusManagementShowStatusHandler = status_management.ShowStatusHandlerFunc(func(params status_management.ShowStatusParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.StatusManagementAPI.ShowStatus(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}

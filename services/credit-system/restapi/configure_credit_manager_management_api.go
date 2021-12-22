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

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/restapi/operations"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/restapi/operations/account_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/restapi/operations/credit_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/restapi/operations/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/credit-system/restapi/operations/trigger_management"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name AccountManagementAPI -inpkg

/* AccountManagementAPI  */
type AccountManagementAPI interface {
	/* CreateAccount Creates a new account in the system */
	CreateAccount(ctx context.Context, params account_management.CreateAccountParams) middleware.Responder

	/* DisableAccount Disables the account with the id provided in the system */
	DisableAccount(ctx context.Context, params account_management.DisableAccountParams) middleware.Responder

	/* EnableAccount Enables the account with the id provided in the system */
	EnableAccount(ctx context.Context, params account_management.EnableAccountParams) middleware.Responder

	/* GetAccountStatus Basic status of the account with the id provided in the system */
	GetAccountStatus(ctx context.Context, params account_management.GetAccountStatusParams) middleware.Responder

	/* ListAccounts List of the accounts in the system */
	ListAccounts(ctx context.Context, params account_management.ListAccountsParams) middleware.Responder
}

//go:generate mockery -name CreditManagementAPI -inpkg

/* CreditManagementAPI  */
type CreditManagementAPI interface {
	/* AddConsumption Adds a consumption to the system */
	AddConsumption(ctx context.Context, params credit_management.AddConsumptionParams) middleware.Responder

	/* DecreaseCredit Insert a new reseller in the system. */
	DecreaseCredit(ctx context.Context, params credit_management.DecreaseCreditParams) middleware.Responder

	/* GetCredit Credit status of the account with the provided id */
	GetCredit(ctx context.Context, params credit_management.GetCreditParams) middleware.Responder

	/* GetHistory Credit history of the customer with id */
	GetHistory(ctx context.Context, params credit_management.GetHistoryParams) middleware.Responder

	/* IncreaseCredit Insert a new reseller in the system. */
	IncreaseCredit(ctx context.Context, params credit_management.IncreaseCreditParams) middleware.Responder
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
	/* ExecSample Sample task trigger */
	ExecSample(ctx context.Context, params trigger_management.ExecSampleParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	AccountManagementAPI
	CreditManagementAPI
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
// and the corresponding *CreditManagerManagementAPI instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.CreditManagerManagementAPIAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewCreditManagerManagementAPIAPI(spec)
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
	api.CreditManagementAddConsumptionHandler = credit_management.AddConsumptionHandlerFunc(func(params credit_management.AddConsumptionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.CreditManagementAPI.AddConsumption(ctx, params)
	})
	api.AccountManagementCreateAccountHandler = account_management.CreateAccountHandlerFunc(func(params account_management.CreateAccountParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.AccountManagementAPI.CreateAccount(ctx, params)
	})
	api.CreditManagementDecreaseCreditHandler = credit_management.DecreaseCreditHandlerFunc(func(params credit_management.DecreaseCreditParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.CreditManagementAPI.DecreaseCredit(ctx, params)
	})
	api.AccountManagementDisableAccountHandler = account_management.DisableAccountHandlerFunc(func(params account_management.DisableAccountParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.AccountManagementAPI.DisableAccount(ctx, params)
	})
	api.AccountManagementEnableAccountHandler = account_management.EnableAccountHandlerFunc(func(params account_management.EnableAccountParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.AccountManagementAPI.EnableAccount(ctx, params)
	})
	api.TriggerManagementExecSampleHandler = trigger_management.ExecSampleHandlerFunc(func(params trigger_management.ExecSampleParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.TriggerManagementAPI.ExecSample(ctx, params)
	})
	api.AccountManagementGetAccountStatusHandler = account_management.GetAccountStatusHandlerFunc(func(params account_management.GetAccountStatusParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.AccountManagementAPI.GetAccountStatus(ctx, params)
	})
	api.CreditManagementGetCreditHandler = credit_management.GetCreditHandlerFunc(func(params credit_management.GetCreditParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.CreditManagementAPI.GetCredit(ctx, params)
	})
	api.CreditManagementGetHistoryHandler = credit_management.GetHistoryHandlerFunc(func(params credit_management.GetHistoryParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.CreditManagementAPI.GetHistory(ctx, params)
	})
	api.StatusManagementGetStatusHandler = status_management.GetStatusHandlerFunc(func(params status_management.GetStatusParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.StatusManagementAPI.GetStatus(ctx, params)
	})
	api.CreditManagementIncreaseCreditHandler = credit_management.IncreaseCreditHandlerFunc(func(params credit_management.IncreaseCreditParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.CreditManagementAPI.IncreaseCredit(ctx, params)
	})
	api.AccountManagementListAccountsHandler = account_management.ListAccountsHandlerFunc(func(params account_management.ListAccountsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.AccountManagementAPI.ListAccounts(ctx, params)
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

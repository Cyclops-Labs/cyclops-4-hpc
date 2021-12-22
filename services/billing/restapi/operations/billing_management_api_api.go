// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/bulk_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/invoice_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/billing/restapi/operations/trigger_management"
)

// NewBillingManagementAPIAPI creates a new BillingManagementAPI instance
func NewBillingManagementAPIAPI(spec *loads.Document) *BillingManagementAPIAPI {
	return &BillingManagementAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		InvoiceManagementGenerateInvoiceForCustomerHandler: invoice_management.GenerateInvoiceForCustomerHandlerFunc(func(params invoice_management.GenerateInvoiceForCustomerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.GenerateInvoiceForCustomer has not yet been implemented")
		}),
		InvoiceManagementGenerateInvoiceForResellerHandler: invoice_management.GenerateInvoiceForResellerHandlerFunc(func(params invoice_management.GenerateInvoiceForResellerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.GenerateInvoiceForReseller has not yet been implemented")
		}),
		BulkManagementGetBillRunHandler: bulk_management.GetBillRunHandlerFunc(func(params bulk_management.GetBillRunParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bulk_management.GetBillRun has not yet been implemented")
		}),
		InvoiceManagementGetInvoiceHandler: invoice_management.GetInvoiceHandlerFunc(func(params invoice_management.GetInvoiceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.GetInvoice has not yet been implemented")
		}),
		InvoiceManagementGetInvoicesByCustomerHandler: invoice_management.GetInvoicesByCustomerHandlerFunc(func(params invoice_management.GetInvoicesByCustomerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.GetInvoicesByCustomer has not yet been implemented")
		}),
		InvoiceManagementGetInvoicesByResellerHandler: invoice_management.GetInvoicesByResellerHandlerFunc(func(params invoice_management.GetInvoicesByResellerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.GetInvoicesByReseller has not yet been implemented")
		}),
		BulkManagementListBillRunsHandler: bulk_management.ListBillRunsHandlerFunc(func(params bulk_management.ListBillRunsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bulk_management.ListBillRuns has not yet been implemented")
		}),
		BulkManagementListBillRunsByOrganizationHandler: bulk_management.ListBillRunsByOrganizationHandlerFunc(func(params bulk_management.ListBillRunsByOrganizationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bulk_management.ListBillRunsByOrganization has not yet been implemented")
		}),
		InvoiceManagementListCustomerInvoicesHandler: invoice_management.ListCustomerInvoicesHandlerFunc(func(params invoice_management.ListCustomerInvoicesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.ListCustomerInvoices has not yet been implemented")
		}),
		InvoiceManagementListInvoicesHandler: invoice_management.ListInvoicesHandlerFunc(func(params invoice_management.ListInvoicesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.ListInvoices has not yet been implemented")
		}),
		InvoiceManagementListResellerInvoicesHandler: invoice_management.ListResellerInvoicesHandlerFunc(func(params invoice_management.ListResellerInvoicesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation invoice_management.ListResellerInvoices has not yet been implemented")
		}),
		BulkManagementReRunAllBillRunsHandler: bulk_management.ReRunAllBillRunsHandlerFunc(func(params bulk_management.ReRunAllBillRunsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bulk_management.ReRunAllBillRuns has not yet been implemented")
		}),
		BulkManagementReRunBillRunHandler: bulk_management.ReRunBillRunHandlerFunc(func(params bulk_management.ReRunBillRunParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bulk_management.ReRunBillRun has not yet been implemented")
		}),
		StatusManagementGetStatusHandler: status_management.GetStatusHandlerFunc(func(params status_management.GetStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation status_management.GetStatus has not yet been implemented")
		}),
		TriggerManagementPeriodicRunHandler: trigger_management.PeriodicRunHandlerFunc(func(params trigger_management.PeriodicRunParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation trigger_management.PeriodicRun has not yet been implemented")
		}),
		StatusManagementShowStatusHandler: status_management.ShowStatusHandlerFunc(func(params status_management.ShowStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation status_management.ShowStatus has not yet been implemented")
		}),

		// Applies when the "X-API-KEY" header is set
		APIKeyHeaderAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (APIKeyHeader) X-API-KEY from header param [X-API-KEY] has not yet been implemented")
		},
		// Applies when the "api_key" query is set
		APIKeyParamAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (APIKeyParam) api_key from query param [api_key] has not yet been implemented")
		},
		KeycloakAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (Keycloak) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*BillingManagementAPIAPI An API which supports creation, deletion, listing etc of Billing */
type BillingManagementAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyHeaderAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-API-KEY provided in the header
	APIKeyHeaderAuth func(string) (interface{}, error)

	// APIKeyParamAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key api_key provided in the query
	APIKeyParamAuth func(string) (interface{}, error)

	// KeycloakAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	KeycloakAuth func(string, []string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// InvoiceManagementGenerateInvoiceForCustomerHandler sets the operation handler for the generate invoice for customer operation
	InvoiceManagementGenerateInvoiceForCustomerHandler invoice_management.GenerateInvoiceForCustomerHandler
	// InvoiceManagementGenerateInvoiceForResellerHandler sets the operation handler for the generate invoice for reseller operation
	InvoiceManagementGenerateInvoiceForResellerHandler invoice_management.GenerateInvoiceForResellerHandler
	// BulkManagementGetBillRunHandler sets the operation handler for the get bill run operation
	BulkManagementGetBillRunHandler bulk_management.GetBillRunHandler
	// InvoiceManagementGetInvoiceHandler sets the operation handler for the get invoice operation
	InvoiceManagementGetInvoiceHandler invoice_management.GetInvoiceHandler
	// InvoiceManagementGetInvoicesByCustomerHandler sets the operation handler for the get invoices by customer operation
	InvoiceManagementGetInvoicesByCustomerHandler invoice_management.GetInvoicesByCustomerHandler
	// InvoiceManagementGetInvoicesByResellerHandler sets the operation handler for the get invoices by reseller operation
	InvoiceManagementGetInvoicesByResellerHandler invoice_management.GetInvoicesByResellerHandler
	// BulkManagementListBillRunsHandler sets the operation handler for the list bill runs operation
	BulkManagementListBillRunsHandler bulk_management.ListBillRunsHandler
	// BulkManagementListBillRunsByOrganizationHandler sets the operation handler for the list bill runs by organization operation
	BulkManagementListBillRunsByOrganizationHandler bulk_management.ListBillRunsByOrganizationHandler
	// InvoiceManagementListCustomerInvoicesHandler sets the operation handler for the list customer invoices operation
	InvoiceManagementListCustomerInvoicesHandler invoice_management.ListCustomerInvoicesHandler
	// InvoiceManagementListInvoicesHandler sets the operation handler for the list invoices operation
	InvoiceManagementListInvoicesHandler invoice_management.ListInvoicesHandler
	// InvoiceManagementListResellerInvoicesHandler sets the operation handler for the list reseller invoices operation
	InvoiceManagementListResellerInvoicesHandler invoice_management.ListResellerInvoicesHandler
	// BulkManagementReRunAllBillRunsHandler sets the operation handler for the re run all bill runs operation
	BulkManagementReRunAllBillRunsHandler bulk_management.ReRunAllBillRunsHandler
	// BulkManagementReRunBillRunHandler sets the operation handler for the re run bill run operation
	BulkManagementReRunBillRunHandler bulk_management.ReRunBillRunHandler
	// StatusManagementGetStatusHandler sets the operation handler for the get status operation
	StatusManagementGetStatusHandler status_management.GetStatusHandler
	// TriggerManagementPeriodicRunHandler sets the operation handler for the periodic run operation
	TriggerManagementPeriodicRunHandler trigger_management.PeriodicRunHandler
	// StatusManagementShowStatusHandler sets the operation handler for the show status operation
	StatusManagementShowStatusHandler status_management.ShowStatusHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BillingManagementAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BillingManagementAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BillingManagementAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BillingManagementAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BillingManagementAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BillingManagementAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BillingManagementAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BillingManagementAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BillingManagementAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BillingManagementAPIAPI
func (o *BillingManagementAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyHeaderAuth == nil {
		unregistered = append(unregistered, "XAPIKEYAuth")
	}
	if o.APIKeyParamAuth == nil {
		unregistered = append(unregistered, "APIKeyAuth")
	}
	if o.KeycloakAuth == nil {
		unregistered = append(unregistered, "KeycloakAuth")
	}

	if o.InvoiceManagementGenerateInvoiceForCustomerHandler == nil {
		unregistered = append(unregistered, "invoice_management.GenerateInvoiceForCustomerHandler")
	}
	if o.InvoiceManagementGenerateInvoiceForResellerHandler == nil {
		unregistered = append(unregistered, "invoice_management.GenerateInvoiceForResellerHandler")
	}
	if o.BulkManagementGetBillRunHandler == nil {
		unregistered = append(unregistered, "bulk_management.GetBillRunHandler")
	}
	if o.InvoiceManagementGetInvoiceHandler == nil {
		unregistered = append(unregistered, "invoice_management.GetInvoiceHandler")
	}
	if o.InvoiceManagementGetInvoicesByCustomerHandler == nil {
		unregistered = append(unregistered, "invoice_management.GetInvoicesByCustomerHandler")
	}
	if o.InvoiceManagementGetInvoicesByResellerHandler == nil {
		unregistered = append(unregistered, "invoice_management.GetInvoicesByResellerHandler")
	}
	if o.BulkManagementListBillRunsHandler == nil {
		unregistered = append(unregistered, "bulk_management.ListBillRunsHandler")
	}
	if o.BulkManagementListBillRunsByOrganizationHandler == nil {
		unregistered = append(unregistered, "bulk_management.ListBillRunsByOrganizationHandler")
	}
	if o.InvoiceManagementListCustomerInvoicesHandler == nil {
		unregistered = append(unregistered, "invoice_management.ListCustomerInvoicesHandler")
	}
	if o.InvoiceManagementListInvoicesHandler == nil {
		unregistered = append(unregistered, "invoice_management.ListInvoicesHandler")
	}
	if o.InvoiceManagementListResellerInvoicesHandler == nil {
		unregistered = append(unregistered, "invoice_management.ListResellerInvoicesHandler")
	}
	if o.BulkManagementReRunAllBillRunsHandler == nil {
		unregistered = append(unregistered, "bulk_management.ReRunAllBillRunsHandler")
	}
	if o.BulkManagementReRunBillRunHandler == nil {
		unregistered = append(unregistered, "bulk_management.ReRunBillRunHandler")
	}
	if o.StatusManagementGetStatusHandler == nil {
		unregistered = append(unregistered, "status_management.GetStatusHandler")
	}
	if o.TriggerManagementPeriodicRunHandler == nil {
		unregistered = append(unregistered, "trigger_management.PeriodicRunHandler")
	}
	if o.StatusManagementShowStatusHandler == nil {
		unregistered = append(unregistered, "status_management.ShowStatusHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BillingManagementAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BillingManagementAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "APIKeyHeader":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyHeaderAuth)

		case "APIKeyParam":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyParamAuth)

		case "Keycloak":
			result[name] = o.BearerAuthenticator(name, o.KeycloakAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *BillingManagementAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BillingManagementAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BillingManagementAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BillingManagementAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the billing management API API
func (o *BillingManagementAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BillingManagementAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/invoice/customer/{id}"] = invoice_management.NewGenerateInvoiceForCustomer(o.context, o.InvoiceManagementGenerateInvoiceForCustomerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/invoice/reseller/{id}"] = invoice_management.NewGenerateInvoiceForReseller(o.context, o.InvoiceManagementGenerateInvoiceForResellerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/billrun/{id}"] = bulk_management.NewGetBillRun(o.context, o.BulkManagementGetBillRunHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice/{id}"] = invoice_management.NewGetInvoice(o.context, o.InvoiceManagementGetInvoiceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice/customer/{id}"] = invoice_management.NewGetInvoicesByCustomer(o.context, o.InvoiceManagementGetInvoicesByCustomerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice/reseller/{id}"] = invoice_management.NewGetInvoicesByReseller(o.context, o.InvoiceManagementGetInvoicesByResellerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/billrun"] = bulk_management.NewListBillRuns(o.context, o.BulkManagementListBillRunsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/billrun/organization/{id}"] = bulk_management.NewListBillRunsByOrganization(o.context, o.BulkManagementListBillRunsByOrganizationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice/customer"] = invoice_management.NewListCustomerInvoices(o.context, o.InvoiceManagementListCustomerInvoicesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice"] = invoice_management.NewListInvoices(o.context, o.InvoiceManagementListInvoicesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/invoice/reseller"] = invoice_management.NewListResellerInvoices(o.context, o.InvoiceManagementListResellerInvoicesHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/billrun"] = bulk_management.NewReRunAllBillRuns(o.context, o.BulkManagementReRunAllBillRunsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/billrun/{id}"] = bulk_management.NewReRunBillRun(o.context, o.BulkManagementReRunBillRunHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status/{id}"] = status_management.NewGetStatus(o.context, o.StatusManagementGetStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/trigger/periodicrun"] = trigger_management.NewPeriodicRun(o.context, o.TriggerManagementPeriodicRunHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = status_management.NewShowStatus(o.context, o.StatusManagementShowStatusHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BillingManagementAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BillingManagementAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BillingManagementAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BillingManagementAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BillingManagementAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

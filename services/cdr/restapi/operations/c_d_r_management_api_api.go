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

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/cdr/restapi/operations/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/cdr/restapi/operations/trigger_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/cdr/restapi/operations/usage_management"
)

// NewCDRManagementAPIAPI creates a new CDRManagementAPI instance
func NewCDRManagementAPIAPI(spec *loads.Document) *CDRManagementAPIAPI {
	return &CDRManagementAPIAPI{
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

		TriggerManagementExecTransformationHandler: trigger_management.ExecTransformationHandlerFunc(func(params trigger_management.ExecTransformationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation trigger_management.ExecTransformation has not yet been implemented")
		}),
		StatusManagementGetStatusHandler: status_management.GetStatusHandlerFunc(func(params status_management.GetStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation status_management.GetStatus has not yet been implemented")
		}),
		UsageManagementGetSystemUsageHandler: usage_management.GetSystemUsageHandlerFunc(func(params usage_management.GetSystemUsageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation usage_management.GetSystemUsage has not yet been implemented")
		}),
		UsageManagementGetUsageHandler: usage_management.GetUsageHandlerFunc(func(params usage_management.GetUsageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation usage_management.GetUsage has not yet been implemented")
		}),
		UsageManagementGetUsageSummaryHandler: usage_management.GetUsageSummaryHandlerFunc(func(params usage_management.GetUsageSummaryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation usage_management.GetUsageSummary has not yet been implemented")
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

/*CDRManagementAPIAPI An API which supports creation, deletion, listing etc of CDR */
type CDRManagementAPIAPI struct {
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

	// TriggerManagementExecTransformationHandler sets the operation handler for the exec transformation operation
	TriggerManagementExecTransformationHandler trigger_management.ExecTransformationHandler
	// StatusManagementGetStatusHandler sets the operation handler for the get status operation
	StatusManagementGetStatusHandler status_management.GetStatusHandler
	// UsageManagementGetSystemUsageHandler sets the operation handler for the get system usage operation
	UsageManagementGetSystemUsageHandler usage_management.GetSystemUsageHandler
	// UsageManagementGetUsageHandler sets the operation handler for the get usage operation
	UsageManagementGetUsageHandler usage_management.GetUsageHandler
	// UsageManagementGetUsageSummaryHandler sets the operation handler for the get usage summary operation
	UsageManagementGetUsageSummaryHandler usage_management.GetUsageSummaryHandler
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
func (o *CDRManagementAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CDRManagementAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CDRManagementAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CDRManagementAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CDRManagementAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CDRManagementAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CDRManagementAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CDRManagementAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CDRManagementAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CDRManagementAPIAPI
func (o *CDRManagementAPIAPI) Validate() error {
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

	if o.TriggerManagementExecTransformationHandler == nil {
		unregistered = append(unregistered, "trigger_management.ExecTransformationHandler")
	}
	if o.StatusManagementGetStatusHandler == nil {
		unregistered = append(unregistered, "status_management.GetStatusHandler")
	}
	if o.UsageManagementGetSystemUsageHandler == nil {
		unregistered = append(unregistered, "usage_management.GetSystemUsageHandler")
	}
	if o.UsageManagementGetUsageHandler == nil {
		unregistered = append(unregistered, "usage_management.GetUsageHandler")
	}
	if o.UsageManagementGetUsageSummaryHandler == nil {
		unregistered = append(unregistered, "usage_management.GetUsageSummaryHandler")
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
func (o *CDRManagementAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CDRManagementAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
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
func (o *CDRManagementAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CDRManagementAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *CDRManagementAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *CDRManagementAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the c d r management API API
func (o *CDRManagementAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CDRManagementAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/trigger/transform"] = trigger_management.NewExecTransformation(o.context, o.TriggerManagementExecTransformationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status/{id}"] = status_management.NewGetStatus(o.context, o.StatusManagementGetStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/usage"] = usage_management.NewGetSystemUsage(o.context, o.UsageManagementGetSystemUsageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/usage/{id}"] = usage_management.NewGetUsage(o.context, o.UsageManagementGetUsageHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/usage/summary/{id}"] = usage_management.NewGetUsageSummary(o.context, o.UsageManagementGetUsageSummaryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = status_management.NewShowStatus(o.context, o.StatusManagementShowStatusHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CDRManagementAPIAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *CDRManagementAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CDRManagementAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CDRManagementAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CDRManagementAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

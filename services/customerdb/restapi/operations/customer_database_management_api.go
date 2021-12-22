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

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/restapi/operations/customer_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/restapi/operations/product_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/restapi/operations/reseller_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/restapi/operations/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/customerdb/restapi/operations/trigger_management"
)

// NewCustomerDatabaseManagementAPI creates a new CustomerDatabaseManagement instance
func NewCustomerDatabaseManagementAPI(spec *loads.Document) *CustomerDatabaseManagementAPI {
	return &CustomerDatabaseManagementAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CustomerManagementAddCustomerHandler: customer_management.AddCustomerHandlerFunc(func(params customer_management.AddCustomerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer_management.AddCustomer has not yet been implemented")
		}),
		ProductManagementAddProductHandler: product_management.AddProductHandlerFunc(func(params product_management.AddProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation product_management.AddProduct has not yet been implemented")
		}),
		ResellerManagementAddResellerHandler: reseller_management.AddResellerHandlerFunc(func(params reseller_management.AddResellerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reseller_management.AddReseller has not yet been implemented")
		}),
		TriggerManagementExecSampleHandler: trigger_management.ExecSampleHandlerFunc(func(params trigger_management.ExecSampleParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation trigger_management.ExecSample has not yet been implemented")
		}),
		CustomerManagementGetCustomerHandler: customer_management.GetCustomerHandlerFunc(func(params customer_management.GetCustomerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer_management.GetCustomer has not yet been implemented")
		}),
		ProductManagementGetProductHandler: product_management.GetProductHandlerFunc(func(params product_management.GetProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation product_management.GetProduct has not yet been implemented")
		}),
		ResellerManagementGetResellerHandler: reseller_management.GetResellerHandlerFunc(func(params reseller_management.GetResellerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reseller_management.GetReseller has not yet been implemented")
		}),
		StatusManagementGetStatusHandler: status_management.GetStatusHandlerFunc(func(params status_management.GetStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation status_management.GetStatus has not yet been implemented")
		}),
		CustomerManagementListCustomersHandler: customer_management.ListCustomersHandlerFunc(func(params customer_management.ListCustomersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer_management.ListCustomers has not yet been implemented")
		}),
		ProductManagementListProductsHandler: product_management.ListProductsHandlerFunc(func(params product_management.ListProductsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation product_management.ListProducts has not yet been implemented")
		}),
		ResellerManagementListResellersHandler: reseller_management.ListResellersHandlerFunc(func(params reseller_management.ListResellersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reseller_management.ListResellers has not yet been implemented")
		}),
		StatusManagementShowStatusHandler: status_management.ShowStatusHandlerFunc(func(params status_management.ShowStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation status_management.ShowStatus has not yet been implemented")
		}),
		CustomerManagementUpdateCustomerHandler: customer_management.UpdateCustomerHandlerFunc(func(params customer_management.UpdateCustomerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation customer_management.UpdateCustomer has not yet been implemented")
		}),
		ProductManagementUpdateProductHandler: product_management.UpdateProductHandlerFunc(func(params product_management.UpdateProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation product_management.UpdateProduct has not yet been implemented")
		}),
		ResellerManagementUpdateResellerHandler: reseller_management.UpdateResellerHandlerFunc(func(params reseller_management.UpdateResellerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reseller_management.UpdateReseller has not yet been implemented")
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

/*CustomerDatabaseManagementAPI An API which supports creation, deletion, listing etc of customers and products */
type CustomerDatabaseManagementAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

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

	// CustomerManagementAddCustomerHandler sets the operation handler for the add customer operation
	CustomerManagementAddCustomerHandler customer_management.AddCustomerHandler
	// ProductManagementAddProductHandler sets the operation handler for the add product operation
	ProductManagementAddProductHandler product_management.AddProductHandler
	// ResellerManagementAddResellerHandler sets the operation handler for the add reseller operation
	ResellerManagementAddResellerHandler reseller_management.AddResellerHandler
	// TriggerManagementExecSampleHandler sets the operation handler for the exec sample operation
	TriggerManagementExecSampleHandler trigger_management.ExecSampleHandler
	// CustomerManagementGetCustomerHandler sets the operation handler for the get customer operation
	CustomerManagementGetCustomerHandler customer_management.GetCustomerHandler
	// ProductManagementGetProductHandler sets the operation handler for the get product operation
	ProductManagementGetProductHandler product_management.GetProductHandler
	// ResellerManagementGetResellerHandler sets the operation handler for the get reseller operation
	ResellerManagementGetResellerHandler reseller_management.GetResellerHandler
	// StatusManagementGetStatusHandler sets the operation handler for the get status operation
	StatusManagementGetStatusHandler status_management.GetStatusHandler
	// CustomerManagementListCustomersHandler sets the operation handler for the list customers operation
	CustomerManagementListCustomersHandler customer_management.ListCustomersHandler
	// ProductManagementListProductsHandler sets the operation handler for the list products operation
	ProductManagementListProductsHandler product_management.ListProductsHandler
	// ResellerManagementListResellersHandler sets the operation handler for the list resellers operation
	ResellerManagementListResellersHandler reseller_management.ListResellersHandler
	// StatusManagementShowStatusHandler sets the operation handler for the show status operation
	StatusManagementShowStatusHandler status_management.ShowStatusHandler
	// CustomerManagementUpdateCustomerHandler sets the operation handler for the update customer operation
	CustomerManagementUpdateCustomerHandler customer_management.UpdateCustomerHandler
	// ProductManagementUpdateProductHandler sets the operation handler for the update product operation
	ProductManagementUpdateProductHandler product_management.UpdateProductHandler
	// ResellerManagementUpdateResellerHandler sets the operation handler for the update reseller operation
	ResellerManagementUpdateResellerHandler reseller_management.UpdateResellerHandler
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

// SetDefaultProduces sets the default produces media type
func (o *CustomerDatabaseManagementAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CustomerDatabaseManagementAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CustomerDatabaseManagementAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CustomerDatabaseManagementAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CustomerDatabaseManagementAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CustomerDatabaseManagementAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CustomerDatabaseManagementAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CustomerDatabaseManagementAPI
func (o *CustomerDatabaseManagementAPI) Validate() error {
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

	if o.CustomerManagementAddCustomerHandler == nil {
		unregistered = append(unregistered, "customer_management.AddCustomerHandler")
	}
	if o.ProductManagementAddProductHandler == nil {
		unregistered = append(unregistered, "product_management.AddProductHandler")
	}
	if o.ResellerManagementAddResellerHandler == nil {
		unregistered = append(unregistered, "reseller_management.AddResellerHandler")
	}
	if o.TriggerManagementExecSampleHandler == nil {
		unregistered = append(unregistered, "trigger_management.ExecSampleHandler")
	}
	if o.CustomerManagementGetCustomerHandler == nil {
		unregistered = append(unregistered, "customer_management.GetCustomerHandler")
	}
	if o.ProductManagementGetProductHandler == nil {
		unregistered = append(unregistered, "product_management.GetProductHandler")
	}
	if o.ResellerManagementGetResellerHandler == nil {
		unregistered = append(unregistered, "reseller_management.GetResellerHandler")
	}
	if o.StatusManagementGetStatusHandler == nil {
		unregistered = append(unregistered, "status_management.GetStatusHandler")
	}
	if o.CustomerManagementListCustomersHandler == nil {
		unregistered = append(unregistered, "customer_management.ListCustomersHandler")
	}
	if o.ProductManagementListProductsHandler == nil {
		unregistered = append(unregistered, "product_management.ListProductsHandler")
	}
	if o.ResellerManagementListResellersHandler == nil {
		unregistered = append(unregistered, "reseller_management.ListResellersHandler")
	}
	if o.StatusManagementShowStatusHandler == nil {
		unregistered = append(unregistered, "status_management.ShowStatusHandler")
	}
	if o.CustomerManagementUpdateCustomerHandler == nil {
		unregistered = append(unregistered, "customer_management.UpdateCustomerHandler")
	}
	if o.ProductManagementUpdateProductHandler == nil {
		unregistered = append(unregistered, "product_management.UpdateProductHandler")
	}
	if o.ResellerManagementUpdateResellerHandler == nil {
		unregistered = append(unregistered, "reseller_management.UpdateResellerHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CustomerDatabaseManagementAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CustomerDatabaseManagementAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
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
func (o *CustomerDatabaseManagementAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CustomerDatabaseManagementAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *CustomerDatabaseManagementAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *CustomerDatabaseManagementAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the customer database management API
func (o *CustomerDatabaseManagementAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CustomerDatabaseManagementAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/customer"] = customer_management.NewAddCustomer(o.context, o.CustomerManagementAddCustomerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/product"] = product_management.NewAddProduct(o.context, o.ProductManagementAddProductHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/reseller"] = reseller_management.NewAddReseller(o.context, o.ResellerManagementAddResellerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/trigger/sample"] = trigger_management.NewExecSample(o.context, o.TriggerManagementExecSampleHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/customer/{id}"] = customer_management.NewGetCustomer(o.context, o.CustomerManagementGetCustomerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product/{id}"] = product_management.NewGetProduct(o.context, o.ProductManagementGetProductHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/reseller/{id}"] = reseller_management.NewGetReseller(o.context, o.ResellerManagementGetResellerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status/{id}"] = status_management.NewGetStatus(o.context, o.StatusManagementGetStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/customer"] = customer_management.NewListCustomers(o.context, o.CustomerManagementListCustomersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product"] = product_management.NewListProducts(o.context, o.ProductManagementListProductsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/reseller"] = reseller_management.NewListResellers(o.context, o.ResellerManagementListResellersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = status_management.NewShowStatus(o.context, o.StatusManagementShowStatusHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/customer/{id}"] = customer_management.NewUpdateCustomer(o.context, o.CustomerManagementUpdateCustomerHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/product/{id}"] = product_management.NewUpdateProduct(o.context, o.ProductManagementUpdateProductHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/reseller/{id}"] = reseller_management.NewUpdateReseller(o.context, o.ResellerManagementUpdateResellerHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CustomerDatabaseManagementAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CustomerDatabaseManagementAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CustomerDatabaseManagementAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CustomerDatabaseManagementAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CustomerDatabaseManagementAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

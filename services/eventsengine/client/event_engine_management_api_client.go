// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/client/event_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/client/status_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/client/trigger_management"
	"github.com/Cyclops-Labs/cyclops-4-hpc.git/services/eventsengine/client/usage_management"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8000"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1.0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new event engine management API HTTP client.
func New(c Config) *EventEngineManagementAPI {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(EventEngineManagementAPI)
	cli.Transport = transport
	cli.EventManagement = event_management.New(transport, strfmt.Default, c.AuthInfo)
	cli.StatusManagement = status_management.New(transport, strfmt.Default, c.AuthInfo)
	cli.TriggerManagement = trigger_management.New(transport, strfmt.Default, c.AuthInfo)
	cli.UsageManagement = usage_management.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// EventEngineManagementAPI is a client for event engine management API
type EventEngineManagementAPI struct {
	EventManagement   *event_management.Client
	StatusManagement  *status_management.Client
	TriggerManagement *trigger_management.Client
	UsageManagement   *usage_management.Client
	Transport         runtime.ClientTransport
}
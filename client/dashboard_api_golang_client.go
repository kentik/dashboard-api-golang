// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kentik/dashboard-api-golang/client/administered"
	"github.com/kentik/dashboard-api-golang/client/appliance"
	"github.com/kentik/dashboard-api-golang/client/camera"
	"github.com/kentik/dashboard-api-golang/client/cellular_gateway"
	"github.com/kentik/dashboard-api-golang/client/devices"
	"github.com/kentik/dashboard-api-golang/client/insight"
	"github.com/kentik/dashboard-api-golang/client/licensing"
	"github.com/kentik/dashboard-api-golang/client/networks"
	"github.com/kentik/dashboard-api-golang/client/organizations"
	"github.com/kentik/dashboard-api-golang/client/sensor"
	"github.com/kentik/dashboard-api-golang/client/sm"
	"github.com/kentik/dashboard-api-golang/client/switch_operations"
	"github.com/kentik/dashboard-api-golang/client/wireless"
)

// Default dashboard API golang HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.meraki.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new dashboard API golang HTTP client.
func NewHTTPClient(formats strfmt.Registry) *DashboardAPIGolang {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new dashboard API golang HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *DashboardAPIGolang {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new dashboard API golang client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *DashboardAPIGolang {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(DashboardAPIGolang)
	cli.Transport = transport
	cli.Administered = administered.New(transport, formats)
	cli.Appliance = appliance.New(transport, formats)
	cli.Camera = camera.New(transport, formats)
	cli.CellularGateway = cellular_gateway.New(transport, formats)
	cli.Devices = devices.New(transport, formats)
	cli.Insight = insight.New(transport, formats)
	cli.Licensing = licensing.New(transport, formats)
	cli.Networks = networks.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Sensor = sensor.New(transport, formats)
	cli.Sm = sm.New(transport, formats)
	cli.SwitchOperations = switch_operations.New(transport, formats)
	cli.Wireless = wireless.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// DashboardAPIGolang is a client for dashboard API golang
type DashboardAPIGolang struct {
	Administered administered.ClientService

	Appliance appliance.ClientService

	Camera camera.ClientService

	CellularGateway cellular_gateway.ClientService

	Devices devices.ClientService

	Insight insight.ClientService

	Licensing licensing.ClientService

	Networks networks.ClientService

	Organizations organizations.ClientService

	Sensor sensor.ClientService

	Sm sm.ClientService

	SwitchOperations switch_operations.ClientService

	Wireless wireless.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *DashboardAPIGolang) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Administered.SetTransport(transport)
	c.Appliance.SetTransport(transport)
	c.Camera.SetTransport(transport)
	c.CellularGateway.SetTransport(transport)
	c.Devices.SetTransport(transport)
	c.Insight.SetTransport(transport)
	c.Licensing.SetTransport(transport)
	c.Networks.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Sensor.SetTransport(transport)
	c.Sm.SetTransport(transport)
	c.SwitchOperations.SetTransport(transport)
	c.Wireless.SetTransport(transport)
}

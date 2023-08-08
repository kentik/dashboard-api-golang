// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkWirelessSsidFirewallL7FirewallRulesParams creates a new GetNetworkWirelessSsidFirewallL7FirewallRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidFirewallL7FirewallRulesParams() *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL7FirewallRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithTimeout creates a new GetNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL7FirewallRulesParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithContext creates a new GetNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL7FirewallRulesParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithHTTPClient creates a new GetNetworkWirelessSsidFirewallL7FirewallRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidFirewallL7FirewallRulesParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	return &GetNetworkWirelessSsidFirewallL7FirewallRulesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessSsidFirewallL7FirewallRulesParams contains all the parameters to send to the API endpoint

	for the get network wireless ssid firewall l7 firewall rules operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidFirewallL7FirewallRulesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Number.

	   Number
	*/
	Number string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless ssid firewall l7 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithDefaults() *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid firewall l7 firewall rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WithNumber(number string) *GetNetworkWirelessSsidFirewallL7FirewallRulesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid firewall l7 firewall rules params
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidFirewallL7FirewallRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

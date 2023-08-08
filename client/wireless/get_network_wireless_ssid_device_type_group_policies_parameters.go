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

// NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParams creates a new GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParams() *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithTimeout creates a new GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithContext creates a new GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithHTTPClient creates a new GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidDeviceTypeGroupPoliciesParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	return &GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams contains all the parameters to send to the API endpoint

	for the get network wireless ssid device type group policies operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams struct {

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

// WithDefaults hydrates default values in the get network wireless ssid device type group policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithDefaults() *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid device type group policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WithNumber(number string) *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid device type group policies params
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidDeviceTypeGroupPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

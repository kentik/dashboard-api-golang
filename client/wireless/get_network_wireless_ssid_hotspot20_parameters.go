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

// NewGetNetworkWirelessSsidHotspot20Params creates a new GetNetworkWirelessSsidHotspot20Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidHotspot20Params() *GetNetworkWirelessSsidHotspot20Params {
	return &GetNetworkWirelessSsidHotspot20Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidHotspot20ParamsWithTimeout creates a new GetNetworkWirelessSsidHotspot20Params object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidHotspot20ParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidHotspot20Params {
	return &GetNetworkWirelessSsidHotspot20Params{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidHotspot20ParamsWithContext creates a new GetNetworkWirelessSsidHotspot20Params object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidHotspot20ParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidHotspot20Params {
	return &GetNetworkWirelessSsidHotspot20Params{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidHotspot20ParamsWithHTTPClient creates a new GetNetworkWirelessSsidHotspot20Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidHotspot20ParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidHotspot20Params {
	return &GetNetworkWirelessSsidHotspot20Params{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessSsidHotspot20Params contains all the parameters to send to the API endpoint

	for the get network wireless ssid hotspot20 operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidHotspot20Params struct {

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

// WithDefaults hydrates default values in the get network wireless ssid hotspot20 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidHotspot20Params) WithDefaults() *GetNetworkWirelessSsidHotspot20Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssid hotspot20 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidHotspot20Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidHotspot20Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) WithContext(ctx context.Context) *GetNetworkWirelessSsidHotspot20Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidHotspot20Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) WithNetworkID(networkID string) *GetNetworkWirelessSsidHotspot20Params {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) WithNumber(number string) *GetNetworkWirelessSsidHotspot20Params {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get network wireless ssid hotspot20 params
func (o *GetNetworkWirelessSsidHotspot20Params) SetNumber(number string) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidHotspot20Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

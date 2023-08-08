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

// NewUpdateNetworkWirelessSsidSplashSettingsParams creates a new UpdateNetworkWirelessSsidSplashSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessSsidSplashSettingsParams() *UpdateNetworkWirelessSsidSplashSettingsParams {
	return &UpdateNetworkWirelessSsidSplashSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessSsidSplashSettingsParamsWithTimeout creates a new UpdateNetworkWirelessSsidSplashSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessSsidSplashSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidSplashSettingsParams {
	return &UpdateNetworkWirelessSsidSplashSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessSsidSplashSettingsParamsWithContext creates a new UpdateNetworkWirelessSsidSplashSettingsParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessSsidSplashSettingsParamsWithContext(ctx context.Context) *UpdateNetworkWirelessSsidSplashSettingsParams {
	return &UpdateNetworkWirelessSsidSplashSettingsParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessSsidSplashSettingsParamsWithHTTPClient creates a new UpdateNetworkWirelessSsidSplashSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessSsidSplashSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidSplashSettingsParams {
	return &UpdateNetworkWirelessSsidSplashSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessSsidSplashSettingsParams contains all the parameters to send to the API endpoint

	for the update network wireless ssid splash settings operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessSsidSplashSettingsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Number.

	   Number
	*/
	Number string

	// UpdateNetworkWirelessSsidSplashSettings.
	UpdateNetworkWirelessSsidSplashSettings UpdateNetworkWirelessSsidSplashSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless ssid splash settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithDefaults() *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless ssid splash settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithContext(ctx context.Context) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithNetworkID(networkID string) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithNumber adds the number to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithNumber(number string) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetNumber(number string) {
	o.Number = number
}

// WithUpdateNetworkWirelessSsidSplashSettings adds the updateNetworkWirelessSsidSplashSettings to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WithUpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings UpdateNetworkWirelessSsidSplashSettingsBody) *UpdateNetworkWirelessSsidSplashSettingsParams {
	o.SetUpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings)
	return o
}

// SetUpdateNetworkWirelessSsidSplashSettings adds the updateNetworkWirelessSsidSplashSettings to the update network wireless ssid splash settings params
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) SetUpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings UpdateNetworkWirelessSsidSplashSettingsBody) {
	o.UpdateNetworkWirelessSsidSplashSettings = updateNetworkWirelessSsidSplashSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessSsidSplashSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkWirelessSsidSplashSettings); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

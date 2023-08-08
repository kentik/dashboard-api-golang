// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewUpdateNetworkAlertsSettingsParams creates a new UpdateNetworkAlertsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkAlertsSettingsParams() *UpdateNetworkAlertsSettingsParams {
	return &UpdateNetworkAlertsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkAlertsSettingsParamsWithTimeout creates a new UpdateNetworkAlertsSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkAlertsSettingsParamsWithTimeout(timeout time.Duration) *UpdateNetworkAlertsSettingsParams {
	return &UpdateNetworkAlertsSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkAlertsSettingsParamsWithContext creates a new UpdateNetworkAlertsSettingsParams object
// with the ability to set a context for a request.
func NewUpdateNetworkAlertsSettingsParamsWithContext(ctx context.Context) *UpdateNetworkAlertsSettingsParams {
	return &UpdateNetworkAlertsSettingsParams{
		Context: ctx,
	}
}

// NewUpdateNetworkAlertsSettingsParamsWithHTTPClient creates a new UpdateNetworkAlertsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkAlertsSettingsParamsWithHTTPClient(client *http.Client) *UpdateNetworkAlertsSettingsParams {
	return &UpdateNetworkAlertsSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkAlertsSettingsParams contains all the parameters to send to the API endpoint

	for the update network alerts settings operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkAlertsSettingsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkAlertsSettings.
	UpdateNetworkAlertsSettings UpdateNetworkAlertsSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network alerts settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkAlertsSettingsParams) WithDefaults() *UpdateNetworkAlertsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network alerts settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkAlertsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) WithTimeout(timeout time.Duration) *UpdateNetworkAlertsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) WithContext(ctx context.Context) *UpdateNetworkAlertsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) WithHTTPClient(client *http.Client) *UpdateNetworkAlertsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) WithNetworkID(networkID string) *UpdateNetworkAlertsSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkAlertsSettings adds the updateNetworkAlertsSettings to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) WithUpdateNetworkAlertsSettings(updateNetworkAlertsSettings UpdateNetworkAlertsSettingsBody) *UpdateNetworkAlertsSettingsParams {
	o.SetUpdateNetworkAlertsSettings(updateNetworkAlertsSettings)
	return o
}

// SetUpdateNetworkAlertsSettings adds the updateNetworkAlertsSettings to the update network alerts settings params
func (o *UpdateNetworkAlertsSettingsParams) SetUpdateNetworkAlertsSettings(updateNetworkAlertsSettings UpdateNetworkAlertsSettingsBody) {
	o.UpdateNetworkAlertsSettings = updateNetworkAlertsSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkAlertsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkAlertsSettings); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

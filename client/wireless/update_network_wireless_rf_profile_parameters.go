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

// NewUpdateNetworkWirelessRfProfileParams creates a new UpdateNetworkWirelessRfProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkWirelessRfProfileParams() *UpdateNetworkWirelessRfProfileParams {
	return &UpdateNetworkWirelessRfProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkWirelessRfProfileParamsWithTimeout creates a new UpdateNetworkWirelessRfProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkWirelessRfProfileParamsWithTimeout(timeout time.Duration) *UpdateNetworkWirelessRfProfileParams {
	return &UpdateNetworkWirelessRfProfileParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkWirelessRfProfileParamsWithContext creates a new UpdateNetworkWirelessRfProfileParams object
// with the ability to set a context for a request.
func NewUpdateNetworkWirelessRfProfileParamsWithContext(ctx context.Context) *UpdateNetworkWirelessRfProfileParams {
	return &UpdateNetworkWirelessRfProfileParams{
		Context: ctx,
	}
}

// NewUpdateNetworkWirelessRfProfileParamsWithHTTPClient creates a new UpdateNetworkWirelessRfProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkWirelessRfProfileParamsWithHTTPClient(client *http.Client) *UpdateNetworkWirelessRfProfileParams {
	return &UpdateNetworkWirelessRfProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkWirelessRfProfileParams contains all the parameters to send to the API endpoint

	for the update network wireless rf profile operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkWirelessRfProfileParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RfProfileID.

	   Rf profile ID
	*/
	RfProfileID string

	// UpdateNetworkWirelessRfProfile.
	UpdateNetworkWirelessRfProfile UpdateNetworkWirelessRfProfileBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessRfProfileParams) WithDefaults() *UpdateNetworkWirelessRfProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkWirelessRfProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithTimeout(timeout time.Duration) *UpdateNetworkWirelessRfProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithContext(ctx context.Context) *UpdateNetworkWirelessRfProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithHTTPClient(client *http.Client) *UpdateNetworkWirelessRfProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithNetworkID(networkID string) *UpdateNetworkWirelessRfProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRfProfileID adds the rfProfileID to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithRfProfileID(rfProfileID string) *UpdateNetworkWirelessRfProfileParams {
	o.SetRfProfileID(rfProfileID)
	return o
}

// SetRfProfileID adds the rfProfileId to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetRfProfileID(rfProfileID string) {
	o.RfProfileID = rfProfileID
}

// WithUpdateNetworkWirelessRfProfile adds the updateNetworkWirelessRfProfile to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) WithUpdateNetworkWirelessRfProfile(updateNetworkWirelessRfProfile UpdateNetworkWirelessRfProfileBody) *UpdateNetworkWirelessRfProfileParams {
	o.SetUpdateNetworkWirelessRfProfile(updateNetworkWirelessRfProfile)
	return o
}

// SetUpdateNetworkWirelessRfProfile adds the updateNetworkWirelessRfProfile to the update network wireless rf profile params
func (o *UpdateNetworkWirelessRfProfileParams) SetUpdateNetworkWirelessRfProfile(updateNetworkWirelessRfProfile UpdateNetworkWirelessRfProfileBody) {
	o.UpdateNetworkWirelessRfProfile = updateNetworkWirelessRfProfile
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkWirelessRfProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param rfProfileId
	if err := r.SetPathParam("rfProfileId", o.RfProfileID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkWirelessRfProfile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

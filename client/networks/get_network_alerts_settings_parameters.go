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

// NewGetNetworkAlertsSettingsParams creates a new GetNetworkAlertsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkAlertsSettingsParams() *GetNetworkAlertsSettingsParams {
	return &GetNetworkAlertsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAlertsSettingsParamsWithTimeout creates a new GetNetworkAlertsSettingsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkAlertsSettingsParamsWithTimeout(timeout time.Duration) *GetNetworkAlertsSettingsParams {
	return &GetNetworkAlertsSettingsParams{
		timeout: timeout,
	}
}

// NewGetNetworkAlertsSettingsParamsWithContext creates a new GetNetworkAlertsSettingsParams object
// with the ability to set a context for a request.
func NewGetNetworkAlertsSettingsParamsWithContext(ctx context.Context) *GetNetworkAlertsSettingsParams {
	return &GetNetworkAlertsSettingsParams{
		Context: ctx,
	}
}

// NewGetNetworkAlertsSettingsParamsWithHTTPClient creates a new GetNetworkAlertsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkAlertsSettingsParamsWithHTTPClient(client *http.Client) *GetNetworkAlertsSettingsParams {
	return &GetNetworkAlertsSettingsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkAlertsSettingsParams contains all the parameters to send to the API endpoint

	for the get network alerts settings operation.

	Typically these are written to a http.Request.
*/
type GetNetworkAlertsSettingsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network alerts settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkAlertsSettingsParams) WithDefaults() *GetNetworkAlertsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network alerts settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkAlertsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) WithTimeout(timeout time.Duration) *GetNetworkAlertsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) WithContext(ctx context.Context) *GetNetworkAlertsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) WithHTTPClient(client *http.Client) *GetNetworkAlertsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) WithNetworkID(networkID string) *GetNetworkAlertsSettingsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network alerts settings params
func (o *GetNetworkAlertsSettingsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAlertsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

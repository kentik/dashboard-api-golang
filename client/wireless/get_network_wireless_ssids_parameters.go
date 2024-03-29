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

// NewGetNetworkWirelessSsidsParams creates a new GetNetworkWirelessSsidsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessSsidsParams() *GetNetworkWirelessSsidsParams {
	return &GetNetworkWirelessSsidsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessSsidsParamsWithTimeout creates a new GetNetworkWirelessSsidsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessSsidsParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessSsidsParams {
	return &GetNetworkWirelessSsidsParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessSsidsParamsWithContext creates a new GetNetworkWirelessSsidsParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessSsidsParamsWithContext(ctx context.Context) *GetNetworkWirelessSsidsParams {
	return &GetNetworkWirelessSsidsParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessSsidsParamsWithHTTPClient creates a new GetNetworkWirelessSsidsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessSsidsParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessSsidsParams {
	return &GetNetworkWirelessSsidsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessSsidsParams contains all the parameters to send to the API endpoint

	for the get network wireless ssids operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessSsidsParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless ssids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidsParams) WithDefaults() *GetNetworkWirelessSsidsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless ssids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessSsidsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessSsidsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) WithContext(ctx context.Context) *GetNetworkWirelessSsidsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessSsidsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) WithNetworkID(networkID string) *GetNetworkWirelessSsidsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless ssids params
func (o *GetNetworkWirelessSsidsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessSsidsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

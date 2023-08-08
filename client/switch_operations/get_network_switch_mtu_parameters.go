// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// NewGetNetworkSwitchMtuParams creates a new GetNetworkSwitchMtuParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchMtuParams() *GetNetworkSwitchMtuParams {
	return &GetNetworkSwitchMtuParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchMtuParamsWithTimeout creates a new GetNetworkSwitchMtuParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchMtuParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchMtuParams {
	return &GetNetworkSwitchMtuParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchMtuParamsWithContext creates a new GetNetworkSwitchMtuParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchMtuParamsWithContext(ctx context.Context) *GetNetworkSwitchMtuParams {
	return &GetNetworkSwitchMtuParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchMtuParamsWithHTTPClient creates a new GetNetworkSwitchMtuParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchMtuParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchMtuParams {
	return &GetNetworkSwitchMtuParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchMtuParams contains all the parameters to send to the API endpoint

	for the get network switch mtu operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchMtuParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch mtu params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchMtuParams) WithDefaults() *GetNetworkSwitchMtuParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch mtu params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchMtuParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchMtuParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) WithContext(ctx context.Context) *GetNetworkSwitchMtuParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchMtuParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) WithNetworkID(networkID string) *GetNetworkSwitchMtuParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch mtu params
func (o *GetNetworkSwitchMtuParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchMtuParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

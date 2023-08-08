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

// NewGetNetworkClientUsageHistoryParams creates a new GetNetworkClientUsageHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkClientUsageHistoryParams() *GetNetworkClientUsageHistoryParams {
	return &GetNetworkClientUsageHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkClientUsageHistoryParamsWithTimeout creates a new GetNetworkClientUsageHistoryParams object
// with the ability to set a timeout on a request.
func NewGetNetworkClientUsageHistoryParamsWithTimeout(timeout time.Duration) *GetNetworkClientUsageHistoryParams {
	return &GetNetworkClientUsageHistoryParams{
		timeout: timeout,
	}
}

// NewGetNetworkClientUsageHistoryParamsWithContext creates a new GetNetworkClientUsageHistoryParams object
// with the ability to set a context for a request.
func NewGetNetworkClientUsageHistoryParamsWithContext(ctx context.Context) *GetNetworkClientUsageHistoryParams {
	return &GetNetworkClientUsageHistoryParams{
		Context: ctx,
	}
}

// NewGetNetworkClientUsageHistoryParamsWithHTTPClient creates a new GetNetworkClientUsageHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkClientUsageHistoryParamsWithHTTPClient(client *http.Client) *GetNetworkClientUsageHistoryParams {
	return &GetNetworkClientUsageHistoryParams{
		HTTPClient: client,
	}
}

/*
GetNetworkClientUsageHistoryParams contains all the parameters to send to the API endpoint

	for the get network client usage history operation.

	Typically these are written to a http.Request.
*/
type GetNetworkClientUsageHistoryParams struct {

	/* ClientID.

	   Client ID
	*/
	ClientID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network client usage history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkClientUsageHistoryParams) WithDefaults() *GetNetworkClientUsageHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network client usage history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkClientUsageHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) WithTimeout(timeout time.Duration) *GetNetworkClientUsageHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) WithContext(ctx context.Context) *GetNetworkClientUsageHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) WithHTTPClient(client *http.Client) *GetNetworkClientUsageHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) WithClientID(clientID string) *GetNetworkClientUsageHistoryParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithNetworkID adds the networkID to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) WithNetworkID(networkID string) *GetNetworkClientUsageHistoryParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network client usage history params
func (o *GetNetworkClientUsageHistoryParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkClientUsageHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

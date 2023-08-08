// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// NewGetNetworkCellularGatewayDhcpParams creates a new GetNetworkCellularGatewayDhcpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCellularGatewayDhcpParams() *GetNetworkCellularGatewayDhcpParams {
	return &GetNetworkCellularGatewayDhcpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCellularGatewayDhcpParamsWithTimeout creates a new GetNetworkCellularGatewayDhcpParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCellularGatewayDhcpParamsWithTimeout(timeout time.Duration) *GetNetworkCellularGatewayDhcpParams {
	return &GetNetworkCellularGatewayDhcpParams{
		timeout: timeout,
	}
}

// NewGetNetworkCellularGatewayDhcpParamsWithContext creates a new GetNetworkCellularGatewayDhcpParams object
// with the ability to set a context for a request.
func NewGetNetworkCellularGatewayDhcpParamsWithContext(ctx context.Context) *GetNetworkCellularGatewayDhcpParams {
	return &GetNetworkCellularGatewayDhcpParams{
		Context: ctx,
	}
}

// NewGetNetworkCellularGatewayDhcpParamsWithHTTPClient creates a new GetNetworkCellularGatewayDhcpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCellularGatewayDhcpParamsWithHTTPClient(client *http.Client) *GetNetworkCellularGatewayDhcpParams {
	return &GetNetworkCellularGatewayDhcpParams{
		HTTPClient: client,
	}
}

/*
GetNetworkCellularGatewayDhcpParams contains all the parameters to send to the API endpoint

	for the get network cellular gateway dhcp operation.

	Typically these are written to a http.Request.
*/
type GetNetworkCellularGatewayDhcpParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network cellular gateway dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCellularGatewayDhcpParams) WithDefaults() *GetNetworkCellularGatewayDhcpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network cellular gateway dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCellularGatewayDhcpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) WithTimeout(timeout time.Duration) *GetNetworkCellularGatewayDhcpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) WithContext(ctx context.Context) *GetNetworkCellularGatewayDhcpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) WithHTTPClient(client *http.Client) *GetNetworkCellularGatewayDhcpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) WithNetworkID(networkID string) *GetNetworkCellularGatewayDhcpParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network cellular gateway dhcp params
func (o *GetNetworkCellularGatewayDhcpParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCellularGatewayDhcpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

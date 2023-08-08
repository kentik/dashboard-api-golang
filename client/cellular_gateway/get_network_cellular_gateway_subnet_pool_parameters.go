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

// NewGetNetworkCellularGatewaySubnetPoolParams creates a new GetNetworkCellularGatewaySubnetPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkCellularGatewaySubnetPoolParams() *GetNetworkCellularGatewaySubnetPoolParams {
	return &GetNetworkCellularGatewaySubnetPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkCellularGatewaySubnetPoolParamsWithTimeout creates a new GetNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a timeout on a request.
func NewGetNetworkCellularGatewaySubnetPoolParamsWithTimeout(timeout time.Duration) *GetNetworkCellularGatewaySubnetPoolParams {
	return &GetNetworkCellularGatewaySubnetPoolParams{
		timeout: timeout,
	}
}

// NewGetNetworkCellularGatewaySubnetPoolParamsWithContext creates a new GetNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a context for a request.
func NewGetNetworkCellularGatewaySubnetPoolParamsWithContext(ctx context.Context) *GetNetworkCellularGatewaySubnetPoolParams {
	return &GetNetworkCellularGatewaySubnetPoolParams{
		Context: ctx,
	}
}

// NewGetNetworkCellularGatewaySubnetPoolParamsWithHTTPClient creates a new GetNetworkCellularGatewaySubnetPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkCellularGatewaySubnetPoolParamsWithHTTPClient(client *http.Client) *GetNetworkCellularGatewaySubnetPoolParams {
	return &GetNetworkCellularGatewaySubnetPoolParams{
		HTTPClient: client,
	}
}

/*
GetNetworkCellularGatewaySubnetPoolParams contains all the parameters to send to the API endpoint

	for the get network cellular gateway subnet pool operation.

	Typically these are written to a http.Request.
*/
type GetNetworkCellularGatewaySubnetPoolParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network cellular gateway subnet pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCellularGatewaySubnetPoolParams) WithDefaults() *GetNetworkCellularGatewaySubnetPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network cellular gateway subnet pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkCellularGatewaySubnetPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) WithTimeout(timeout time.Duration) *GetNetworkCellularGatewaySubnetPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) WithContext(ctx context.Context) *GetNetworkCellularGatewaySubnetPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) WithHTTPClient(client *http.Client) *GetNetworkCellularGatewaySubnetPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) WithNetworkID(networkID string) *GetNetworkCellularGatewaySubnetPoolParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network cellular gateway subnet pool params
func (o *GetNetworkCellularGatewaySubnetPoolParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkCellularGatewaySubnetPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

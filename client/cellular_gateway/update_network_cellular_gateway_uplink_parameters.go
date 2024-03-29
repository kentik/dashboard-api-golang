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

// NewUpdateNetworkCellularGatewayUplinkParams creates a new UpdateNetworkCellularGatewayUplinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkCellularGatewayUplinkParams() *UpdateNetworkCellularGatewayUplinkParams {
	return &UpdateNetworkCellularGatewayUplinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkCellularGatewayUplinkParamsWithTimeout creates a new UpdateNetworkCellularGatewayUplinkParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkCellularGatewayUplinkParamsWithTimeout(timeout time.Duration) *UpdateNetworkCellularGatewayUplinkParams {
	return &UpdateNetworkCellularGatewayUplinkParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkCellularGatewayUplinkParamsWithContext creates a new UpdateNetworkCellularGatewayUplinkParams object
// with the ability to set a context for a request.
func NewUpdateNetworkCellularGatewayUplinkParamsWithContext(ctx context.Context) *UpdateNetworkCellularGatewayUplinkParams {
	return &UpdateNetworkCellularGatewayUplinkParams{
		Context: ctx,
	}
}

// NewUpdateNetworkCellularGatewayUplinkParamsWithHTTPClient creates a new UpdateNetworkCellularGatewayUplinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkCellularGatewayUplinkParamsWithHTTPClient(client *http.Client) *UpdateNetworkCellularGatewayUplinkParams {
	return &UpdateNetworkCellularGatewayUplinkParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkCellularGatewayUplinkParams contains all the parameters to send to the API endpoint

	for the update network cellular gateway uplink operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkCellularGatewayUplinkParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkCellularGatewayUplink.
	UpdateNetworkCellularGatewayUplink UpdateNetworkCellularGatewayUplinkBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network cellular gateway uplink params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCellularGatewayUplinkParams) WithDefaults() *UpdateNetworkCellularGatewayUplinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network cellular gateway uplink params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkCellularGatewayUplinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) WithTimeout(timeout time.Duration) *UpdateNetworkCellularGatewayUplinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) WithContext(ctx context.Context) *UpdateNetworkCellularGatewayUplinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) WithHTTPClient(client *http.Client) *UpdateNetworkCellularGatewayUplinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) WithNetworkID(networkID string) *UpdateNetworkCellularGatewayUplinkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkCellularGatewayUplink adds the updateNetworkCellularGatewayUplink to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) WithUpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink UpdateNetworkCellularGatewayUplinkBody) *UpdateNetworkCellularGatewayUplinkParams {
	o.SetUpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink)
	return o
}

// SetUpdateNetworkCellularGatewayUplink adds the updateNetworkCellularGatewayUplink to the update network cellular gateway uplink params
func (o *UpdateNetworkCellularGatewayUplinkParams) SetUpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink UpdateNetworkCellularGatewayUplinkBody) {
	o.UpdateNetworkCellularGatewayUplink = updateNetworkCellularGatewayUplink
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkCellularGatewayUplinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkCellularGatewayUplink); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

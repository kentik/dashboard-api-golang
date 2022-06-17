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

// NewGetNetworkSwitchStackRoutingInterfaceDhcpParams creates a new GetNetworkSwitchStackRoutingInterfaceDhcpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchStackRoutingInterfaceDhcpParams() *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	return &GetNetworkSwitchStackRoutingInterfaceDhcpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithTimeout creates a new GetNetworkSwitchStackRoutingInterfaceDhcpParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	return &GetNetworkSwitchStackRoutingInterfaceDhcpParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithContext creates a new GetNetworkSwitchStackRoutingInterfaceDhcpParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithContext(ctx context.Context) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	return &GetNetworkSwitchStackRoutingInterfaceDhcpParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithHTTPClient creates a new GetNetworkSwitchStackRoutingInterfaceDhcpParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchStackRoutingInterfaceDhcpParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	return &GetNetworkSwitchStackRoutingInterfaceDhcpParams{
		HTTPClient: client,
	}
}

/* GetNetworkSwitchStackRoutingInterfaceDhcpParams contains all the parameters to send to the API endpoint
   for the get network switch stack routing interface dhcp operation.

   Typically these are written to a http.Request.
*/
type GetNetworkSwitchStackRoutingInterfaceDhcpParams struct {

	// InterfaceID.
	InterfaceID string

	// NetworkID.
	NetworkID string

	// SwitchStackID.
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch stack routing interface dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithDefaults() *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch stack routing interface dhcp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithContext(ctx context.Context) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterfaceID adds the interfaceID to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithInterfaceID(interfaceID string) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetInterfaceID(interfaceID string) {
	o.InterfaceID = interfaceID
}

// WithNetworkID adds the networkID to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithNetworkID(networkID string) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSwitchStackID adds the switchStackID to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WithSwitchStackID(switchStackID string) *GetNetworkSwitchStackRoutingInterfaceDhcpParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the get network switch stack routing interface dhcp params
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchStackRoutingInterfaceDhcpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interfaceId
	if err := r.SetPathParam("interfaceId", o.InterfaceID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param switchStackId
	if err := r.SetPathParam("switchStackId", o.SwitchStackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

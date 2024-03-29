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

// NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams creates a new DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams() *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	return &DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithTimeout creates a new DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithTimeout(timeout time.Duration) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	return &DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithContext creates a new DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams object
// with the ability to set a context for a request.
func NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithContext(ctx context.Context) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	return &DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams{
		Context: ctx,
	}
}

// NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithHTTPClient creates a new DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsWithHTTPClient(client *http.Client) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	return &DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams contains all the parameters to send to the API endpoint

	for the delete network switch dhcp server policy arp inspection trusted server operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* TrustedServerID.

	   Trusted server ID
	*/
	TrustedServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network switch dhcp server policy arp inspection trusted server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithDefaults() *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network switch dhcp server policy arp inspection trusted server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithTimeout(timeout time.Duration) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithContext(ctx context.Context) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithHTTPClient(client *http.Client) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithNetworkID(networkID string) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithTrustedServerID adds the trustedServerID to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WithTrustedServerID(trustedServerID string) *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams {
	o.SetTrustedServerID(trustedServerID)
	return o
}

// SetTrustedServerID adds the trustedServerId to the delete network switch dhcp server policy arp inspection trusted server params
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) SetTrustedServerID(trustedServerID string) {
	o.TrustedServerID = trustedServerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param trustedServerId
	if err := r.SetPathParam("trustedServerId", o.TrustedServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

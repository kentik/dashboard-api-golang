// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesParams creates a new UpdateNetworkApplianceFirewallOneToManyNatRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesParams() *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	return &UpdateNetworkApplianceFirewallOneToManyNatRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithTimeout creates a new UpdateNetworkApplianceFirewallOneToManyNatRulesParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	return &UpdateNetworkApplianceFirewallOneToManyNatRulesParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithContext creates a new UpdateNetworkApplianceFirewallOneToManyNatRulesParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithContext(ctx context.Context) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	return &UpdateNetworkApplianceFirewallOneToManyNatRulesParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithHTTPClient creates a new UpdateNetworkApplianceFirewallOneToManyNatRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	return &UpdateNetworkApplianceFirewallOneToManyNatRulesParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceFirewallOneToManyNatRulesParams contains all the parameters to send to the API endpoint

	for the update network appliance firewall one to many nat rules operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceFirewallOneToManyNatRulesParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkApplianceFirewallOneToManyNatRules.
	UpdateNetworkApplianceFirewallOneToManyNatRules UpdateNetworkApplianceFirewallOneToManyNatRulesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance firewall one to many nat rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithDefaults() *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance firewall one to many nat rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithContext(ctx context.Context) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithNetworkID(networkID string) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceFirewallOneToManyNatRules adds the updateNetworkApplianceFirewallOneToManyNatRules to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WithUpdateNetworkApplianceFirewallOneToManyNatRules(updateNetworkApplianceFirewallOneToManyNatRules UpdateNetworkApplianceFirewallOneToManyNatRulesBody) *UpdateNetworkApplianceFirewallOneToManyNatRulesParams {
	o.SetUpdateNetworkApplianceFirewallOneToManyNatRules(updateNetworkApplianceFirewallOneToManyNatRules)
	return o
}

// SetUpdateNetworkApplianceFirewallOneToManyNatRules adds the updateNetworkApplianceFirewallOneToManyNatRules to the update network appliance firewall one to many nat rules params
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) SetUpdateNetworkApplianceFirewallOneToManyNatRules(updateNetworkApplianceFirewallOneToManyNatRules UpdateNetworkApplianceFirewallOneToManyNatRulesBody) {
	o.UpdateNetworkApplianceFirewallOneToManyNatRules = updateNetworkApplianceFirewallOneToManyNatRules
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceFirewallOneToManyNatRules); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

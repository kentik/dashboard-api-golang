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

// NewUpdateNetworkSwitchStackRoutingInterfaceParams creates a new UpdateNetworkSwitchStackRoutingInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSwitchStackRoutingInterfaceParams() *UpdateNetworkSwitchStackRoutingInterfaceParams {
	return &UpdateNetworkSwitchStackRoutingInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithTimeout creates a new UpdateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	return &UpdateNetworkSwitchStackRoutingInterfaceParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithContext creates a new UpdateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithContext(ctx context.Context) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	return &UpdateNetworkSwitchStackRoutingInterfaceParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithHTTPClient creates a new UpdateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSwitchStackRoutingInterfaceParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	return &UpdateNetworkSwitchStackRoutingInterfaceParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkSwitchStackRoutingInterfaceParams contains all the parameters to send to the API endpoint

	for the update network switch stack routing interface operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkSwitchStackRoutingInterfaceParams struct {

	/* InterfaceID.

	   Interface ID
	*/
	InterfaceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* SwitchStackID.

	   Switch stack ID
	*/
	SwitchStackID string

	// UpdateNetworkSwitchStackRoutingInterface.
	UpdateNetworkSwitchStackRoutingInterface UpdateNetworkSwitchStackRoutingInterfaceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network switch stack routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithDefaults() *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network switch stack routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithContext(ctx context.Context) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterfaceID adds the interfaceID to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithInterfaceID(interfaceID string) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetInterfaceID(interfaceID string) {
	o.InterfaceID = interfaceID
}

// WithNetworkID adds the networkID to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithNetworkID(networkID string) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSwitchStackID adds the switchStackID to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithSwitchStackID(switchStackID string) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WithUpdateNetworkSwitchStackRoutingInterface adds the updateNetworkSwitchStackRoutingInterface to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WithUpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface UpdateNetworkSwitchStackRoutingInterfaceBody) *UpdateNetworkSwitchStackRoutingInterfaceParams {
	o.SetUpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface)
	return o
}

// SetUpdateNetworkSwitchStackRoutingInterface adds the updateNetworkSwitchStackRoutingInterface to the update network switch stack routing interface params
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) SetUpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface UpdateNetworkSwitchStackRoutingInterfaceBody) {
	o.UpdateNetworkSwitchStackRoutingInterface = updateNetworkSwitchStackRoutingInterface
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchStackRoutingInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkSwitchStackRoutingInterface); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

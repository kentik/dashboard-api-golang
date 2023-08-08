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

// NewCreateNetworkSwitchStackRoutingInterfaceParams creates a new CreateNetworkSwitchStackRoutingInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSwitchStackRoutingInterfaceParams() *CreateNetworkSwitchStackRoutingInterfaceParams {
	return &CreateNetworkSwitchStackRoutingInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSwitchStackRoutingInterfaceParamsWithTimeout creates a new CreateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSwitchStackRoutingInterfaceParamsWithTimeout(timeout time.Duration) *CreateNetworkSwitchStackRoutingInterfaceParams {
	return &CreateNetworkSwitchStackRoutingInterfaceParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSwitchStackRoutingInterfaceParamsWithContext creates a new CreateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a context for a request.
func NewCreateNetworkSwitchStackRoutingInterfaceParamsWithContext(ctx context.Context) *CreateNetworkSwitchStackRoutingInterfaceParams {
	return &CreateNetworkSwitchStackRoutingInterfaceParams{
		Context: ctx,
	}
}

// NewCreateNetworkSwitchStackRoutingInterfaceParamsWithHTTPClient creates a new CreateNetworkSwitchStackRoutingInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSwitchStackRoutingInterfaceParamsWithHTTPClient(client *http.Client) *CreateNetworkSwitchStackRoutingInterfaceParams {
	return &CreateNetworkSwitchStackRoutingInterfaceParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkSwitchStackRoutingInterfaceParams contains all the parameters to send to the API endpoint

	for the create network switch stack routing interface operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkSwitchStackRoutingInterfaceParams struct {

	// CreateNetworkSwitchStackRoutingInterface.
	CreateNetworkSwitchStackRoutingInterface CreateNetworkSwitchStackRoutingInterfaceBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* SwitchStackID.

	   Switch stack ID
	*/
	SwitchStackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network switch stack routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithDefaults() *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network switch stack routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithTimeout(timeout time.Duration) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithContext(ctx context.Context) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithHTTPClient(client *http.Client) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSwitchStackRoutingInterface adds the createNetworkSwitchStackRoutingInterface to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithCreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface CreateNetworkSwitchStackRoutingInterfaceBody) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetCreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface)
	return o
}

// SetCreateNetworkSwitchStackRoutingInterface adds the createNetworkSwitchStackRoutingInterface to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetCreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface CreateNetworkSwitchStackRoutingInterfaceBody) {
	o.CreateNetworkSwitchStackRoutingInterface = createNetworkSwitchStackRoutingInterface
}

// WithNetworkID adds the networkID to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithNetworkID(networkID string) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSwitchStackID adds the switchStackID to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WithSwitchStackID(switchStackID string) *CreateNetworkSwitchStackRoutingInterfaceParams {
	o.SetSwitchStackID(switchStackID)
	return o
}

// SetSwitchStackID adds the switchStackId to the create network switch stack routing interface params
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) SetSwitchStackID(switchStackID string) {
	o.SwitchStackID = switchStackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSwitchStackRoutingInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSwitchStackRoutingInterface); err != nil {
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

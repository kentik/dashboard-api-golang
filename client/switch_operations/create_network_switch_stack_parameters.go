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

// NewCreateNetworkSwitchStackParams creates a new CreateNetworkSwitchStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSwitchStackParams() *CreateNetworkSwitchStackParams {
	return &CreateNetworkSwitchStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSwitchStackParamsWithTimeout creates a new CreateNetworkSwitchStackParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSwitchStackParamsWithTimeout(timeout time.Duration) *CreateNetworkSwitchStackParams {
	return &CreateNetworkSwitchStackParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSwitchStackParamsWithContext creates a new CreateNetworkSwitchStackParams object
// with the ability to set a context for a request.
func NewCreateNetworkSwitchStackParamsWithContext(ctx context.Context) *CreateNetworkSwitchStackParams {
	return &CreateNetworkSwitchStackParams{
		Context: ctx,
	}
}

// NewCreateNetworkSwitchStackParamsWithHTTPClient creates a new CreateNetworkSwitchStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSwitchStackParamsWithHTTPClient(client *http.Client) *CreateNetworkSwitchStackParams {
	return &CreateNetworkSwitchStackParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkSwitchStackParams contains all the parameters to send to the API endpoint

	for the create network switch stack operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkSwitchStackParams struct {

	// CreateNetworkSwitchStack.
	CreateNetworkSwitchStack CreateNetworkSwitchStackBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchStackParams) WithDefaults() *CreateNetworkSwitchStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network switch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSwitchStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) WithTimeout(timeout time.Duration) *CreateNetworkSwitchStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) WithContext(ctx context.Context) *CreateNetworkSwitchStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) WithHTTPClient(client *http.Client) *CreateNetworkSwitchStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSwitchStack adds the createNetworkSwitchStack to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) WithCreateNetworkSwitchStack(createNetworkSwitchStack CreateNetworkSwitchStackBody) *CreateNetworkSwitchStackParams {
	o.SetCreateNetworkSwitchStack(createNetworkSwitchStack)
	return o
}

// SetCreateNetworkSwitchStack adds the createNetworkSwitchStack to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) SetCreateNetworkSwitchStack(createNetworkSwitchStack CreateNetworkSwitchStackBody) {
	o.CreateNetworkSwitchStack = createNetworkSwitchStack
}

// WithNetworkID adds the networkID to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) WithNetworkID(networkID string) *CreateNetworkSwitchStackParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network switch stack params
func (o *CreateNetworkSwitchStackParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSwitchStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSwitchStack); err != nil {
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

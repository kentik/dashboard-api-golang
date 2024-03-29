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

// NewUpdateNetworkClientPolicyParams creates a new UpdateNetworkClientPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkClientPolicyParams() *UpdateNetworkClientPolicyParams {
	return &UpdateNetworkClientPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkClientPolicyParamsWithTimeout creates a new UpdateNetworkClientPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkClientPolicyParamsWithTimeout(timeout time.Duration) *UpdateNetworkClientPolicyParams {
	return &UpdateNetworkClientPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkClientPolicyParamsWithContext creates a new UpdateNetworkClientPolicyParams object
// with the ability to set a context for a request.
func NewUpdateNetworkClientPolicyParamsWithContext(ctx context.Context) *UpdateNetworkClientPolicyParams {
	return &UpdateNetworkClientPolicyParams{
		Context: ctx,
	}
}

// NewUpdateNetworkClientPolicyParamsWithHTTPClient creates a new UpdateNetworkClientPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkClientPolicyParamsWithHTTPClient(client *http.Client) *UpdateNetworkClientPolicyParams {
	return &UpdateNetworkClientPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkClientPolicyParams contains all the parameters to send to the API endpoint

	for the update network client policy operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkClientPolicyParams struct {

	/* ClientID.

	   Client ID
	*/
	ClientID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkClientPolicy.
	UpdateNetworkClientPolicy UpdateNetworkClientPolicyBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network client policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkClientPolicyParams) WithDefaults() *UpdateNetworkClientPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network client policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkClientPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithTimeout(timeout time.Duration) *UpdateNetworkClientPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithContext(ctx context.Context) *UpdateNetworkClientPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithHTTPClient(client *http.Client) *UpdateNetworkClientPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithClientID(clientID string) *UpdateNetworkClientPolicyParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithNetworkID adds the networkID to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithNetworkID(networkID string) *UpdateNetworkClientPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkClientPolicy adds the updateNetworkClientPolicy to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) WithUpdateNetworkClientPolicy(updateNetworkClientPolicy UpdateNetworkClientPolicyBody) *UpdateNetworkClientPolicyParams {
	o.SetUpdateNetworkClientPolicy(updateNetworkClientPolicy)
	return o
}

// SetUpdateNetworkClientPolicy adds the updateNetworkClientPolicy to the update network client policy params
func (o *UpdateNetworkClientPolicyParams) SetUpdateNetworkClientPolicy(updateNetworkClientPolicy UpdateNetworkClientPolicyBody) {
	o.UpdateNetworkClientPolicy = updateNetworkClientPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkClientPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.UpdateNetworkClientPolicy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewGetNetworkSwitchAccessPolicyParams creates a new GetNetworkSwitchAccessPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSwitchAccessPolicyParams() *GetNetworkSwitchAccessPolicyParams {
	return &GetNetworkSwitchAccessPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSwitchAccessPolicyParamsWithTimeout creates a new GetNetworkSwitchAccessPolicyParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSwitchAccessPolicyParamsWithTimeout(timeout time.Duration) *GetNetworkSwitchAccessPolicyParams {
	return &GetNetworkSwitchAccessPolicyParams{
		timeout: timeout,
	}
}

// NewGetNetworkSwitchAccessPolicyParamsWithContext creates a new GetNetworkSwitchAccessPolicyParams object
// with the ability to set a context for a request.
func NewGetNetworkSwitchAccessPolicyParamsWithContext(ctx context.Context) *GetNetworkSwitchAccessPolicyParams {
	return &GetNetworkSwitchAccessPolicyParams{
		Context: ctx,
	}
}

// NewGetNetworkSwitchAccessPolicyParamsWithHTTPClient creates a new GetNetworkSwitchAccessPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSwitchAccessPolicyParamsWithHTTPClient(client *http.Client) *GetNetworkSwitchAccessPolicyParams {
	return &GetNetworkSwitchAccessPolicyParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSwitchAccessPolicyParams contains all the parameters to send to the API endpoint

	for the get network switch access policy operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSwitchAccessPolicyParams struct {

	/* AccessPolicyNumber.

	   Access policy number
	*/
	AccessPolicyNumber string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network switch access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchAccessPolicyParams) WithDefaults() *GetNetworkSwitchAccessPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network switch access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSwitchAccessPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) WithTimeout(timeout time.Duration) *GetNetworkSwitchAccessPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) WithContext(ctx context.Context) *GetNetworkSwitchAccessPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) WithHTTPClient(client *http.Client) *GetNetworkSwitchAccessPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessPolicyNumber adds the accessPolicyNumber to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) WithAccessPolicyNumber(accessPolicyNumber string) *GetNetworkSwitchAccessPolicyParams {
	o.SetAccessPolicyNumber(accessPolicyNumber)
	return o
}

// SetAccessPolicyNumber adds the accessPolicyNumber to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) SetAccessPolicyNumber(accessPolicyNumber string) {
	o.AccessPolicyNumber = accessPolicyNumber
}

// WithNetworkID adds the networkID to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) WithNetworkID(networkID string) *GetNetworkSwitchAccessPolicyParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network switch access policy params
func (o *GetNetworkSwitchAccessPolicyParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSwitchAccessPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accessPolicyNumber
	if err := r.SetPathParam("accessPolicyNumber", o.AccessPolicyNumber); err != nil {
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

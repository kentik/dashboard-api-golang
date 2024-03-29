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

// NewUpdateNetworkApplianceSingleLanParams creates a new UpdateNetworkApplianceSingleLanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceSingleLanParams() *UpdateNetworkApplianceSingleLanParams {
	return &UpdateNetworkApplianceSingleLanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceSingleLanParamsWithTimeout creates a new UpdateNetworkApplianceSingleLanParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceSingleLanParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceSingleLanParams {
	return &UpdateNetworkApplianceSingleLanParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceSingleLanParamsWithContext creates a new UpdateNetworkApplianceSingleLanParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceSingleLanParamsWithContext(ctx context.Context) *UpdateNetworkApplianceSingleLanParams {
	return &UpdateNetworkApplianceSingleLanParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceSingleLanParamsWithHTTPClient creates a new UpdateNetworkApplianceSingleLanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceSingleLanParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceSingleLanParams {
	return &UpdateNetworkApplianceSingleLanParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceSingleLanParams contains all the parameters to send to the API endpoint

	for the update network appliance single lan operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceSingleLanParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkApplianceSingleLan.
	UpdateNetworkApplianceSingleLan UpdateNetworkApplianceSingleLanBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance single lan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSingleLanParams) WithDefaults() *UpdateNetworkApplianceSingleLanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance single lan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceSingleLanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceSingleLanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) WithContext(ctx context.Context) *UpdateNetworkApplianceSingleLanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceSingleLanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) WithNetworkID(networkID string) *UpdateNetworkApplianceSingleLanParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceSingleLan adds the updateNetworkApplianceSingleLan to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) WithUpdateNetworkApplianceSingleLan(updateNetworkApplianceSingleLan UpdateNetworkApplianceSingleLanBody) *UpdateNetworkApplianceSingleLanParams {
	o.SetUpdateNetworkApplianceSingleLan(updateNetworkApplianceSingleLan)
	return o
}

// SetUpdateNetworkApplianceSingleLan adds the updateNetworkApplianceSingleLan to the update network appliance single lan params
func (o *UpdateNetworkApplianceSingleLanParams) SetUpdateNetworkApplianceSingleLan(updateNetworkApplianceSingleLan UpdateNetworkApplianceSingleLanBody) {
	o.UpdateNetworkApplianceSingleLan = updateNetworkApplianceSingleLan
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceSingleLanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceSingleLan); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

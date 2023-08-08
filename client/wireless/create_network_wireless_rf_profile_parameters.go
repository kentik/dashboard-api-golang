// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewCreateNetworkWirelessRfProfileParams creates a new CreateNetworkWirelessRfProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkWirelessRfProfileParams() *CreateNetworkWirelessRfProfileParams {
	return &CreateNetworkWirelessRfProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkWirelessRfProfileParamsWithTimeout creates a new CreateNetworkWirelessRfProfileParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkWirelessRfProfileParamsWithTimeout(timeout time.Duration) *CreateNetworkWirelessRfProfileParams {
	return &CreateNetworkWirelessRfProfileParams{
		timeout: timeout,
	}
}

// NewCreateNetworkWirelessRfProfileParamsWithContext creates a new CreateNetworkWirelessRfProfileParams object
// with the ability to set a context for a request.
func NewCreateNetworkWirelessRfProfileParamsWithContext(ctx context.Context) *CreateNetworkWirelessRfProfileParams {
	return &CreateNetworkWirelessRfProfileParams{
		Context: ctx,
	}
}

// NewCreateNetworkWirelessRfProfileParamsWithHTTPClient creates a new CreateNetworkWirelessRfProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkWirelessRfProfileParamsWithHTTPClient(client *http.Client) *CreateNetworkWirelessRfProfileParams {
	return &CreateNetworkWirelessRfProfileParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkWirelessRfProfileParams contains all the parameters to send to the API endpoint

	for the create network wireless rf profile operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkWirelessRfProfileParams struct {

	// CreateNetworkWirelessRfProfile.
	CreateNetworkWirelessRfProfile CreateNetworkWirelessRfProfileBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkWirelessRfProfileParams) WithDefaults() *CreateNetworkWirelessRfProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkWirelessRfProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) WithTimeout(timeout time.Duration) *CreateNetworkWirelessRfProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) WithContext(ctx context.Context) *CreateNetworkWirelessRfProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) WithHTTPClient(client *http.Client) *CreateNetworkWirelessRfProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkWirelessRfProfile adds the createNetworkWirelessRfProfile to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) WithCreateNetworkWirelessRfProfile(createNetworkWirelessRfProfile CreateNetworkWirelessRfProfileBody) *CreateNetworkWirelessRfProfileParams {
	o.SetCreateNetworkWirelessRfProfile(createNetworkWirelessRfProfile)
	return o
}

// SetCreateNetworkWirelessRfProfile adds the createNetworkWirelessRfProfile to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) SetCreateNetworkWirelessRfProfile(createNetworkWirelessRfProfile CreateNetworkWirelessRfProfileBody) {
	o.CreateNetworkWirelessRfProfile = createNetworkWirelessRfProfile
}

// WithNetworkID adds the networkID to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) WithNetworkID(networkID string) *CreateNetworkWirelessRfProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network wireless rf profile params
func (o *CreateNetworkWirelessRfProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkWirelessRfProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkWirelessRfProfile); err != nil {
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

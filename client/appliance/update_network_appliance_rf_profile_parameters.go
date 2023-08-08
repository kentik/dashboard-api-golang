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

// NewUpdateNetworkApplianceRfProfileParams creates a new UpdateNetworkApplianceRfProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceRfProfileParams() *UpdateNetworkApplianceRfProfileParams {
	return &UpdateNetworkApplianceRfProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceRfProfileParamsWithTimeout creates a new UpdateNetworkApplianceRfProfileParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceRfProfileParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceRfProfileParams {
	return &UpdateNetworkApplianceRfProfileParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceRfProfileParamsWithContext creates a new UpdateNetworkApplianceRfProfileParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceRfProfileParamsWithContext(ctx context.Context) *UpdateNetworkApplianceRfProfileParams {
	return &UpdateNetworkApplianceRfProfileParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceRfProfileParamsWithHTTPClient creates a new UpdateNetworkApplianceRfProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceRfProfileParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceRfProfileParams {
	return &UpdateNetworkApplianceRfProfileParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceRfProfileParams contains all the parameters to send to the API endpoint

	for the update network appliance rf profile operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceRfProfileParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RfProfileID.

	   Rf profile ID
	*/
	RfProfileID string

	// UpdateNetworkApplianceRfProfile.
	UpdateNetworkApplianceRfProfile UpdateNetworkApplianceRfProfileBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceRfProfileParams) WithDefaults() *UpdateNetworkApplianceRfProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceRfProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceRfProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithContext(ctx context.Context) *UpdateNetworkApplianceRfProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceRfProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithNetworkID(networkID string) *UpdateNetworkApplianceRfProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRfProfileID adds the rfProfileID to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithRfProfileID(rfProfileID string) *UpdateNetworkApplianceRfProfileParams {
	o.SetRfProfileID(rfProfileID)
	return o
}

// SetRfProfileID adds the rfProfileId to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetRfProfileID(rfProfileID string) {
	o.RfProfileID = rfProfileID
}

// WithUpdateNetworkApplianceRfProfile adds the updateNetworkApplianceRfProfile to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) WithUpdateNetworkApplianceRfProfile(updateNetworkApplianceRfProfile UpdateNetworkApplianceRfProfileBody) *UpdateNetworkApplianceRfProfileParams {
	o.SetUpdateNetworkApplianceRfProfile(updateNetworkApplianceRfProfile)
	return o
}

// SetUpdateNetworkApplianceRfProfile adds the updateNetworkApplianceRfProfile to the update network appliance rf profile params
func (o *UpdateNetworkApplianceRfProfileParams) SetUpdateNetworkApplianceRfProfile(updateNetworkApplianceRfProfile UpdateNetworkApplianceRfProfileBody) {
	o.UpdateNetworkApplianceRfProfile = updateNetworkApplianceRfProfile
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceRfProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param rfProfileId
	if err := r.SetPathParam("rfProfileId", o.RfProfileID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceRfProfile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

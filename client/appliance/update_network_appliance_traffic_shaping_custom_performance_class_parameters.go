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

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams creates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithTimeout creates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithTimeout(timeout time.Duration) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithContext creates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a context for a request.
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithContext(ctx context.Context) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		Context: ctx,
	}
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithHTTPClient creates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassParamsWithHTTPClient(client *http.Client) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	return &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams contains all the parameters to send to the API endpoint

	for the update network appliance traffic shaping custom performance class operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams struct {

	/* CustomPerformanceClassID.

	   Custom performance class ID
	*/
	CustomPerformanceClassID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkApplianceTrafficShapingCustomPerformanceClass.
	UpdateNetworkApplianceTrafficShapingCustomPerformanceClass UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network appliance traffic shaping custom performance class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithDefaults() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network appliance traffic shaping custom performance class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithTimeout(timeout time.Duration) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithContext(ctx context.Context) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithHTTPClient(client *http.Client) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomPerformanceClassID adds the customPerformanceClassID to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithCustomPerformanceClassID(customPerformanceClassID string) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetCustomPerformanceClassID(customPerformanceClassID)
	return o
}

// SetCustomPerformanceClassID adds the customPerformanceClassId to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetCustomPerformanceClassID(customPerformanceClassID string) {
	o.CustomPerformanceClassID = customPerformanceClassID
}

// WithNetworkID adds the networkID to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithNetworkID(networkID string) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkApplianceTrafficShapingCustomPerformanceClass adds the updateNetworkApplianceTrafficShapingCustomPerformanceClass to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WithUpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams {
	o.SetUpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass)
	return o
}

// SetUpdateNetworkApplianceTrafficShapingCustomPerformanceClass adds the updateNetworkApplianceTrafficShapingCustomPerformanceClass to the update network appliance traffic shaping custom performance class params
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) SetUpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) {
	o.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass = updateNetworkApplianceTrafficShapingCustomPerformanceClass
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customPerformanceClassId
	if err := r.SetPathParam("customPerformanceClassId", o.CustomPerformanceClassID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

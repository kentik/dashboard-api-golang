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

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParams creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParams() *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithTimeout creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams object
// with the ability to set a timeout on a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams{
		timeout: timeout,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithContext creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams object
// with the ability to set a context for a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams{
		Context: ctx,
	}
}

// NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithHTTPClient creates a new GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkApplianceTrafficShapingCustomPerformanceClassesParamsWithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	return &GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams{
		HTTPClient: client,
	}
}

/* GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams contains all the parameters to send to the API endpoint
   for the get network appliance traffic shaping custom performance classes operation.

   Typically these are written to a http.Request.
*/
type GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams struct {

	// NetworkID.
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network appliance traffic shaping custom performance classes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WithDefaults() *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network appliance traffic shaping custom performance classes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WithTimeout(timeout time.Duration) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WithContext(ctx context.Context) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WithHTTPClient(client *http.Client) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WithNetworkID(networkID string) *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network appliance traffic shaping custom performance classes params
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkApplianceTrafficShapingCustomPerformanceClassesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
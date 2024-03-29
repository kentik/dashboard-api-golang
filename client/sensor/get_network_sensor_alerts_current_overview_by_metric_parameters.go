// Code generated by go-swagger; DO NOT EDIT.

package sensor

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

// NewGetNetworkSensorAlertsCurrentOverviewByMetricParams creates a new GetNetworkSensorAlertsCurrentOverviewByMetricParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSensorAlertsCurrentOverviewByMetricParams() *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	return &GetNetworkSensorAlertsCurrentOverviewByMetricParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithTimeout creates a new GetNetworkSensorAlertsCurrentOverviewByMetricParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithTimeout(timeout time.Duration) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	return &GetNetworkSensorAlertsCurrentOverviewByMetricParams{
		timeout: timeout,
	}
}

// NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithContext creates a new GetNetworkSensorAlertsCurrentOverviewByMetricParams object
// with the ability to set a context for a request.
func NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithContext(ctx context.Context) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	return &GetNetworkSensorAlertsCurrentOverviewByMetricParams{
		Context: ctx,
	}
}

// NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithHTTPClient creates a new GetNetworkSensorAlertsCurrentOverviewByMetricParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSensorAlertsCurrentOverviewByMetricParamsWithHTTPClient(client *http.Client) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	return &GetNetworkSensorAlertsCurrentOverviewByMetricParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSensorAlertsCurrentOverviewByMetricParams contains all the parameters to send to the API endpoint

	for the get network sensor alerts current overview by metric operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSensorAlertsCurrentOverviewByMetricParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sensor alerts current overview by metric params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WithDefaults() *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sensor alerts current overview by metric params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WithTimeout(timeout time.Duration) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WithContext(ctx context.Context) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WithHTTPClient(client *http.Client) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WithNetworkID(networkID string) *GetNetworkSensorAlertsCurrentOverviewByMetricParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sensor alerts current overview by metric params
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSensorAlertsCurrentOverviewByMetricParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

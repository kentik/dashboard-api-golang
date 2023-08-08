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
	"github.com/go-openapi/swag"
)

// NewGetNetworkWirelessAirMarshalParams creates a new GetNetworkWirelessAirMarshalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessAirMarshalParams() *GetNetworkWirelessAirMarshalParams {
	return &GetNetworkWirelessAirMarshalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessAirMarshalParamsWithTimeout creates a new GetNetworkWirelessAirMarshalParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessAirMarshalParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessAirMarshalParams {
	return &GetNetworkWirelessAirMarshalParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessAirMarshalParamsWithContext creates a new GetNetworkWirelessAirMarshalParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessAirMarshalParamsWithContext(ctx context.Context) *GetNetworkWirelessAirMarshalParams {
	return &GetNetworkWirelessAirMarshalParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessAirMarshalParamsWithHTTPClient creates a new GetNetworkWirelessAirMarshalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessAirMarshalParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessAirMarshalParams {
	return &GetNetworkWirelessAirMarshalParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessAirMarshalParams contains all the parameters to send to the API endpoint

	for the get network wireless air marshal operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessAirMarshalParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless air marshal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessAirMarshalParams) WithDefaults() *GetNetworkWirelessAirMarshalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless air marshal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessAirMarshalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessAirMarshalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithContext(ctx context.Context) *GetNetworkWirelessAirMarshalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessAirMarshalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithNetworkID(networkID string) *GetNetworkWirelessAirMarshalParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithT0 adds the t0 to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithT0(t0 *string) *GetNetworkWirelessAirMarshalParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) WithTimespan(timespan *float32) *GetNetworkWirelessAirMarshalParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network wireless air marshal params
func (o *GetNetworkWirelessAirMarshalParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessAirMarshalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewGetNetworkWirelessClientsLatencyStatsParams creates a new GetNetworkWirelessClientsLatencyStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessClientsLatencyStatsParams() *GetNetworkWirelessClientsLatencyStatsParams {
	return &GetNetworkWirelessClientsLatencyStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessClientsLatencyStatsParamsWithTimeout creates a new GetNetworkWirelessClientsLatencyStatsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessClientsLatencyStatsParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessClientsLatencyStatsParams {
	return &GetNetworkWirelessClientsLatencyStatsParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessClientsLatencyStatsParamsWithContext creates a new GetNetworkWirelessClientsLatencyStatsParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessClientsLatencyStatsParamsWithContext(ctx context.Context) *GetNetworkWirelessClientsLatencyStatsParams {
	return &GetNetworkWirelessClientsLatencyStatsParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessClientsLatencyStatsParamsWithHTTPClient creates a new GetNetworkWirelessClientsLatencyStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessClientsLatencyStatsParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessClientsLatencyStatsParams {
	return &GetNetworkWirelessClientsLatencyStatsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessClientsLatencyStatsParams contains all the parameters to send to the API endpoint

	for the get network wireless clients latency stats operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessClientsLatencyStatsParams struct {

	/* ApTag.

	   Filter results by AP Tag
	*/
	ApTag *string

	/* Band.

	   Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.
	*/
	Band *string

	/* Fields.

	   Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
	*/
	Fields *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Ssid.

	   Filter results by SSID
	*/
	Ssid *int64

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.

	   Format: float
	*/
	Timespan *float32

	/* Vlan.

	   Filter results by VLAN
	*/
	Vlan *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless clients latency stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithDefaults() *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless clients latency stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithContext(ctx context.Context) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApTag adds the apTag to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithApTag(apTag *string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetApTag(apTag)
	return o
}

// SetApTag adds the apTag to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetApTag(apTag *string) {
	o.ApTag = apTag
}

// WithBand adds the band to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithBand(band *string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetBand(band)
	return o
}

// SetBand adds the band to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetBand(band *string) {
	o.Band = band
}

// WithFields adds the fields to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithFields(fields *string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithNetworkID adds the networkID to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithNetworkID(networkID string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithSsid adds the ssid to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithSsid(ssid *int64) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetSsid(ssid)
	return o
}

// SetSsid adds the ssid to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetSsid(ssid *int64) {
	o.Ssid = ssid
}

// WithT0 adds the t0 to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithT0(t0 *string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithT1(t1 *string) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithTimespan(timespan *float32) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WithVlan adds the vlan to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) WithVlan(vlan *int64) *GetNetworkWirelessClientsLatencyStatsParams {
	o.SetVlan(vlan)
	return o
}

// SetVlan adds the vlan to the get network wireless clients latency stats params
func (o *GetNetworkWirelessClientsLatencyStatsParams) SetVlan(vlan *int64) {
	o.Vlan = vlan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessClientsLatencyStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApTag != nil {

		// query param apTag
		var qrApTag string

		if o.ApTag != nil {
			qrApTag = *o.ApTag
		}
		qApTag := qrApTag
		if qApTag != "" {

			if err := r.SetQueryParam("apTag", qApTag); err != nil {
				return err
			}
		}
	}

	if o.Band != nil {

		// query param band
		var qrBand string

		if o.Band != nil {
			qrBand = *o.Band
		}
		qBand := qrBand
		if qBand != "" {

			if err := r.SetQueryParam("band", qBand); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.Ssid != nil {

		// query param ssid
		var qrSsid int64

		if o.Ssid != nil {
			qrSsid = *o.Ssid
		}
		qSsid := swag.FormatInt64(qrSsid)
		if qSsid != "" {

			if err := r.SetQueryParam("ssid", qSsid); err != nil {
				return err
			}
		}
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

	if o.T1 != nil {

		// query param t1
		var qrT1 string

		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {

			if err := r.SetQueryParam("t1", qT1); err != nil {
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

	if o.Vlan != nil {

		// query param vlan
		var qrVlan int64

		if o.Vlan != nil {
			qrVlan = *o.Vlan
		}
		qVlan := swag.FormatInt64(qrVlan)
		if qVlan != "" {

			if err := r.SetQueryParam("vlan", qVlan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

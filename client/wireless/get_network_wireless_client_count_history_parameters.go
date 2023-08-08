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

// NewGetNetworkWirelessClientCountHistoryParams creates a new GetNetworkWirelessClientCountHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkWirelessClientCountHistoryParams() *GetNetworkWirelessClientCountHistoryParams {
	return &GetNetworkWirelessClientCountHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkWirelessClientCountHistoryParamsWithTimeout creates a new GetNetworkWirelessClientCountHistoryParams object
// with the ability to set a timeout on a request.
func NewGetNetworkWirelessClientCountHistoryParamsWithTimeout(timeout time.Duration) *GetNetworkWirelessClientCountHistoryParams {
	return &GetNetworkWirelessClientCountHistoryParams{
		timeout: timeout,
	}
}

// NewGetNetworkWirelessClientCountHistoryParamsWithContext creates a new GetNetworkWirelessClientCountHistoryParams object
// with the ability to set a context for a request.
func NewGetNetworkWirelessClientCountHistoryParamsWithContext(ctx context.Context) *GetNetworkWirelessClientCountHistoryParams {
	return &GetNetworkWirelessClientCountHistoryParams{
		Context: ctx,
	}
}

// NewGetNetworkWirelessClientCountHistoryParamsWithHTTPClient creates a new GetNetworkWirelessClientCountHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkWirelessClientCountHistoryParamsWithHTTPClient(client *http.Client) *GetNetworkWirelessClientCountHistoryParams {
	return &GetNetworkWirelessClientCountHistoryParams{
		HTTPClient: client,
	}
}

/*
GetNetworkWirelessClientCountHistoryParams contains all the parameters to send to the API endpoint

	for the get network wireless client count history operation.

	Typically these are written to a http.Request.
*/
type GetNetworkWirelessClientCountHistoryParams struct {

	/* ApTag.

	   Filter results by AP tag.
	*/
	ApTag *string

	/* AutoResolution.

	   Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
	*/
	AutoResolution *bool

	/* Band.

	   Filter results by band (either '2.4', '5' or '6').
	*/
	Band *string

	/* ClientID.

	   Filter results by network client to return per-device client counts over time inner joined by the queried client's connection history.
	*/
	ClientID *string

	/* DeviceSerial.

	   Filter results by device.
	*/
	DeviceSerial *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* Resolution.

	   The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
	*/
	Resolution *int64

	/* Ssid.

	   Filter results by SSID number.
	*/
	Ssid *int64

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network wireless client count history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessClientCountHistoryParams) WithDefaults() *GetNetworkWirelessClientCountHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network wireless client count history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkWirelessClientCountHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithTimeout(timeout time.Duration) *GetNetworkWirelessClientCountHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithContext(ctx context.Context) *GetNetworkWirelessClientCountHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithHTTPClient(client *http.Client) *GetNetworkWirelessClientCountHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApTag adds the apTag to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithApTag(apTag *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetApTag(apTag)
	return o
}

// SetApTag adds the apTag to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetApTag(apTag *string) {
	o.ApTag = apTag
}

// WithAutoResolution adds the autoResolution to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithAutoResolution(autoResolution *bool) *GetNetworkWirelessClientCountHistoryParams {
	o.SetAutoResolution(autoResolution)
	return o
}

// SetAutoResolution adds the autoResolution to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetAutoResolution(autoResolution *bool) {
	o.AutoResolution = autoResolution
}

// WithBand adds the band to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithBand(band *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetBand(band)
	return o
}

// SetBand adds the band to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetBand(band *string) {
	o.Band = band
}

// WithClientID adds the clientID to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithClientID(clientID *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithDeviceSerial adds the deviceSerial to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithDeviceSerial(deviceSerial *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetDeviceSerial(deviceSerial)
	return o
}

// SetDeviceSerial adds the deviceSerial to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetDeviceSerial(deviceSerial *string) {
	o.DeviceSerial = deviceSerial
}

// WithNetworkID adds the networkID to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithNetworkID(networkID string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithResolution adds the resolution to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithResolution(resolution *int64) *GetNetworkWirelessClientCountHistoryParams {
	o.SetResolution(resolution)
	return o
}

// SetResolution adds the resolution to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetResolution(resolution *int64) {
	o.Resolution = resolution
}

// WithSsid adds the ssid to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithSsid(ssid *int64) *GetNetworkWirelessClientCountHistoryParams {
	o.SetSsid(ssid)
	return o
}

// SetSsid adds the ssid to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetSsid(ssid *int64) {
	o.Ssid = ssid
}

// WithT0 adds the t0 to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithT0(t0 *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithT1(t1 *string) *GetNetworkWirelessClientCountHistoryParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) WithTimespan(timespan *float32) *GetNetworkWirelessClientCountHistoryParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network wireless client count history params
func (o *GetNetworkWirelessClientCountHistoryParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkWirelessClientCountHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AutoResolution != nil {

		// query param autoResolution
		var qrAutoResolution bool

		if o.AutoResolution != nil {
			qrAutoResolution = *o.AutoResolution
		}
		qAutoResolution := swag.FormatBool(qrAutoResolution)
		if qAutoResolution != "" {

			if err := r.SetQueryParam("autoResolution", qAutoResolution); err != nil {
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

	if o.ClientID != nil {

		// query param clientId
		var qrClientID string

		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("clientId", qClientID); err != nil {
				return err
			}
		}
	}

	if o.DeviceSerial != nil {

		// query param deviceSerial
		var qrDeviceSerial string

		if o.DeviceSerial != nil {
			qrDeviceSerial = *o.DeviceSerial
		}
		qDeviceSerial := qrDeviceSerial
		if qDeviceSerial != "" {

			if err := r.SetQueryParam("deviceSerial", qDeviceSerial); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.Resolution != nil {

		// query param resolution
		var qrResolution int64

		if o.Resolution != nil {
			qrResolution = *o.Resolution
		}
		qResolution := swag.FormatInt64(qrResolution)
		if qResolution != "" {

			if err := r.SetQueryParam("resolution", qResolution); err != nil {
				return err
			}
		}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

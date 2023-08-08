// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationDevicesUplinksLossAndLatencyParams creates a new GetOrganizationDevicesUplinksLossAndLatencyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationDevicesUplinksLossAndLatencyParams() *GetOrganizationDevicesUplinksLossAndLatencyParams {
	return &GetOrganizationDevicesUplinksLossAndLatencyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithTimeout creates a new GetOrganizationDevicesUplinksLossAndLatencyParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithTimeout(timeout time.Duration) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	return &GetOrganizationDevicesUplinksLossAndLatencyParams{
		timeout: timeout,
	}
}

// NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithContext creates a new GetOrganizationDevicesUplinksLossAndLatencyParams object
// with the ability to set a context for a request.
func NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithContext(ctx context.Context) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	return &GetOrganizationDevicesUplinksLossAndLatencyParams{
		Context: ctx,
	}
}

// NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithHTTPClient creates a new GetOrganizationDevicesUplinksLossAndLatencyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationDevicesUplinksLossAndLatencyParamsWithHTTPClient(client *http.Client) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	return &GetOrganizationDevicesUplinksLossAndLatencyParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationDevicesUplinksLossAndLatencyParams contains all the parameters to send to the API endpoint

	for the get organization devices uplinks loss and latency operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationDevicesUplinksLossAndLatencyParams struct {

	/* IP.

	   Optional filter for a specific destination IP. Default will return all destination IPs.
	*/
	IP *string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 60 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes.

	   Format: float
	*/
	Timespan *float32

	/* Uplink.

	   Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks.
	*/
	Uplink *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization devices uplinks loss and latency params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithDefaults() *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization devices uplinks loss and latency params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithTimeout(timeout time.Duration) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithContext(ctx context.Context) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithHTTPClient(client *http.Client) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithIP(ip *string) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetIP(ip *string) {
	o.IP = ip
}

// WithOrganizationID adds the organizationID to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithOrganizationID(organizationID string) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithT0 adds the t0 to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithT0(t0 *string) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithT1(t1 *string) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithTimespan(timespan *float32) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WithUplink adds the uplink to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WithUplink(uplink *string) *GetOrganizationDevicesUplinksLossAndLatencyParams {
	o.SetUplink(uplink)
	return o
}

// SetUplink adds the uplink to the get organization devices uplinks loss and latency params
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) SetUplink(uplink *string) {
	o.Uplink = uplink
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationDevicesUplinksLossAndLatencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IP != nil {

		// query param ip
		var qrIP string

		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {

			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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

	if o.Uplink != nil {

		// query param uplink
		var qrUplink string

		if o.Uplink != nil {
			qrUplink = *o.Uplink
		}
		qUplink := qrUplink
		if qUplink != "" {

			if err := r.SetQueryParam("uplink", qUplink); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

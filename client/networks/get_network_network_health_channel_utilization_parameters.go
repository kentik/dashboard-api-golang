// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkNetworkHealthChannelUtilizationParams creates a new GetNetworkNetworkHealthChannelUtilizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkNetworkHealthChannelUtilizationParams() *GetNetworkNetworkHealthChannelUtilizationParams {
	return &GetNetworkNetworkHealthChannelUtilizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkNetworkHealthChannelUtilizationParamsWithTimeout creates a new GetNetworkNetworkHealthChannelUtilizationParams object
// with the ability to set a timeout on a request.
func NewGetNetworkNetworkHealthChannelUtilizationParamsWithTimeout(timeout time.Duration) *GetNetworkNetworkHealthChannelUtilizationParams {
	return &GetNetworkNetworkHealthChannelUtilizationParams{
		timeout: timeout,
	}
}

// NewGetNetworkNetworkHealthChannelUtilizationParamsWithContext creates a new GetNetworkNetworkHealthChannelUtilizationParams object
// with the ability to set a context for a request.
func NewGetNetworkNetworkHealthChannelUtilizationParamsWithContext(ctx context.Context) *GetNetworkNetworkHealthChannelUtilizationParams {
	return &GetNetworkNetworkHealthChannelUtilizationParams{
		Context: ctx,
	}
}

// NewGetNetworkNetworkHealthChannelUtilizationParamsWithHTTPClient creates a new GetNetworkNetworkHealthChannelUtilizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkNetworkHealthChannelUtilizationParamsWithHTTPClient(client *http.Client) *GetNetworkNetworkHealthChannelUtilizationParams {
	return &GetNetworkNetworkHealthChannelUtilizationParams{
		HTTPClient: client,
	}
}

/*
GetNetworkNetworkHealthChannelUtilizationParams contains all the parameters to send to the API endpoint

	for the get network network health channel utilization operation.

	Typically these are written to a http.Request.
*/
type GetNetworkNetworkHealthChannelUtilizationParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 100. Default is 10.
	*/
	PerPage *int64

	/* Resolution.

	   The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600.
	*/
	Resolution *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network network health channel utilization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithDefaults() *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network network health channel utilization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithTimeout(timeout time.Duration) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithContext(ctx context.Context) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithHTTPClient(client *http.Client) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithEndingBefore(endingBefore *string) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithNetworkID(networkID string) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithPerPage(perPage *int64) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithResolution adds the resolution to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithResolution(resolution *int64) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetResolution(resolution)
	return o
}

// SetResolution adds the resolution to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetResolution(resolution *int64) {
	o.Resolution = resolution
}

// WithStartingAfter adds the startingAfter to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithStartingAfter(startingAfter *string) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithT0(t0 *string) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithT1(t1 *string) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WithTimespan(timespan *float32) *GetNetworkNetworkHealthChannelUtilizationParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network network health channel utilization params
func (o *GetNetworkNetworkHealthChannelUtilizationParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkNetworkHealthChannelUtilizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
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

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
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

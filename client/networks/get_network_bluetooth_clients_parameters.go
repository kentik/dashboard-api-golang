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

// NewGetNetworkBluetoothClientsParams creates a new GetNetworkBluetoothClientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkBluetoothClientsParams() *GetNetworkBluetoothClientsParams {
	return &GetNetworkBluetoothClientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkBluetoothClientsParamsWithTimeout creates a new GetNetworkBluetoothClientsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkBluetoothClientsParamsWithTimeout(timeout time.Duration) *GetNetworkBluetoothClientsParams {
	return &GetNetworkBluetoothClientsParams{
		timeout: timeout,
	}
}

// NewGetNetworkBluetoothClientsParamsWithContext creates a new GetNetworkBluetoothClientsParams object
// with the ability to set a context for a request.
func NewGetNetworkBluetoothClientsParamsWithContext(ctx context.Context) *GetNetworkBluetoothClientsParams {
	return &GetNetworkBluetoothClientsParams{
		Context: ctx,
	}
}

// NewGetNetworkBluetoothClientsParamsWithHTTPClient creates a new GetNetworkBluetoothClientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkBluetoothClientsParamsWithHTTPClient(client *http.Client) *GetNetworkBluetoothClientsParams {
	return &GetNetworkBluetoothClientsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkBluetoothClientsParams contains all the parameters to send to the API endpoint

	for the get network bluetooth clients operation.

	Typically these are written to a http.Request.
*/
type GetNetworkBluetoothClientsParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* IncludeConnectivityHistory.

	   Include the connectivity history for this client
	*/
	IncludeConnectivityHistory *bool

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
	*/
	T0 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network bluetooth clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkBluetoothClientsParams) WithDefaults() *GetNetworkBluetoothClientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network bluetooth clients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkBluetoothClientsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithTimeout(timeout time.Duration) *GetNetworkBluetoothClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithContext(ctx context.Context) *GetNetworkBluetoothClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithHTTPClient(client *http.Client) *GetNetworkBluetoothClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithEndingBefore(endingBefore *string) *GetNetworkBluetoothClientsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithIncludeConnectivityHistory adds the includeConnectivityHistory to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithIncludeConnectivityHistory(includeConnectivityHistory *bool) *GetNetworkBluetoothClientsParams {
	o.SetIncludeConnectivityHistory(includeConnectivityHistory)
	return o
}

// SetIncludeConnectivityHistory adds the includeConnectivityHistory to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetIncludeConnectivityHistory(includeConnectivityHistory *bool) {
	o.IncludeConnectivityHistory = includeConnectivityHistory
}

// WithNetworkID adds the networkID to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithNetworkID(networkID string) *GetNetworkBluetoothClientsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithPerPage(perPage *int64) *GetNetworkBluetoothClientsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithStartingAfter(startingAfter *string) *GetNetworkBluetoothClientsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithT0(t0 *string) *GetNetworkBluetoothClientsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) WithTimespan(timespan *float32) *GetNetworkBluetoothClientsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get network bluetooth clients params
func (o *GetNetworkBluetoothClientsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkBluetoothClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeConnectivityHistory != nil {

		// query param includeConnectivityHistory
		var qrIncludeConnectivityHistory bool

		if o.IncludeConnectivityHistory != nil {
			qrIncludeConnectivityHistory = *o.IncludeConnectivityHistory
		}
		qIncludeConnectivityHistory := swag.FormatBool(qrIncludeConnectivityHistory)
		if qIncludeConnectivityHistory != "" {

			if err := r.SetQueryParam("includeConnectivityHistory", qIncludeConnectivityHistory); err != nil {
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

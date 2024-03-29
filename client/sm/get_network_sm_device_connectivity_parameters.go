// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// NewGetNetworkSmDeviceConnectivityParams creates a new GetNetworkSmDeviceConnectivityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmDeviceConnectivityParams() *GetNetworkSmDeviceConnectivityParams {
	return &GetNetworkSmDeviceConnectivityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmDeviceConnectivityParamsWithTimeout creates a new GetNetworkSmDeviceConnectivityParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmDeviceConnectivityParamsWithTimeout(timeout time.Duration) *GetNetworkSmDeviceConnectivityParams {
	return &GetNetworkSmDeviceConnectivityParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmDeviceConnectivityParamsWithContext creates a new GetNetworkSmDeviceConnectivityParams object
// with the ability to set a context for a request.
func NewGetNetworkSmDeviceConnectivityParamsWithContext(ctx context.Context) *GetNetworkSmDeviceConnectivityParams {
	return &GetNetworkSmDeviceConnectivityParams{
		Context: ctx,
	}
}

// NewGetNetworkSmDeviceConnectivityParamsWithHTTPClient creates a new GetNetworkSmDeviceConnectivityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmDeviceConnectivityParamsWithHTTPClient(client *http.Client) *GetNetworkSmDeviceConnectivityParams {
	return &GetNetworkSmDeviceConnectivityParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSmDeviceConnectivityParams contains all the parameters to send to the API endpoint

	for the get network sm device connectivity operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSmDeviceConnectivityParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm device connectivity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceConnectivityParams) WithDefaults() *GetNetworkSmDeviceConnectivityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm device connectivity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceConnectivityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithTimeout(timeout time.Duration) *GetNetworkSmDeviceConnectivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithContext(ctx context.Context) *GetNetworkSmDeviceConnectivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithHTTPClient(client *http.Client) *GetNetworkSmDeviceConnectivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithDeviceID(deviceID string) *GetNetworkSmDeviceConnectivityParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithEndingBefore adds the endingBefore to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithEndingBefore(endingBefore *string) *GetNetworkSmDeviceConnectivityParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithNetworkID(networkID string) *GetNetworkSmDeviceConnectivityParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithPerPage(perPage *int64) *GetNetworkSmDeviceConnectivityParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) WithStartingAfter(startingAfter *string) *GetNetworkSmDeviceConnectivityParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network sm device connectivity params
func (o *GetNetworkSmDeviceConnectivityParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmDeviceConnectivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

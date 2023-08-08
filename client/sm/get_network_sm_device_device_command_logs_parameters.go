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

// NewGetNetworkSmDeviceDeviceCommandLogsParams creates a new GetNetworkSmDeviceDeviceCommandLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmDeviceDeviceCommandLogsParams() *GetNetworkSmDeviceDeviceCommandLogsParams {
	return &GetNetworkSmDeviceDeviceCommandLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmDeviceDeviceCommandLogsParamsWithTimeout creates a new GetNetworkSmDeviceDeviceCommandLogsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmDeviceDeviceCommandLogsParamsWithTimeout(timeout time.Duration) *GetNetworkSmDeviceDeviceCommandLogsParams {
	return &GetNetworkSmDeviceDeviceCommandLogsParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmDeviceDeviceCommandLogsParamsWithContext creates a new GetNetworkSmDeviceDeviceCommandLogsParams object
// with the ability to set a context for a request.
func NewGetNetworkSmDeviceDeviceCommandLogsParamsWithContext(ctx context.Context) *GetNetworkSmDeviceDeviceCommandLogsParams {
	return &GetNetworkSmDeviceDeviceCommandLogsParams{
		Context: ctx,
	}
}

// NewGetNetworkSmDeviceDeviceCommandLogsParamsWithHTTPClient creates a new GetNetworkSmDeviceDeviceCommandLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmDeviceDeviceCommandLogsParamsWithHTTPClient(client *http.Client) *GetNetworkSmDeviceDeviceCommandLogsParams {
	return &GetNetworkSmDeviceDeviceCommandLogsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSmDeviceDeviceCommandLogsParams contains all the parameters to send to the API endpoint

	for the get network sm device device command logs operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSmDeviceDeviceCommandLogsParams struct {

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

// WithDefaults hydrates default values in the get network sm device device command logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithDefaults() *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm device device command logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithTimeout(timeout time.Duration) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithContext(ctx context.Context) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithHTTPClient(client *http.Client) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithDeviceID(deviceID string) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithEndingBefore adds the endingBefore to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithEndingBefore(endingBefore *string) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithNetworkID adds the networkID to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithNetworkID(networkID string) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPerPage adds the perPage to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithPerPage(perPage *int64) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WithStartingAfter(startingAfter *string) *GetNetworkSmDeviceDeviceCommandLogsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get network sm device device command logs params
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmDeviceDeviceCommandLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

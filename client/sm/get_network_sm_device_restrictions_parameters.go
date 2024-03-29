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
)

// NewGetNetworkSmDeviceRestrictionsParams creates a new GetNetworkSmDeviceRestrictionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmDeviceRestrictionsParams() *GetNetworkSmDeviceRestrictionsParams {
	return &GetNetworkSmDeviceRestrictionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmDeviceRestrictionsParamsWithTimeout creates a new GetNetworkSmDeviceRestrictionsParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmDeviceRestrictionsParamsWithTimeout(timeout time.Duration) *GetNetworkSmDeviceRestrictionsParams {
	return &GetNetworkSmDeviceRestrictionsParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmDeviceRestrictionsParamsWithContext creates a new GetNetworkSmDeviceRestrictionsParams object
// with the ability to set a context for a request.
func NewGetNetworkSmDeviceRestrictionsParamsWithContext(ctx context.Context) *GetNetworkSmDeviceRestrictionsParams {
	return &GetNetworkSmDeviceRestrictionsParams{
		Context: ctx,
	}
}

// NewGetNetworkSmDeviceRestrictionsParamsWithHTTPClient creates a new GetNetworkSmDeviceRestrictionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmDeviceRestrictionsParamsWithHTTPClient(client *http.Client) *GetNetworkSmDeviceRestrictionsParams {
	return &GetNetworkSmDeviceRestrictionsParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSmDeviceRestrictionsParams contains all the parameters to send to the API endpoint

	for the get network sm device restrictions operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSmDeviceRestrictionsParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network sm device restrictions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceRestrictionsParams) WithDefaults() *GetNetworkSmDeviceRestrictionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm device restrictions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceRestrictionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) WithTimeout(timeout time.Duration) *GetNetworkSmDeviceRestrictionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) WithContext(ctx context.Context) *GetNetworkSmDeviceRestrictionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) WithHTTPClient(client *http.Client) *GetNetworkSmDeviceRestrictionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) WithDeviceID(deviceID string) *GetNetworkSmDeviceRestrictionsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNetworkID adds the networkID to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) WithNetworkID(networkID string) *GetNetworkSmDeviceRestrictionsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm device restrictions params
func (o *GetNetworkSmDeviceRestrictionsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmDeviceRestrictionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

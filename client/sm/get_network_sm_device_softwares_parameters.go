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

// NewGetNetworkSmDeviceSoftwaresParams creates a new GetNetworkSmDeviceSoftwaresParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkSmDeviceSoftwaresParams() *GetNetworkSmDeviceSoftwaresParams {
	return &GetNetworkSmDeviceSoftwaresParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmDeviceSoftwaresParamsWithTimeout creates a new GetNetworkSmDeviceSoftwaresParams object
// with the ability to set a timeout on a request.
func NewGetNetworkSmDeviceSoftwaresParamsWithTimeout(timeout time.Duration) *GetNetworkSmDeviceSoftwaresParams {
	return &GetNetworkSmDeviceSoftwaresParams{
		timeout: timeout,
	}
}

// NewGetNetworkSmDeviceSoftwaresParamsWithContext creates a new GetNetworkSmDeviceSoftwaresParams object
// with the ability to set a context for a request.
func NewGetNetworkSmDeviceSoftwaresParamsWithContext(ctx context.Context) *GetNetworkSmDeviceSoftwaresParams {
	return &GetNetworkSmDeviceSoftwaresParams{
		Context: ctx,
	}
}

// NewGetNetworkSmDeviceSoftwaresParamsWithHTTPClient creates a new GetNetworkSmDeviceSoftwaresParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkSmDeviceSoftwaresParamsWithHTTPClient(client *http.Client) *GetNetworkSmDeviceSoftwaresParams {
	return &GetNetworkSmDeviceSoftwaresParams{
		HTTPClient: client,
	}
}

/*
GetNetworkSmDeviceSoftwaresParams contains all the parameters to send to the API endpoint

	for the get network sm device softwares operation.

	Typically these are written to a http.Request.
*/
type GetNetworkSmDeviceSoftwaresParams struct {

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

// WithDefaults hydrates default values in the get network sm device softwares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceSoftwaresParams) WithDefaults() *GetNetworkSmDeviceSoftwaresParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network sm device softwares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkSmDeviceSoftwaresParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) WithTimeout(timeout time.Duration) *GetNetworkSmDeviceSoftwaresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) WithContext(ctx context.Context) *GetNetworkSmDeviceSoftwaresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) WithHTTPClient(client *http.Client) *GetNetworkSmDeviceSoftwaresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) WithDeviceID(deviceID string) *GetNetworkSmDeviceSoftwaresParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNetworkID adds the networkID to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) WithNetworkID(networkID string) *GetNetworkSmDeviceSoftwaresParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm device softwares params
func (o *GetNetworkSmDeviceSoftwaresParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmDeviceSoftwaresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

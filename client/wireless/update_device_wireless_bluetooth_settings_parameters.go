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
)

// NewUpdateDeviceWirelessBluetoothSettingsParams creates a new UpdateDeviceWirelessBluetoothSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceWirelessBluetoothSettingsParams() *UpdateDeviceWirelessBluetoothSettingsParams {
	return &UpdateDeviceWirelessBluetoothSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceWirelessBluetoothSettingsParamsWithTimeout creates a new UpdateDeviceWirelessBluetoothSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceWirelessBluetoothSettingsParamsWithTimeout(timeout time.Duration) *UpdateDeviceWirelessBluetoothSettingsParams {
	return &UpdateDeviceWirelessBluetoothSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceWirelessBluetoothSettingsParamsWithContext creates a new UpdateDeviceWirelessBluetoothSettingsParams object
// with the ability to set a context for a request.
func NewUpdateDeviceWirelessBluetoothSettingsParamsWithContext(ctx context.Context) *UpdateDeviceWirelessBluetoothSettingsParams {
	return &UpdateDeviceWirelessBluetoothSettingsParams{
		Context: ctx,
	}
}

// NewUpdateDeviceWirelessBluetoothSettingsParamsWithHTTPClient creates a new UpdateDeviceWirelessBluetoothSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceWirelessBluetoothSettingsParamsWithHTTPClient(client *http.Client) *UpdateDeviceWirelessBluetoothSettingsParams {
	return &UpdateDeviceWirelessBluetoothSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceWirelessBluetoothSettingsParams contains all the parameters to send to the API endpoint

	for the update device wireless bluetooth settings operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceWirelessBluetoothSettingsParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	// UpdateDeviceWirelessBluetoothSettings.
	UpdateDeviceWirelessBluetoothSettings UpdateDeviceWirelessBluetoothSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device wireless bluetooth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithDefaults() *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device wireless bluetooth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithTimeout(timeout time.Duration) *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithContext(ctx context.Context) *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithHTTPClient(client *http.Client) *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithSerial(serial string) *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceWirelessBluetoothSettings adds the updateDeviceWirelessBluetoothSettings to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WithUpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings UpdateDeviceWirelessBluetoothSettingsBody) *UpdateDeviceWirelessBluetoothSettingsParams {
	o.SetUpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings)
	return o
}

// SetUpdateDeviceWirelessBluetoothSettings adds the updateDeviceWirelessBluetoothSettings to the update device wireless bluetooth settings params
func (o *UpdateDeviceWirelessBluetoothSettingsParams) SetUpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings UpdateDeviceWirelessBluetoothSettingsBody) {
	o.UpdateDeviceWirelessBluetoothSettings = updateDeviceWirelessBluetoothSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceWirelessBluetoothSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceWirelessBluetoothSettings); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewGetDeviceWirelessBluetoothSettingsParams creates a new GetDeviceWirelessBluetoothSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceWirelessBluetoothSettingsParams() *GetDeviceWirelessBluetoothSettingsParams {
	return &GetDeviceWirelessBluetoothSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceWirelessBluetoothSettingsParamsWithTimeout creates a new GetDeviceWirelessBluetoothSettingsParams object
// with the ability to set a timeout on a request.
func NewGetDeviceWirelessBluetoothSettingsParamsWithTimeout(timeout time.Duration) *GetDeviceWirelessBluetoothSettingsParams {
	return &GetDeviceWirelessBluetoothSettingsParams{
		timeout: timeout,
	}
}

// NewGetDeviceWirelessBluetoothSettingsParamsWithContext creates a new GetDeviceWirelessBluetoothSettingsParams object
// with the ability to set a context for a request.
func NewGetDeviceWirelessBluetoothSettingsParamsWithContext(ctx context.Context) *GetDeviceWirelessBluetoothSettingsParams {
	return &GetDeviceWirelessBluetoothSettingsParams{
		Context: ctx,
	}
}

// NewGetDeviceWirelessBluetoothSettingsParamsWithHTTPClient creates a new GetDeviceWirelessBluetoothSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceWirelessBluetoothSettingsParamsWithHTTPClient(client *http.Client) *GetDeviceWirelessBluetoothSettingsParams {
	return &GetDeviceWirelessBluetoothSettingsParams{
		HTTPClient: client,
	}
}

/*
GetDeviceWirelessBluetoothSettingsParams contains all the parameters to send to the API endpoint

	for the get device wireless bluetooth settings operation.

	Typically these are written to a http.Request.
*/
type GetDeviceWirelessBluetoothSettingsParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device wireless bluetooth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceWirelessBluetoothSettingsParams) WithDefaults() *GetDeviceWirelessBluetoothSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device wireless bluetooth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceWirelessBluetoothSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) WithTimeout(timeout time.Duration) *GetDeviceWirelessBluetoothSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) WithContext(ctx context.Context) *GetDeviceWirelessBluetoothSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) WithHTTPClient(client *http.Client) *GetDeviceWirelessBluetoothSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) WithSerial(serial string) *GetDeviceWirelessBluetoothSettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device wireless bluetooth settings params
func (o *GetDeviceWirelessBluetoothSettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceWirelessBluetoothSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

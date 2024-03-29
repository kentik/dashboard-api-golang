// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewUpdateDeviceParams creates a new UpdateDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceParams() *UpdateDeviceParams {
	return &UpdateDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceParamsWithTimeout creates a new UpdateDeviceParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceParamsWithTimeout(timeout time.Duration) *UpdateDeviceParams {
	return &UpdateDeviceParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceParamsWithContext creates a new UpdateDeviceParams object
// with the ability to set a context for a request.
func NewUpdateDeviceParamsWithContext(ctx context.Context) *UpdateDeviceParams {
	return &UpdateDeviceParams{
		Context: ctx,
	}
}

// NewUpdateDeviceParamsWithHTTPClient creates a new UpdateDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceParamsWithHTTPClient(client *http.Client) *UpdateDeviceParams {
	return &UpdateDeviceParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceParams contains all the parameters to send to the API endpoint

	for the update device operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	// UpdateDevice.
	UpdateDevice UpdateDeviceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceParams) WithDefaults() *UpdateDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device params
func (o *UpdateDeviceParams) WithTimeout(timeout time.Duration) *UpdateDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device params
func (o *UpdateDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device params
func (o *UpdateDeviceParams) WithContext(ctx context.Context) *UpdateDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device params
func (o *UpdateDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device params
func (o *UpdateDeviceParams) WithHTTPClient(client *http.Client) *UpdateDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device params
func (o *UpdateDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device params
func (o *UpdateDeviceParams) WithSerial(serial string) *UpdateDeviceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device params
func (o *UpdateDeviceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDevice adds the updateDevice to the update device params
func (o *UpdateDeviceParams) WithUpdateDevice(updateDevice UpdateDeviceBody) *UpdateDeviceParams {
	o.SetUpdateDevice(updateDevice)
	return o
}

// SetUpdateDevice adds the updateDevice to the update device params
func (o *UpdateDeviceParams) SetUpdateDevice(updateDevice UpdateDeviceBody) {
	o.UpdateDevice = updateDevice
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDevice); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

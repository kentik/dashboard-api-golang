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

// NewUpdateDeviceManagementInterfaceParams creates a new UpdateDeviceManagementInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceManagementInterfaceParams() *UpdateDeviceManagementInterfaceParams {
	return &UpdateDeviceManagementInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceManagementInterfaceParamsWithTimeout creates a new UpdateDeviceManagementInterfaceParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceManagementInterfaceParamsWithTimeout(timeout time.Duration) *UpdateDeviceManagementInterfaceParams {
	return &UpdateDeviceManagementInterfaceParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceManagementInterfaceParamsWithContext creates a new UpdateDeviceManagementInterfaceParams object
// with the ability to set a context for a request.
func NewUpdateDeviceManagementInterfaceParamsWithContext(ctx context.Context) *UpdateDeviceManagementInterfaceParams {
	return &UpdateDeviceManagementInterfaceParams{
		Context: ctx,
	}
}

// NewUpdateDeviceManagementInterfaceParamsWithHTTPClient creates a new UpdateDeviceManagementInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceManagementInterfaceParamsWithHTTPClient(client *http.Client) *UpdateDeviceManagementInterfaceParams {
	return &UpdateDeviceManagementInterfaceParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceManagementInterfaceParams contains all the parameters to send to the API endpoint

	for the update device management interface operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceManagementInterfaceParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	// UpdateDeviceManagementInterface.
	UpdateDeviceManagementInterface UpdateDeviceManagementInterfaceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device management interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceManagementInterfaceParams) WithDefaults() *UpdateDeviceManagementInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device management interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceManagementInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) WithTimeout(timeout time.Duration) *UpdateDeviceManagementInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) WithContext(ctx context.Context) *UpdateDeviceManagementInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) WithHTTPClient(client *http.Client) *UpdateDeviceManagementInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) WithSerial(serial string) *UpdateDeviceManagementInterfaceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceManagementInterface adds the updateDeviceManagementInterface to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) WithUpdateDeviceManagementInterface(updateDeviceManagementInterface UpdateDeviceManagementInterfaceBody) *UpdateDeviceManagementInterfaceParams {
	o.SetUpdateDeviceManagementInterface(updateDeviceManagementInterface)
	return o
}

// SetUpdateDeviceManagementInterface adds the updateDeviceManagementInterface to the update device management interface params
func (o *UpdateDeviceManagementInterfaceParams) SetUpdateDeviceManagementInterface(updateDeviceManagementInterface UpdateDeviceManagementInterfaceBody) {
	o.UpdateDeviceManagementInterface = updateDeviceManagementInterface
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceManagementInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceManagementInterface); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

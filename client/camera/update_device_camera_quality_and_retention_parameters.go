// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewUpdateDeviceCameraQualityAndRetentionParams creates a new UpdateDeviceCameraQualityAndRetentionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceCameraQualityAndRetentionParams() *UpdateDeviceCameraQualityAndRetentionParams {
	return &UpdateDeviceCameraQualityAndRetentionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceCameraQualityAndRetentionParamsWithTimeout creates a new UpdateDeviceCameraQualityAndRetentionParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceCameraQualityAndRetentionParamsWithTimeout(timeout time.Duration) *UpdateDeviceCameraQualityAndRetentionParams {
	return &UpdateDeviceCameraQualityAndRetentionParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceCameraQualityAndRetentionParamsWithContext creates a new UpdateDeviceCameraQualityAndRetentionParams object
// with the ability to set a context for a request.
func NewUpdateDeviceCameraQualityAndRetentionParamsWithContext(ctx context.Context) *UpdateDeviceCameraQualityAndRetentionParams {
	return &UpdateDeviceCameraQualityAndRetentionParams{
		Context: ctx,
	}
}

// NewUpdateDeviceCameraQualityAndRetentionParamsWithHTTPClient creates a new UpdateDeviceCameraQualityAndRetentionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceCameraQualityAndRetentionParamsWithHTTPClient(client *http.Client) *UpdateDeviceCameraQualityAndRetentionParams {
	return &UpdateDeviceCameraQualityAndRetentionParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceCameraQualityAndRetentionParams contains all the parameters to send to the API endpoint

	for the update device camera quality and retention operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceCameraQualityAndRetentionParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	// UpdateDeviceCameraQualityAndRetention.
	UpdateDeviceCameraQualityAndRetention UpdateDeviceCameraQualityAndRetentionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device camera quality and retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithDefaults() *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device camera quality and retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithTimeout(timeout time.Duration) *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithContext(ctx context.Context) *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithHTTPClient(client *http.Client) *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithSerial(serial string) *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceCameraQualityAndRetention adds the updateDeviceCameraQualityAndRetention to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) WithUpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention UpdateDeviceCameraQualityAndRetentionBody) *UpdateDeviceCameraQualityAndRetentionParams {
	o.SetUpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention)
	return o
}

// SetUpdateDeviceCameraQualityAndRetention adds the updateDeviceCameraQualityAndRetention to the update device camera quality and retention params
func (o *UpdateDeviceCameraQualityAndRetentionParams) SetUpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention UpdateDeviceCameraQualityAndRetentionBody) {
	o.UpdateDeviceCameraQualityAndRetention = updateDeviceCameraQualityAndRetention
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceCameraQualityAndRetentionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceCameraQualityAndRetention); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

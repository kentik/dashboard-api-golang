// Code generated by go-swagger; DO NOT EDIT.

package sensor

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

// NewUpdateDeviceSensorRelationshipsParams creates a new UpdateDeviceSensorRelationshipsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceSensorRelationshipsParams() *UpdateDeviceSensorRelationshipsParams {
	return &UpdateDeviceSensorRelationshipsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceSensorRelationshipsParamsWithTimeout creates a new UpdateDeviceSensorRelationshipsParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceSensorRelationshipsParamsWithTimeout(timeout time.Duration) *UpdateDeviceSensorRelationshipsParams {
	return &UpdateDeviceSensorRelationshipsParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceSensorRelationshipsParamsWithContext creates a new UpdateDeviceSensorRelationshipsParams object
// with the ability to set a context for a request.
func NewUpdateDeviceSensorRelationshipsParamsWithContext(ctx context.Context) *UpdateDeviceSensorRelationshipsParams {
	return &UpdateDeviceSensorRelationshipsParams{
		Context: ctx,
	}
}

// NewUpdateDeviceSensorRelationshipsParamsWithHTTPClient creates a new UpdateDeviceSensorRelationshipsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceSensorRelationshipsParamsWithHTTPClient(client *http.Client) *UpdateDeviceSensorRelationshipsParams {
	return &UpdateDeviceSensorRelationshipsParams{
		HTTPClient: client,
	}
}

/*
UpdateDeviceSensorRelationshipsParams contains all the parameters to send to the API endpoint

	for the update device sensor relationships operation.

	Typically these are written to a http.Request.
*/
type UpdateDeviceSensorRelationshipsParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	// UpdateDeviceSensorRelationships.
	UpdateDeviceSensorRelationships UpdateDeviceSensorRelationshipsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device sensor relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceSensorRelationshipsParams) WithDefaults() *UpdateDeviceSensorRelationshipsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device sensor relationships params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceSensorRelationshipsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) WithTimeout(timeout time.Duration) *UpdateDeviceSensorRelationshipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) WithContext(ctx context.Context) *UpdateDeviceSensorRelationshipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) WithHTTPClient(client *http.Client) *UpdateDeviceSensorRelationshipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) WithSerial(serial string) *UpdateDeviceSensorRelationshipsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithUpdateDeviceSensorRelationships adds the updateDeviceSensorRelationships to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) WithUpdateDeviceSensorRelationships(updateDeviceSensorRelationships UpdateDeviceSensorRelationshipsBody) *UpdateDeviceSensorRelationshipsParams {
	o.SetUpdateDeviceSensorRelationships(updateDeviceSensorRelationships)
	return o
}

// SetUpdateDeviceSensorRelationships adds the updateDeviceSensorRelationships to the update device sensor relationships params
func (o *UpdateDeviceSensorRelationshipsParams) SetUpdateDeviceSensorRelationships(updateDeviceSensorRelationships UpdateDeviceSensorRelationshipsBody) {
	o.UpdateDeviceSensorRelationships = updateDeviceSensorRelationships
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceSensorRelationshipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateDeviceSensorRelationships); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

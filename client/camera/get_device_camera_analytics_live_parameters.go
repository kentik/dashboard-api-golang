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

// NewGetDeviceCameraAnalyticsLiveParams creates a new GetDeviceCameraAnalyticsLiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCameraAnalyticsLiveParams() *GetDeviceCameraAnalyticsLiveParams {
	return &GetDeviceCameraAnalyticsLiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCameraAnalyticsLiveParamsWithTimeout creates a new GetDeviceCameraAnalyticsLiveParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCameraAnalyticsLiveParamsWithTimeout(timeout time.Duration) *GetDeviceCameraAnalyticsLiveParams {
	return &GetDeviceCameraAnalyticsLiveParams{
		timeout: timeout,
	}
}

// NewGetDeviceCameraAnalyticsLiveParamsWithContext creates a new GetDeviceCameraAnalyticsLiveParams object
// with the ability to set a context for a request.
func NewGetDeviceCameraAnalyticsLiveParamsWithContext(ctx context.Context) *GetDeviceCameraAnalyticsLiveParams {
	return &GetDeviceCameraAnalyticsLiveParams{
		Context: ctx,
	}
}

// NewGetDeviceCameraAnalyticsLiveParamsWithHTTPClient creates a new GetDeviceCameraAnalyticsLiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCameraAnalyticsLiveParamsWithHTTPClient(client *http.Client) *GetDeviceCameraAnalyticsLiveParams {
	return &GetDeviceCameraAnalyticsLiveParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCameraAnalyticsLiveParams contains all the parameters to send to the API endpoint

	for the get device camera analytics live operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCameraAnalyticsLiveParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device camera analytics live params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraAnalyticsLiveParams) WithDefaults() *GetDeviceCameraAnalyticsLiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device camera analytics live params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraAnalyticsLiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) WithTimeout(timeout time.Duration) *GetDeviceCameraAnalyticsLiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) WithContext(ctx context.Context) *GetDeviceCameraAnalyticsLiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) WithHTTPClient(client *http.Client) *GetDeviceCameraAnalyticsLiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) WithSerial(serial string) *GetDeviceCameraAnalyticsLiveParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device camera analytics live params
func (o *GetDeviceCameraAnalyticsLiveParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCameraAnalyticsLiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

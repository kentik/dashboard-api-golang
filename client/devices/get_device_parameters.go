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

// NewGetDeviceParams creates a new GetDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceParams() *GetDeviceParams {
	return &GetDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceParamsWithTimeout creates a new GetDeviceParams object
// with the ability to set a timeout on a request.
func NewGetDeviceParamsWithTimeout(timeout time.Duration) *GetDeviceParams {
	return &GetDeviceParams{
		timeout: timeout,
	}
}

// NewGetDeviceParamsWithContext creates a new GetDeviceParams object
// with the ability to set a context for a request.
func NewGetDeviceParamsWithContext(ctx context.Context) *GetDeviceParams {
	return &GetDeviceParams{
		Context: ctx,
	}
}

// NewGetDeviceParamsWithHTTPClient creates a new GetDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceParamsWithHTTPClient(client *http.Client) *GetDeviceParams {
	return &GetDeviceParams{
		HTTPClient: client,
	}
}

/*
GetDeviceParams contains all the parameters to send to the API endpoint

	for the get device operation.

	Typically these are written to a http.Request.
*/
type GetDeviceParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceParams) WithDefaults() *GetDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device params
func (o *GetDeviceParams) WithTimeout(timeout time.Duration) *GetDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device params
func (o *GetDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device params
func (o *GetDeviceParams) WithContext(ctx context.Context) *GetDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device params
func (o *GetDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device params
func (o *GetDeviceParams) WithHTTPClient(client *http.Client) *GetDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device params
func (o *GetDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device params
func (o *GetDeviceParams) WithSerial(serial string) *GetDeviceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device params
func (o *GetDeviceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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
	"github.com/go-openapi/swag"
)

// NewGetDeviceSwitchPortsStatusesParams creates a new GetDeviceSwitchPortsStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceSwitchPortsStatusesParams() *GetDeviceSwitchPortsStatusesParams {
	return &GetDeviceSwitchPortsStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceSwitchPortsStatusesParamsWithTimeout creates a new GetDeviceSwitchPortsStatusesParams object
// with the ability to set a timeout on a request.
func NewGetDeviceSwitchPortsStatusesParamsWithTimeout(timeout time.Duration) *GetDeviceSwitchPortsStatusesParams {
	return &GetDeviceSwitchPortsStatusesParams{
		timeout: timeout,
	}
}

// NewGetDeviceSwitchPortsStatusesParamsWithContext creates a new GetDeviceSwitchPortsStatusesParams object
// with the ability to set a context for a request.
func NewGetDeviceSwitchPortsStatusesParamsWithContext(ctx context.Context) *GetDeviceSwitchPortsStatusesParams {
	return &GetDeviceSwitchPortsStatusesParams{
		Context: ctx,
	}
}

// NewGetDeviceSwitchPortsStatusesParamsWithHTTPClient creates a new GetDeviceSwitchPortsStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceSwitchPortsStatusesParamsWithHTTPClient(client *http.Client) *GetDeviceSwitchPortsStatusesParams {
	return &GetDeviceSwitchPortsStatusesParams{
		HTTPClient: client,
	}
}

/*
GetDeviceSwitchPortsStatusesParams contains all the parameters to send to the API endpoint

	for the get device switch ports statuses operation.

	Typically these are written to a http.Request.
*/
type GetDeviceSwitchPortsStatusesParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device switch ports statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchPortsStatusesParams) WithDefaults() *GetDeviceSwitchPortsStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device switch ports statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceSwitchPortsStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithTimeout(timeout time.Duration) *GetDeviceSwitchPortsStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithContext(ctx context.Context) *GetDeviceSwitchPortsStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithHTTPClient(client *http.Client) *GetDeviceSwitchPortsStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithSerial(serial string) *GetDeviceSwitchPortsStatusesParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithT0 adds the t0 to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithT0(t0 *string) *GetDeviceSwitchPortsStatusesParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithTimespan adds the timespan to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) WithTimespan(timespan *float32) *GetDeviceSwitchPortsStatusesParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get device switch ports statuses params
func (o *GetDeviceSwitchPortsStatusesParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceSwitchPortsStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

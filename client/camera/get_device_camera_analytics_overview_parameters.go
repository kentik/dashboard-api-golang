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
	"github.com/go-openapi/swag"
)

// NewGetDeviceCameraAnalyticsOverviewParams creates a new GetDeviceCameraAnalyticsOverviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceCameraAnalyticsOverviewParams() *GetDeviceCameraAnalyticsOverviewParams {
	return &GetDeviceCameraAnalyticsOverviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceCameraAnalyticsOverviewParamsWithTimeout creates a new GetDeviceCameraAnalyticsOverviewParams object
// with the ability to set a timeout on a request.
func NewGetDeviceCameraAnalyticsOverviewParamsWithTimeout(timeout time.Duration) *GetDeviceCameraAnalyticsOverviewParams {
	return &GetDeviceCameraAnalyticsOverviewParams{
		timeout: timeout,
	}
}

// NewGetDeviceCameraAnalyticsOverviewParamsWithContext creates a new GetDeviceCameraAnalyticsOverviewParams object
// with the ability to set a context for a request.
func NewGetDeviceCameraAnalyticsOverviewParamsWithContext(ctx context.Context) *GetDeviceCameraAnalyticsOverviewParams {
	return &GetDeviceCameraAnalyticsOverviewParams{
		Context: ctx,
	}
}

// NewGetDeviceCameraAnalyticsOverviewParamsWithHTTPClient creates a new GetDeviceCameraAnalyticsOverviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceCameraAnalyticsOverviewParamsWithHTTPClient(client *http.Client) *GetDeviceCameraAnalyticsOverviewParams {
	return &GetDeviceCameraAnalyticsOverviewParams{
		HTTPClient: client,
	}
}

/*
GetDeviceCameraAnalyticsOverviewParams contains all the parameters to send to the API endpoint

	for the get device camera analytics overview operation.

	Typically these are written to a http.Request.
*/
type GetDeviceCameraAnalyticsOverviewParams struct {

	/* ObjectType.

	   [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
	*/
	ObjectType *string

	/* Serial.

	   Serial
	*/
	Serial string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device camera analytics overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraAnalyticsOverviewParams) WithDefaults() *GetDeviceCameraAnalyticsOverviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device camera analytics overview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceCameraAnalyticsOverviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithTimeout(timeout time.Duration) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithContext(ctx context.Context) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithHTTPClient(client *http.Client) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithObjectType adds the objectType to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithObjectType(objectType *string) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetObjectType(objectType *string) {
	o.ObjectType = objectType
}

// WithSerial adds the serial to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithSerial(serial string) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithT0 adds the t0 to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithT0(t0 *string) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithT1(t1 *string) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) WithTimespan(timespan *float32) *GetDeviceCameraAnalyticsOverviewParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get device camera analytics overview params
func (o *GetDeviceCameraAnalyticsOverviewParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceCameraAnalyticsOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ObjectType != nil {

		// query param objectType
		var qrObjectType string

		if o.ObjectType != nil {
			qrObjectType = *o.ObjectType
		}
		qObjectType := qrObjectType
		if qObjectType != "" {

			if err := r.SetQueryParam("objectType", qObjectType); err != nil {
				return err
			}
		}
	}

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

	if o.T1 != nil {

		// query param t1
		var qrT1 string

		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {

			if err := r.SetQueryParam("t1", qT1); err != nil {
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

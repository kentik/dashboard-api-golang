// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrganizationSummaryTopDevicesByUsageParams creates a new GetOrganizationSummaryTopDevicesByUsageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationSummaryTopDevicesByUsageParams() *GetOrganizationSummaryTopDevicesByUsageParams {
	return &GetOrganizationSummaryTopDevicesByUsageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSummaryTopDevicesByUsageParamsWithTimeout creates a new GetOrganizationSummaryTopDevicesByUsageParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationSummaryTopDevicesByUsageParamsWithTimeout(timeout time.Duration) *GetOrganizationSummaryTopDevicesByUsageParams {
	return &GetOrganizationSummaryTopDevicesByUsageParams{
		timeout: timeout,
	}
}

// NewGetOrganizationSummaryTopDevicesByUsageParamsWithContext creates a new GetOrganizationSummaryTopDevicesByUsageParams object
// with the ability to set a context for a request.
func NewGetOrganizationSummaryTopDevicesByUsageParamsWithContext(ctx context.Context) *GetOrganizationSummaryTopDevicesByUsageParams {
	return &GetOrganizationSummaryTopDevicesByUsageParams{
		Context: ctx,
	}
}

// NewGetOrganizationSummaryTopDevicesByUsageParamsWithHTTPClient creates a new GetOrganizationSummaryTopDevicesByUsageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationSummaryTopDevicesByUsageParamsWithHTTPClient(client *http.Client) *GetOrganizationSummaryTopDevicesByUsageParams {
	return &GetOrganizationSummaryTopDevicesByUsageParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationSummaryTopDevicesByUsageParams contains all the parameters to send to the API endpoint

	for the get organization summary top devices by usage operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationSummaryTopDevicesByUsageParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* T0.

	   The beginning of the timespan for the data.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization summary top devices by usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithDefaults() *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization summary top devices by usage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithTimeout(timeout time.Duration) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithContext(ctx context.Context) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithHTTPClient(client *http.Client) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithOrganizationID(organizationID string) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithT0 adds the t0 to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithT0(t0 *string) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithT1(t1 *string) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WithTimespan(timespan *float32) *GetOrganizationSummaryTopDevicesByUsageParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization summary top devices by usage params
func (o *GetOrganizationSummaryTopDevicesByUsageParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSummaryTopDevicesByUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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

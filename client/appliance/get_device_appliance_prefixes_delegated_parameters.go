// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// NewGetDeviceAppliancePrefixesDelegatedParams creates a new GetDeviceAppliancePrefixesDelegatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceAppliancePrefixesDelegatedParams() *GetDeviceAppliancePrefixesDelegatedParams {
	return &GetDeviceAppliancePrefixesDelegatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceAppliancePrefixesDelegatedParamsWithTimeout creates a new GetDeviceAppliancePrefixesDelegatedParams object
// with the ability to set a timeout on a request.
func NewGetDeviceAppliancePrefixesDelegatedParamsWithTimeout(timeout time.Duration) *GetDeviceAppliancePrefixesDelegatedParams {
	return &GetDeviceAppliancePrefixesDelegatedParams{
		timeout: timeout,
	}
}

// NewGetDeviceAppliancePrefixesDelegatedParamsWithContext creates a new GetDeviceAppliancePrefixesDelegatedParams object
// with the ability to set a context for a request.
func NewGetDeviceAppliancePrefixesDelegatedParamsWithContext(ctx context.Context) *GetDeviceAppliancePrefixesDelegatedParams {
	return &GetDeviceAppliancePrefixesDelegatedParams{
		Context: ctx,
	}
}

// NewGetDeviceAppliancePrefixesDelegatedParamsWithHTTPClient creates a new GetDeviceAppliancePrefixesDelegatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceAppliancePrefixesDelegatedParamsWithHTTPClient(client *http.Client) *GetDeviceAppliancePrefixesDelegatedParams {
	return &GetDeviceAppliancePrefixesDelegatedParams{
		HTTPClient: client,
	}
}

/*
GetDeviceAppliancePrefixesDelegatedParams contains all the parameters to send to the API endpoint

	for the get device appliance prefixes delegated operation.

	Typically these are written to a http.Request.
*/
type GetDeviceAppliancePrefixesDelegatedParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device appliance prefixes delegated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceAppliancePrefixesDelegatedParams) WithDefaults() *GetDeviceAppliancePrefixesDelegatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device appliance prefixes delegated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceAppliancePrefixesDelegatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) WithTimeout(timeout time.Duration) *GetDeviceAppliancePrefixesDelegatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) WithContext(ctx context.Context) *GetDeviceAppliancePrefixesDelegatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) WithHTTPClient(client *http.Client) *GetDeviceAppliancePrefixesDelegatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) WithSerial(serial string) *GetDeviceAppliancePrefixesDelegatedParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device appliance prefixes delegated params
func (o *GetDeviceAppliancePrefixesDelegatedParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceAppliancePrefixesDelegatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

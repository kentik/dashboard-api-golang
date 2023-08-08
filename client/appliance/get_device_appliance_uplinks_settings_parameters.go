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

// NewGetDeviceApplianceUplinksSettingsParams creates a new GetDeviceApplianceUplinksSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceApplianceUplinksSettingsParams() *GetDeviceApplianceUplinksSettingsParams {
	return &GetDeviceApplianceUplinksSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceApplianceUplinksSettingsParamsWithTimeout creates a new GetDeviceApplianceUplinksSettingsParams object
// with the ability to set a timeout on a request.
func NewGetDeviceApplianceUplinksSettingsParamsWithTimeout(timeout time.Duration) *GetDeviceApplianceUplinksSettingsParams {
	return &GetDeviceApplianceUplinksSettingsParams{
		timeout: timeout,
	}
}

// NewGetDeviceApplianceUplinksSettingsParamsWithContext creates a new GetDeviceApplianceUplinksSettingsParams object
// with the ability to set a context for a request.
func NewGetDeviceApplianceUplinksSettingsParamsWithContext(ctx context.Context) *GetDeviceApplianceUplinksSettingsParams {
	return &GetDeviceApplianceUplinksSettingsParams{
		Context: ctx,
	}
}

// NewGetDeviceApplianceUplinksSettingsParamsWithHTTPClient creates a new GetDeviceApplianceUplinksSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceApplianceUplinksSettingsParamsWithHTTPClient(client *http.Client) *GetDeviceApplianceUplinksSettingsParams {
	return &GetDeviceApplianceUplinksSettingsParams{
		HTTPClient: client,
	}
}

/*
GetDeviceApplianceUplinksSettingsParams contains all the parameters to send to the API endpoint

	for the get device appliance uplinks settings operation.

	Typically these are written to a http.Request.
*/
type GetDeviceApplianceUplinksSettingsParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device appliance uplinks settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceApplianceUplinksSettingsParams) WithDefaults() *GetDeviceApplianceUplinksSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device appliance uplinks settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceApplianceUplinksSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) WithTimeout(timeout time.Duration) *GetDeviceApplianceUplinksSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) WithContext(ctx context.Context) *GetDeviceApplianceUplinksSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) WithHTTPClient(client *http.Client) *GetDeviceApplianceUplinksSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) WithSerial(serial string) *GetDeviceApplianceUplinksSettingsParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the get device appliance uplinks settings params
func (o *GetDeviceApplianceUplinksSettingsParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceApplianceUplinksSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

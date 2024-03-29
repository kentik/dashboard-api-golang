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

// NewCreateDeviceApplianceVmxAuthenticationTokenParams creates a new CreateDeviceApplianceVmxAuthenticationTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeviceApplianceVmxAuthenticationTokenParams() *CreateDeviceApplianceVmxAuthenticationTokenParams {
	return &CreateDeviceApplianceVmxAuthenticationTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithTimeout creates a new CreateDeviceApplianceVmxAuthenticationTokenParams object
// with the ability to set a timeout on a request.
func NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithTimeout(timeout time.Duration) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	return &CreateDeviceApplianceVmxAuthenticationTokenParams{
		timeout: timeout,
	}
}

// NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithContext creates a new CreateDeviceApplianceVmxAuthenticationTokenParams object
// with the ability to set a context for a request.
func NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithContext(ctx context.Context) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	return &CreateDeviceApplianceVmxAuthenticationTokenParams{
		Context: ctx,
	}
}

// NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithHTTPClient creates a new CreateDeviceApplianceVmxAuthenticationTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeviceApplianceVmxAuthenticationTokenParamsWithHTTPClient(client *http.Client) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	return &CreateDeviceApplianceVmxAuthenticationTokenParams{
		HTTPClient: client,
	}
}

/*
CreateDeviceApplianceVmxAuthenticationTokenParams contains all the parameters to send to the API endpoint

	for the create device appliance vmx authentication token operation.

	Typically these are written to a http.Request.
*/
type CreateDeviceApplianceVmxAuthenticationTokenParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create device appliance vmx authentication token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WithDefaults() *CreateDeviceApplianceVmxAuthenticationTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create device appliance vmx authentication token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WithTimeout(timeout time.Duration) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WithContext(ctx context.Context) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WithHTTPClient(client *http.Client) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WithSerial(serial string) *CreateDeviceApplianceVmxAuthenticationTokenParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the create device appliance vmx authentication token params
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceApplianceVmxAuthenticationTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

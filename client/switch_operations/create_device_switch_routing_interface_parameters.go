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
)

// NewCreateDeviceSwitchRoutingInterfaceParams creates a new CreateDeviceSwitchRoutingInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeviceSwitchRoutingInterfaceParams() *CreateDeviceSwitchRoutingInterfaceParams {
	return &CreateDeviceSwitchRoutingInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceSwitchRoutingInterfaceParamsWithTimeout creates a new CreateDeviceSwitchRoutingInterfaceParams object
// with the ability to set a timeout on a request.
func NewCreateDeviceSwitchRoutingInterfaceParamsWithTimeout(timeout time.Duration) *CreateDeviceSwitchRoutingInterfaceParams {
	return &CreateDeviceSwitchRoutingInterfaceParams{
		timeout: timeout,
	}
}

// NewCreateDeviceSwitchRoutingInterfaceParamsWithContext creates a new CreateDeviceSwitchRoutingInterfaceParams object
// with the ability to set a context for a request.
func NewCreateDeviceSwitchRoutingInterfaceParamsWithContext(ctx context.Context) *CreateDeviceSwitchRoutingInterfaceParams {
	return &CreateDeviceSwitchRoutingInterfaceParams{
		Context: ctx,
	}
}

// NewCreateDeviceSwitchRoutingInterfaceParamsWithHTTPClient creates a new CreateDeviceSwitchRoutingInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeviceSwitchRoutingInterfaceParamsWithHTTPClient(client *http.Client) *CreateDeviceSwitchRoutingInterfaceParams {
	return &CreateDeviceSwitchRoutingInterfaceParams{
		HTTPClient: client,
	}
}

/*
CreateDeviceSwitchRoutingInterfaceParams contains all the parameters to send to the API endpoint

	for the create device switch routing interface operation.

	Typically these are written to a http.Request.
*/
type CreateDeviceSwitchRoutingInterfaceParams struct {

	// CreateDeviceSwitchRoutingInterface.
	CreateDeviceSwitchRoutingInterface CreateDeviceSwitchRoutingInterfaceBody

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create device switch routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithDefaults() *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create device switch routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithTimeout(timeout time.Duration) *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithContext(ctx context.Context) *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithHTTPClient(client *http.Client) *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateDeviceSwitchRoutingInterface adds the createDeviceSwitchRoutingInterface to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithCreateDeviceSwitchRoutingInterface(createDeviceSwitchRoutingInterface CreateDeviceSwitchRoutingInterfaceBody) *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetCreateDeviceSwitchRoutingInterface(createDeviceSwitchRoutingInterface)
	return o
}

// SetCreateDeviceSwitchRoutingInterface adds the createDeviceSwitchRoutingInterface to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetCreateDeviceSwitchRoutingInterface(createDeviceSwitchRoutingInterface CreateDeviceSwitchRoutingInterfaceBody) {
	o.CreateDeviceSwitchRoutingInterface = createDeviceSwitchRoutingInterface
}

// WithSerial adds the serial to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) WithSerial(serial string) *CreateDeviceSwitchRoutingInterfaceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the create device switch routing interface params
func (o *CreateDeviceSwitchRoutingInterfaceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceSwitchRoutingInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateDeviceSwitchRoutingInterface); err != nil {
		return err
	}

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

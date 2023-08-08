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

// NewDeleteDeviceSwitchRoutingInterfaceParams creates a new DeleteDeviceSwitchRoutingInterfaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceSwitchRoutingInterfaceParams() *DeleteDeviceSwitchRoutingInterfaceParams {
	return &DeleteDeviceSwitchRoutingInterfaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceSwitchRoutingInterfaceParamsWithTimeout creates a new DeleteDeviceSwitchRoutingInterfaceParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceSwitchRoutingInterfaceParamsWithTimeout(timeout time.Duration) *DeleteDeviceSwitchRoutingInterfaceParams {
	return &DeleteDeviceSwitchRoutingInterfaceParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceSwitchRoutingInterfaceParamsWithContext creates a new DeleteDeviceSwitchRoutingInterfaceParams object
// with the ability to set a context for a request.
func NewDeleteDeviceSwitchRoutingInterfaceParamsWithContext(ctx context.Context) *DeleteDeviceSwitchRoutingInterfaceParams {
	return &DeleteDeviceSwitchRoutingInterfaceParams{
		Context: ctx,
	}
}

// NewDeleteDeviceSwitchRoutingInterfaceParamsWithHTTPClient creates a new DeleteDeviceSwitchRoutingInterfaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceSwitchRoutingInterfaceParamsWithHTTPClient(client *http.Client) *DeleteDeviceSwitchRoutingInterfaceParams {
	return &DeleteDeviceSwitchRoutingInterfaceParams{
		HTTPClient: client,
	}
}

/*
DeleteDeviceSwitchRoutingInterfaceParams contains all the parameters to send to the API endpoint

	for the delete device switch routing interface operation.

	Typically these are written to a http.Request.
*/
type DeleteDeviceSwitchRoutingInterfaceParams struct {

	/* InterfaceID.

	   Interface ID
	*/
	InterfaceID string

	/* Serial.

	   Serial
	*/
	Serial string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device switch routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithDefaults() *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device switch routing interface params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithTimeout(timeout time.Duration) *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithContext(ctx context.Context) *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithHTTPClient(client *http.Client) *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterfaceID adds the interfaceID to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithInterfaceID(interfaceID string) *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetInterfaceID(interfaceID string) {
	o.InterfaceID = interfaceID
}

// WithSerial adds the serial to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WithSerial(serial string) *DeleteDeviceSwitchRoutingInterfaceParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the delete device switch routing interface params
func (o *DeleteDeviceSwitchRoutingInterfaceParams) SetSerial(serial string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceSwitchRoutingInterfaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param interfaceId
	if err := r.SetPathParam("interfaceId", o.InterfaceID); err != nil {
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

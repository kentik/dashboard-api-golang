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

// NewDeleteDeviceSwitchRoutingStaticRouteParams creates a new DeleteDeviceSwitchRoutingStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceSwitchRoutingStaticRouteParams() *DeleteDeviceSwitchRoutingStaticRouteParams {
	return &DeleteDeviceSwitchRoutingStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceSwitchRoutingStaticRouteParamsWithTimeout creates a new DeleteDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceSwitchRoutingStaticRouteParamsWithTimeout(timeout time.Duration) *DeleteDeviceSwitchRoutingStaticRouteParams {
	return &DeleteDeviceSwitchRoutingStaticRouteParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceSwitchRoutingStaticRouteParamsWithContext creates a new DeleteDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a context for a request.
func NewDeleteDeviceSwitchRoutingStaticRouteParamsWithContext(ctx context.Context) *DeleteDeviceSwitchRoutingStaticRouteParams {
	return &DeleteDeviceSwitchRoutingStaticRouteParams{
		Context: ctx,
	}
}

// NewDeleteDeviceSwitchRoutingStaticRouteParamsWithHTTPClient creates a new DeleteDeviceSwitchRoutingStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceSwitchRoutingStaticRouteParamsWithHTTPClient(client *http.Client) *DeleteDeviceSwitchRoutingStaticRouteParams {
	return &DeleteDeviceSwitchRoutingStaticRouteParams{
		HTTPClient: client,
	}
}

/*
DeleteDeviceSwitchRoutingStaticRouteParams contains all the parameters to send to the API endpoint

	for the delete device switch routing static route operation.

	Typically these are written to a http.Request.
*/
type DeleteDeviceSwitchRoutingStaticRouteParams struct {

	/* Serial.

	   Serial
	*/
	Serial string

	/* StaticRouteID.

	   Static route ID
	*/
	StaticRouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device switch routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithDefaults() *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device switch routing static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithTimeout(timeout time.Duration) *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithContext(ctx context.Context) *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithHTTPClient(client *http.Client) *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSerial adds the serial to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithSerial(serial string) *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetSerial(serial string) {
	o.Serial = serial
}

// WithStaticRouteID adds the staticRouteID to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WithStaticRouteID(staticRouteID string) *DeleteDeviceSwitchRoutingStaticRouteParams {
	o.SetStaticRouteID(staticRouteID)
	return o
}

// SetStaticRouteID adds the staticRouteId to the delete device switch routing static route params
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) SetStaticRouteID(staticRouteID string) {
	o.StaticRouteID = staticRouteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceSwitchRoutingStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serial
	if err := r.SetPathParam("serial", o.Serial); err != nil {
		return err
	}

	// path param staticRouteId
	if err := r.SetPathParam("staticRouteId", o.StaticRouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

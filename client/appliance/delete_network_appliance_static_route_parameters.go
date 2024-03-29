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

// NewDeleteNetworkApplianceStaticRouteParams creates a new DeleteNetworkApplianceStaticRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkApplianceStaticRouteParams() *DeleteNetworkApplianceStaticRouteParams {
	return &DeleteNetworkApplianceStaticRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkApplianceStaticRouteParamsWithTimeout creates a new DeleteNetworkApplianceStaticRouteParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkApplianceStaticRouteParamsWithTimeout(timeout time.Duration) *DeleteNetworkApplianceStaticRouteParams {
	return &DeleteNetworkApplianceStaticRouteParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkApplianceStaticRouteParamsWithContext creates a new DeleteNetworkApplianceStaticRouteParams object
// with the ability to set a context for a request.
func NewDeleteNetworkApplianceStaticRouteParamsWithContext(ctx context.Context) *DeleteNetworkApplianceStaticRouteParams {
	return &DeleteNetworkApplianceStaticRouteParams{
		Context: ctx,
	}
}

// NewDeleteNetworkApplianceStaticRouteParamsWithHTTPClient creates a new DeleteNetworkApplianceStaticRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkApplianceStaticRouteParamsWithHTTPClient(client *http.Client) *DeleteNetworkApplianceStaticRouteParams {
	return &DeleteNetworkApplianceStaticRouteParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkApplianceStaticRouteParams contains all the parameters to send to the API endpoint

	for the delete network appliance static route operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkApplianceStaticRouteParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* StaticRouteID.

	   Static route ID
	*/
	StaticRouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkApplianceStaticRouteParams) WithDefaults() *DeleteNetworkApplianceStaticRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network appliance static route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkApplianceStaticRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) WithTimeout(timeout time.Duration) *DeleteNetworkApplianceStaticRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) WithContext(ctx context.Context) *DeleteNetworkApplianceStaticRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) WithHTTPClient(client *http.Client) *DeleteNetworkApplianceStaticRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) WithNetworkID(networkID string) *DeleteNetworkApplianceStaticRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithStaticRouteID adds the staticRouteID to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) WithStaticRouteID(staticRouteID string) *DeleteNetworkApplianceStaticRouteParams {
	o.SetStaticRouteID(staticRouteID)
	return o
}

// SetStaticRouteID adds the staticRouteId to the delete network appliance static route params
func (o *DeleteNetworkApplianceStaticRouteParams) SetStaticRouteID(staticRouteID string) {
	o.StaticRouteID = staticRouteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkApplianceStaticRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
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

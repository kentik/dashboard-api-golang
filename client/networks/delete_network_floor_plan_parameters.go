// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewDeleteNetworkFloorPlanParams creates a new DeleteNetworkFloorPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkFloorPlanParams() *DeleteNetworkFloorPlanParams {
	return &DeleteNetworkFloorPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkFloorPlanParamsWithTimeout creates a new DeleteNetworkFloorPlanParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkFloorPlanParamsWithTimeout(timeout time.Duration) *DeleteNetworkFloorPlanParams {
	return &DeleteNetworkFloorPlanParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkFloorPlanParamsWithContext creates a new DeleteNetworkFloorPlanParams object
// with the ability to set a context for a request.
func NewDeleteNetworkFloorPlanParamsWithContext(ctx context.Context) *DeleteNetworkFloorPlanParams {
	return &DeleteNetworkFloorPlanParams{
		Context: ctx,
	}
}

// NewDeleteNetworkFloorPlanParamsWithHTTPClient creates a new DeleteNetworkFloorPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkFloorPlanParamsWithHTTPClient(client *http.Client) *DeleteNetworkFloorPlanParams {
	return &DeleteNetworkFloorPlanParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkFloorPlanParams contains all the parameters to send to the API endpoint

	for the delete network floor plan operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkFloorPlanParams struct {

	/* FloorPlanID.

	   Floor plan ID
	*/
	FloorPlanID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network floor plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkFloorPlanParams) WithDefaults() *DeleteNetworkFloorPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network floor plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkFloorPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) WithTimeout(timeout time.Duration) *DeleteNetworkFloorPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) WithContext(ctx context.Context) *DeleteNetworkFloorPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) WithHTTPClient(client *http.Client) *DeleteNetworkFloorPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFloorPlanID adds the floorPlanID to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) WithFloorPlanID(floorPlanID string) *DeleteNetworkFloorPlanParams {
	o.SetFloorPlanID(floorPlanID)
	return o
}

// SetFloorPlanID adds the floorPlanId to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) SetFloorPlanID(floorPlanID string) {
	o.FloorPlanID = floorPlanID
}

// WithNetworkID adds the networkID to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) WithNetworkID(networkID string) *DeleteNetworkFloorPlanParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network floor plan params
func (o *DeleteNetworkFloorPlanParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkFloorPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param floorPlanId
	if err := r.SetPathParam("floorPlanId", o.FloorPlanID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

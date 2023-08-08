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

// NewDeleteNetworkSwitchLinkAggregationParams creates a new DeleteNetworkSwitchLinkAggregationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkSwitchLinkAggregationParams() *DeleteNetworkSwitchLinkAggregationParams {
	return &DeleteNetworkSwitchLinkAggregationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkSwitchLinkAggregationParamsWithTimeout creates a new DeleteNetworkSwitchLinkAggregationParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkSwitchLinkAggregationParamsWithTimeout(timeout time.Duration) *DeleteNetworkSwitchLinkAggregationParams {
	return &DeleteNetworkSwitchLinkAggregationParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkSwitchLinkAggregationParamsWithContext creates a new DeleteNetworkSwitchLinkAggregationParams object
// with the ability to set a context for a request.
func NewDeleteNetworkSwitchLinkAggregationParamsWithContext(ctx context.Context) *DeleteNetworkSwitchLinkAggregationParams {
	return &DeleteNetworkSwitchLinkAggregationParams{
		Context: ctx,
	}
}

// NewDeleteNetworkSwitchLinkAggregationParamsWithHTTPClient creates a new DeleteNetworkSwitchLinkAggregationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkSwitchLinkAggregationParamsWithHTTPClient(client *http.Client) *DeleteNetworkSwitchLinkAggregationParams {
	return &DeleteNetworkSwitchLinkAggregationParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkSwitchLinkAggregationParams contains all the parameters to send to the API endpoint

	for the delete network switch link aggregation operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkSwitchLinkAggregationParams struct {

	/* LinkAggregationID.

	   Link aggregation ID
	*/
	LinkAggregationID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network switch link aggregation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchLinkAggregationParams) WithDefaults() *DeleteNetworkSwitchLinkAggregationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network switch link aggregation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkSwitchLinkAggregationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) WithTimeout(timeout time.Duration) *DeleteNetworkSwitchLinkAggregationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) WithContext(ctx context.Context) *DeleteNetworkSwitchLinkAggregationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) WithHTTPClient(client *http.Client) *DeleteNetworkSwitchLinkAggregationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkAggregationID adds the linkAggregationID to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) WithLinkAggregationID(linkAggregationID string) *DeleteNetworkSwitchLinkAggregationParams {
	o.SetLinkAggregationID(linkAggregationID)
	return o
}

// SetLinkAggregationID adds the linkAggregationId to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) SetLinkAggregationID(linkAggregationID string) {
	o.LinkAggregationID = linkAggregationID
}

// WithNetworkID adds the networkID to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) WithNetworkID(networkID string) *DeleteNetworkSwitchLinkAggregationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network switch link aggregation params
func (o *DeleteNetworkSwitchLinkAggregationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkSwitchLinkAggregationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linkAggregationId
	if err := r.SetPathParam("linkAggregationId", o.LinkAggregationID); err != nil {
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

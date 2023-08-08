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

// NewDeleteNetworkApplianceRfProfileParams creates a new DeleteNetworkApplianceRfProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkApplianceRfProfileParams() *DeleteNetworkApplianceRfProfileParams {
	return &DeleteNetworkApplianceRfProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkApplianceRfProfileParamsWithTimeout creates a new DeleteNetworkApplianceRfProfileParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkApplianceRfProfileParamsWithTimeout(timeout time.Duration) *DeleteNetworkApplianceRfProfileParams {
	return &DeleteNetworkApplianceRfProfileParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkApplianceRfProfileParamsWithContext creates a new DeleteNetworkApplianceRfProfileParams object
// with the ability to set a context for a request.
func NewDeleteNetworkApplianceRfProfileParamsWithContext(ctx context.Context) *DeleteNetworkApplianceRfProfileParams {
	return &DeleteNetworkApplianceRfProfileParams{
		Context: ctx,
	}
}

// NewDeleteNetworkApplianceRfProfileParamsWithHTTPClient creates a new DeleteNetworkApplianceRfProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkApplianceRfProfileParamsWithHTTPClient(client *http.Client) *DeleteNetworkApplianceRfProfileParams {
	return &DeleteNetworkApplianceRfProfileParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkApplianceRfProfileParams contains all the parameters to send to the API endpoint

	for the delete network appliance rf profile operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkApplianceRfProfileParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* RfProfileID.

	   Rf profile ID
	*/
	RfProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network appliance rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkApplianceRfProfileParams) WithDefaults() *DeleteNetworkApplianceRfProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network appliance rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkApplianceRfProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) WithTimeout(timeout time.Duration) *DeleteNetworkApplianceRfProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) WithContext(ctx context.Context) *DeleteNetworkApplianceRfProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) WithHTTPClient(client *http.Client) *DeleteNetworkApplianceRfProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) WithNetworkID(networkID string) *DeleteNetworkApplianceRfProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRfProfileID adds the rfProfileID to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) WithRfProfileID(rfProfileID string) *DeleteNetworkApplianceRfProfileParams {
	o.SetRfProfileID(rfProfileID)
	return o
}

// SetRfProfileID adds the rfProfileId to the delete network appliance rf profile params
func (o *DeleteNetworkApplianceRfProfileParams) SetRfProfileID(rfProfileID string) {
	o.RfProfileID = rfProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkApplianceRfProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	// path param rfProfileId
	if err := r.SetPathParam("rfProfileId", o.RfProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
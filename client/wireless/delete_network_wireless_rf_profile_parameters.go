// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewDeleteNetworkWirelessRfProfileParams creates a new DeleteNetworkWirelessRfProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkWirelessRfProfileParams() *DeleteNetworkWirelessRfProfileParams {
	return &DeleteNetworkWirelessRfProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkWirelessRfProfileParamsWithTimeout creates a new DeleteNetworkWirelessRfProfileParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkWirelessRfProfileParamsWithTimeout(timeout time.Duration) *DeleteNetworkWirelessRfProfileParams {
	return &DeleteNetworkWirelessRfProfileParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkWirelessRfProfileParamsWithContext creates a new DeleteNetworkWirelessRfProfileParams object
// with the ability to set a context for a request.
func NewDeleteNetworkWirelessRfProfileParamsWithContext(ctx context.Context) *DeleteNetworkWirelessRfProfileParams {
	return &DeleteNetworkWirelessRfProfileParams{
		Context: ctx,
	}
}

// NewDeleteNetworkWirelessRfProfileParamsWithHTTPClient creates a new DeleteNetworkWirelessRfProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkWirelessRfProfileParamsWithHTTPClient(client *http.Client) *DeleteNetworkWirelessRfProfileParams {
	return &DeleteNetworkWirelessRfProfileParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkWirelessRfProfileParams contains all the parameters to send to the API endpoint

	for the delete network wireless rf profile operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkWirelessRfProfileParams struct {

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

// WithDefaults hydrates default values in the delete network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkWirelessRfProfileParams) WithDefaults() *DeleteNetworkWirelessRfProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network wireless rf profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkWirelessRfProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) WithTimeout(timeout time.Duration) *DeleteNetworkWirelessRfProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) WithContext(ctx context.Context) *DeleteNetworkWirelessRfProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) WithHTTPClient(client *http.Client) *DeleteNetworkWirelessRfProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) WithNetworkID(networkID string) *DeleteNetworkWirelessRfProfileParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRfProfileID adds the rfProfileID to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) WithRfProfileID(rfProfileID string) *DeleteNetworkWirelessRfProfileParams {
	o.SetRfProfileID(rfProfileID)
	return o
}

// SetRfProfileID adds the rfProfileId to the delete network wireless rf profile params
func (o *DeleteNetworkWirelessRfProfileParams) SetRfProfileID(rfProfileID string) {
	o.RfProfileID = rfProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkWirelessRfProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDeleteNetworkParams creates a new DeleteNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNetworkParams() *DeleteNetworkParams {
	return &DeleteNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkParamsWithTimeout creates a new DeleteNetworkParams object
// with the ability to set a timeout on a request.
func NewDeleteNetworkParamsWithTimeout(timeout time.Duration) *DeleteNetworkParams {
	return &DeleteNetworkParams{
		timeout: timeout,
	}
}

// NewDeleteNetworkParamsWithContext creates a new DeleteNetworkParams object
// with the ability to set a context for a request.
func NewDeleteNetworkParamsWithContext(ctx context.Context) *DeleteNetworkParams {
	return &DeleteNetworkParams{
		Context: ctx,
	}
}

// NewDeleteNetworkParamsWithHTTPClient creates a new DeleteNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNetworkParamsWithHTTPClient(client *http.Client) *DeleteNetworkParams {
	return &DeleteNetworkParams{
		HTTPClient: client,
	}
}

/*
DeleteNetworkParams contains all the parameters to send to the API endpoint

	for the delete network operation.

	Typically these are written to a http.Request.
*/
type DeleteNetworkParams struct {

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkParams) WithDefaults() *DeleteNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete network params
func (o *DeleteNetworkParams) WithTimeout(timeout time.Duration) *DeleteNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network params
func (o *DeleteNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network params
func (o *DeleteNetworkParams) WithContext(ctx context.Context) *DeleteNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network params
func (o *DeleteNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network params
func (o *DeleteNetworkParams) WithHTTPClient(client *http.Client) *DeleteNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network params
func (o *DeleteNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete network params
func (o *DeleteNetworkParams) WithNetworkID(networkID string) *DeleteNetworkParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete network params
func (o *DeleteNetworkParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

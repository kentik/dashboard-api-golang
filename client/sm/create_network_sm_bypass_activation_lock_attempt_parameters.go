// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// NewCreateNetworkSmBypassActivationLockAttemptParams creates a new CreateNetworkSmBypassActivationLockAttemptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkSmBypassActivationLockAttemptParams() *CreateNetworkSmBypassActivationLockAttemptParams {
	return &CreateNetworkSmBypassActivationLockAttemptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkSmBypassActivationLockAttemptParamsWithTimeout creates a new CreateNetworkSmBypassActivationLockAttemptParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkSmBypassActivationLockAttemptParamsWithTimeout(timeout time.Duration) *CreateNetworkSmBypassActivationLockAttemptParams {
	return &CreateNetworkSmBypassActivationLockAttemptParams{
		timeout: timeout,
	}
}

// NewCreateNetworkSmBypassActivationLockAttemptParamsWithContext creates a new CreateNetworkSmBypassActivationLockAttemptParams object
// with the ability to set a context for a request.
func NewCreateNetworkSmBypassActivationLockAttemptParamsWithContext(ctx context.Context) *CreateNetworkSmBypassActivationLockAttemptParams {
	return &CreateNetworkSmBypassActivationLockAttemptParams{
		Context: ctx,
	}
}

// NewCreateNetworkSmBypassActivationLockAttemptParamsWithHTTPClient creates a new CreateNetworkSmBypassActivationLockAttemptParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkSmBypassActivationLockAttemptParamsWithHTTPClient(client *http.Client) *CreateNetworkSmBypassActivationLockAttemptParams {
	return &CreateNetworkSmBypassActivationLockAttemptParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkSmBypassActivationLockAttemptParams contains all the parameters to send to the API endpoint

	for the create network sm bypass activation lock attempt operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkSmBypassActivationLockAttemptParams struct {

	// CreateNetworkSmBypassActivationLockAttempt.
	CreateNetworkSmBypassActivationLockAttempt CreateNetworkSmBypassActivationLockAttemptBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network sm bypass activation lock attempt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithDefaults() *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network sm bypass activation lock attempt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithTimeout(timeout time.Duration) *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithContext(ctx context.Context) *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithHTTPClient(client *http.Client) *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkSmBypassActivationLockAttempt adds the createNetworkSmBypassActivationLockAttempt to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithCreateNetworkSmBypassActivationLockAttempt(createNetworkSmBypassActivationLockAttempt CreateNetworkSmBypassActivationLockAttemptBody) *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetCreateNetworkSmBypassActivationLockAttempt(createNetworkSmBypassActivationLockAttempt)
	return o
}

// SetCreateNetworkSmBypassActivationLockAttempt adds the createNetworkSmBypassActivationLockAttempt to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetCreateNetworkSmBypassActivationLockAttempt(createNetworkSmBypassActivationLockAttempt CreateNetworkSmBypassActivationLockAttemptBody) {
	o.CreateNetworkSmBypassActivationLockAttempt = createNetworkSmBypassActivationLockAttempt
}

// WithNetworkID adds the networkID to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WithNetworkID(networkID string) *CreateNetworkSmBypassActivationLockAttemptParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network sm bypass activation lock attempt params
func (o *CreateNetworkSmBypassActivationLockAttemptParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkSmBypassActivationLockAttemptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkSmBypassActivationLockAttempt); err != nil {
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

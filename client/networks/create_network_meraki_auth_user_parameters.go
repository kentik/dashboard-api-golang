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

// NewCreateNetworkMerakiAuthUserParams creates a new CreateNetworkMerakiAuthUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkMerakiAuthUserParams() *CreateNetworkMerakiAuthUserParams {
	return &CreateNetworkMerakiAuthUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkMerakiAuthUserParamsWithTimeout creates a new CreateNetworkMerakiAuthUserParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkMerakiAuthUserParamsWithTimeout(timeout time.Duration) *CreateNetworkMerakiAuthUserParams {
	return &CreateNetworkMerakiAuthUserParams{
		timeout: timeout,
	}
}

// NewCreateNetworkMerakiAuthUserParamsWithContext creates a new CreateNetworkMerakiAuthUserParams object
// with the ability to set a context for a request.
func NewCreateNetworkMerakiAuthUserParamsWithContext(ctx context.Context) *CreateNetworkMerakiAuthUserParams {
	return &CreateNetworkMerakiAuthUserParams{
		Context: ctx,
	}
}

// NewCreateNetworkMerakiAuthUserParamsWithHTTPClient creates a new CreateNetworkMerakiAuthUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkMerakiAuthUserParamsWithHTTPClient(client *http.Client) *CreateNetworkMerakiAuthUserParams {
	return &CreateNetworkMerakiAuthUserParams{
		HTTPClient: client,
	}
}

/*
CreateNetworkMerakiAuthUserParams contains all the parameters to send to the API endpoint

	for the create network meraki auth user operation.

	Typically these are written to a http.Request.
*/
type CreateNetworkMerakiAuthUserParams struct {

	// CreateNetworkMerakiAuthUser.
	CreateNetworkMerakiAuthUser CreateNetworkMerakiAuthUserBody

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkMerakiAuthUserParams) WithDefaults() *CreateNetworkMerakiAuthUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network meraki auth user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkMerakiAuthUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) WithTimeout(timeout time.Duration) *CreateNetworkMerakiAuthUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) WithContext(ctx context.Context) *CreateNetworkMerakiAuthUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) WithHTTPClient(client *http.Client) *CreateNetworkMerakiAuthUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateNetworkMerakiAuthUser adds the createNetworkMerakiAuthUser to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) WithCreateNetworkMerakiAuthUser(createNetworkMerakiAuthUser CreateNetworkMerakiAuthUserBody) *CreateNetworkMerakiAuthUserParams {
	o.SetCreateNetworkMerakiAuthUser(createNetworkMerakiAuthUser)
	return o
}

// SetCreateNetworkMerakiAuthUser adds the createNetworkMerakiAuthUser to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) SetCreateNetworkMerakiAuthUser(createNetworkMerakiAuthUser CreateNetworkMerakiAuthUserBody) {
	o.CreateNetworkMerakiAuthUser = createNetworkMerakiAuthUser
}

// WithNetworkID adds the networkID to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) WithNetworkID(networkID string) *CreateNetworkMerakiAuthUserParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the create network meraki auth user params
func (o *CreateNetworkMerakiAuthUserParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkMerakiAuthUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateNetworkMerakiAuthUser); err != nil {
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

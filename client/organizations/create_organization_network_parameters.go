// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewCreateOrganizationNetworkParams creates a new CreateOrganizationNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationNetworkParams() *CreateOrganizationNetworkParams {
	return &CreateOrganizationNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationNetworkParamsWithTimeout creates a new CreateOrganizationNetworkParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationNetworkParamsWithTimeout(timeout time.Duration) *CreateOrganizationNetworkParams {
	return &CreateOrganizationNetworkParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationNetworkParamsWithContext creates a new CreateOrganizationNetworkParams object
// with the ability to set a context for a request.
func NewCreateOrganizationNetworkParamsWithContext(ctx context.Context) *CreateOrganizationNetworkParams {
	return &CreateOrganizationNetworkParams{
		Context: ctx,
	}
}

// NewCreateOrganizationNetworkParamsWithHTTPClient creates a new CreateOrganizationNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationNetworkParamsWithHTTPClient(client *http.Client) *CreateOrganizationNetworkParams {
	return &CreateOrganizationNetworkParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationNetworkParams contains all the parameters to send to the API endpoint

	for the create organization network operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationNetworkParams struct {

	// CreateOrganizationNetwork.
	CreateOrganizationNetwork CreateOrganizationNetworkBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationNetworkParams) WithDefaults() *CreateOrganizationNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization network params
func (o *CreateOrganizationNetworkParams) WithTimeout(timeout time.Duration) *CreateOrganizationNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization network params
func (o *CreateOrganizationNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization network params
func (o *CreateOrganizationNetworkParams) WithContext(ctx context.Context) *CreateOrganizationNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization network params
func (o *CreateOrganizationNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization network params
func (o *CreateOrganizationNetworkParams) WithHTTPClient(client *http.Client) *CreateOrganizationNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization network params
func (o *CreateOrganizationNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationNetwork adds the createOrganizationNetwork to the create organization network params
func (o *CreateOrganizationNetworkParams) WithCreateOrganizationNetwork(createOrganizationNetwork CreateOrganizationNetworkBody) *CreateOrganizationNetworkParams {
	o.SetCreateOrganizationNetwork(createOrganizationNetwork)
	return o
}

// SetCreateOrganizationNetwork adds the createOrganizationNetwork to the create organization network params
func (o *CreateOrganizationNetworkParams) SetCreateOrganizationNetwork(createOrganizationNetwork CreateOrganizationNetworkBody) {
	o.CreateOrganizationNetwork = createOrganizationNetwork
}

// WithOrganizationID adds the organizationID to the create organization network params
func (o *CreateOrganizationNetworkParams) WithOrganizationID(organizationID string) *CreateOrganizationNetworkParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization network params
func (o *CreateOrganizationNetworkParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationNetwork); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

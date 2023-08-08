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

// NewCreateOrganizationSamlIdpParams creates a new CreateOrganizationSamlIdpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationSamlIdpParams() *CreateOrganizationSamlIdpParams {
	return &CreateOrganizationSamlIdpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationSamlIdpParamsWithTimeout creates a new CreateOrganizationSamlIdpParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationSamlIdpParamsWithTimeout(timeout time.Duration) *CreateOrganizationSamlIdpParams {
	return &CreateOrganizationSamlIdpParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationSamlIdpParamsWithContext creates a new CreateOrganizationSamlIdpParams object
// with the ability to set a context for a request.
func NewCreateOrganizationSamlIdpParamsWithContext(ctx context.Context) *CreateOrganizationSamlIdpParams {
	return &CreateOrganizationSamlIdpParams{
		Context: ctx,
	}
}

// NewCreateOrganizationSamlIdpParamsWithHTTPClient creates a new CreateOrganizationSamlIdpParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationSamlIdpParamsWithHTTPClient(client *http.Client) *CreateOrganizationSamlIdpParams {
	return &CreateOrganizationSamlIdpParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationSamlIdpParams contains all the parameters to send to the API endpoint

	for the create organization saml idp operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationSamlIdpParams struct {

	// CreateOrganizationSamlIdp.
	CreateOrganizationSamlIdp CreateOrganizationSamlIdpBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization saml idp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationSamlIdpParams) WithDefaults() *CreateOrganizationSamlIdpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization saml idp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationSamlIdpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) WithTimeout(timeout time.Duration) *CreateOrganizationSamlIdpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) WithContext(ctx context.Context) *CreateOrganizationSamlIdpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) WithHTTPClient(client *http.Client) *CreateOrganizationSamlIdpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationSamlIdp adds the createOrganizationSamlIdp to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) WithCreateOrganizationSamlIdp(createOrganizationSamlIdp CreateOrganizationSamlIdpBody) *CreateOrganizationSamlIdpParams {
	o.SetCreateOrganizationSamlIdp(createOrganizationSamlIdp)
	return o
}

// SetCreateOrganizationSamlIdp adds the createOrganizationSamlIdp to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) SetCreateOrganizationSamlIdp(createOrganizationSamlIdp CreateOrganizationSamlIdpBody) {
	o.CreateOrganizationSamlIdp = createOrganizationSamlIdp
}

// WithOrganizationID adds the organizationID to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) WithOrganizationID(organizationID string) *CreateOrganizationSamlIdpParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization saml idp params
func (o *CreateOrganizationSamlIdpParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationSamlIdpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationSamlIdp); err != nil {
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

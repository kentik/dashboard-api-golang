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

// NewCreateOrganizationBrandingPolicyParams creates a new CreateOrganizationBrandingPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationBrandingPolicyParams() *CreateOrganizationBrandingPolicyParams {
	return &CreateOrganizationBrandingPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationBrandingPolicyParamsWithTimeout creates a new CreateOrganizationBrandingPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationBrandingPolicyParamsWithTimeout(timeout time.Duration) *CreateOrganizationBrandingPolicyParams {
	return &CreateOrganizationBrandingPolicyParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationBrandingPolicyParamsWithContext creates a new CreateOrganizationBrandingPolicyParams object
// with the ability to set a context for a request.
func NewCreateOrganizationBrandingPolicyParamsWithContext(ctx context.Context) *CreateOrganizationBrandingPolicyParams {
	return &CreateOrganizationBrandingPolicyParams{
		Context: ctx,
	}
}

// NewCreateOrganizationBrandingPolicyParamsWithHTTPClient creates a new CreateOrganizationBrandingPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationBrandingPolicyParamsWithHTTPClient(client *http.Client) *CreateOrganizationBrandingPolicyParams {
	return &CreateOrganizationBrandingPolicyParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationBrandingPolicyParams contains all the parameters to send to the API endpoint

	for the create organization branding policy operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationBrandingPolicyParams struct {

	// CreateOrganizationBrandingPolicy.
	CreateOrganizationBrandingPolicy CreateOrganizationBrandingPolicyBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization branding policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationBrandingPolicyParams) WithDefaults() *CreateOrganizationBrandingPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization branding policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationBrandingPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) WithTimeout(timeout time.Duration) *CreateOrganizationBrandingPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) WithContext(ctx context.Context) *CreateOrganizationBrandingPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) WithHTTPClient(client *http.Client) *CreateOrganizationBrandingPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationBrandingPolicy adds the createOrganizationBrandingPolicy to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) WithCreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy CreateOrganizationBrandingPolicyBody) *CreateOrganizationBrandingPolicyParams {
	o.SetCreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy)
	return o
}

// SetCreateOrganizationBrandingPolicy adds the createOrganizationBrandingPolicy to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) SetCreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy CreateOrganizationBrandingPolicyBody) {
	o.CreateOrganizationBrandingPolicy = createOrganizationBrandingPolicy
}

// WithOrganizationID adds the organizationID to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) WithOrganizationID(organizationID string) *CreateOrganizationBrandingPolicyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization branding policy params
func (o *CreateOrganizationBrandingPolicyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationBrandingPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationBrandingPolicy); err != nil {
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

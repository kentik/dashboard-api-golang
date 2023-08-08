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

// NewGetOrganizationBrandingPoliciesParams creates a new GetOrganizationBrandingPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationBrandingPoliciesParams() *GetOrganizationBrandingPoliciesParams {
	return &GetOrganizationBrandingPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationBrandingPoliciesParamsWithTimeout creates a new GetOrganizationBrandingPoliciesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationBrandingPoliciesParamsWithTimeout(timeout time.Duration) *GetOrganizationBrandingPoliciesParams {
	return &GetOrganizationBrandingPoliciesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationBrandingPoliciesParamsWithContext creates a new GetOrganizationBrandingPoliciesParams object
// with the ability to set a context for a request.
func NewGetOrganizationBrandingPoliciesParamsWithContext(ctx context.Context) *GetOrganizationBrandingPoliciesParams {
	return &GetOrganizationBrandingPoliciesParams{
		Context: ctx,
	}
}

// NewGetOrganizationBrandingPoliciesParamsWithHTTPClient creates a new GetOrganizationBrandingPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationBrandingPoliciesParamsWithHTTPClient(client *http.Client) *GetOrganizationBrandingPoliciesParams {
	return &GetOrganizationBrandingPoliciesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationBrandingPoliciesParams contains all the parameters to send to the API endpoint

	for the get organization branding policies operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationBrandingPoliciesParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization branding policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationBrandingPoliciesParams) WithDefaults() *GetOrganizationBrandingPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization branding policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationBrandingPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) WithTimeout(timeout time.Duration) *GetOrganizationBrandingPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) WithContext(ctx context.Context) *GetOrganizationBrandingPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) WithHTTPClient(client *http.Client) *GetOrganizationBrandingPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) WithOrganizationID(organizationID string) *GetOrganizationBrandingPoliciesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization branding policies params
func (o *GetOrganizationBrandingPoliciesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationBrandingPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

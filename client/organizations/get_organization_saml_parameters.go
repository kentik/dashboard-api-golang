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

// NewGetOrganizationSamlParams creates a new GetOrganizationSamlParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationSamlParams() *GetOrganizationSamlParams {
	return &GetOrganizationSamlParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationSamlParamsWithTimeout creates a new GetOrganizationSamlParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationSamlParamsWithTimeout(timeout time.Duration) *GetOrganizationSamlParams {
	return &GetOrganizationSamlParams{
		timeout: timeout,
	}
}

// NewGetOrganizationSamlParamsWithContext creates a new GetOrganizationSamlParams object
// with the ability to set a context for a request.
func NewGetOrganizationSamlParamsWithContext(ctx context.Context) *GetOrganizationSamlParams {
	return &GetOrganizationSamlParams{
		Context: ctx,
	}
}

// NewGetOrganizationSamlParamsWithHTTPClient creates a new GetOrganizationSamlParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationSamlParamsWithHTTPClient(client *http.Client) *GetOrganizationSamlParams {
	return &GetOrganizationSamlParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationSamlParams contains all the parameters to send to the API endpoint

	for the get organization saml operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationSamlParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization saml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSamlParams) WithDefaults() *GetOrganizationSamlParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization saml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationSamlParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization saml params
func (o *GetOrganizationSamlParams) WithTimeout(timeout time.Duration) *GetOrganizationSamlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization saml params
func (o *GetOrganizationSamlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization saml params
func (o *GetOrganizationSamlParams) WithContext(ctx context.Context) *GetOrganizationSamlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization saml params
func (o *GetOrganizationSamlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization saml params
func (o *GetOrganizationSamlParams) WithHTTPClient(client *http.Client) *GetOrganizationSamlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization saml params
func (o *GetOrganizationSamlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization saml params
func (o *GetOrganizationSamlParams) WithOrganizationID(organizationID string) *GetOrganizationSamlParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization saml params
func (o *GetOrganizationSamlParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationSamlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

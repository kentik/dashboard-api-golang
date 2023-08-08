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

// NewGetOrganizationConfigTemplatesParams creates a new GetOrganizationConfigTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationConfigTemplatesParams() *GetOrganizationConfigTemplatesParams {
	return &GetOrganizationConfigTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationConfigTemplatesParamsWithTimeout creates a new GetOrganizationConfigTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationConfigTemplatesParamsWithTimeout(timeout time.Duration) *GetOrganizationConfigTemplatesParams {
	return &GetOrganizationConfigTemplatesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationConfigTemplatesParamsWithContext creates a new GetOrganizationConfigTemplatesParams object
// with the ability to set a context for a request.
func NewGetOrganizationConfigTemplatesParamsWithContext(ctx context.Context) *GetOrganizationConfigTemplatesParams {
	return &GetOrganizationConfigTemplatesParams{
		Context: ctx,
	}
}

// NewGetOrganizationConfigTemplatesParamsWithHTTPClient creates a new GetOrganizationConfigTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationConfigTemplatesParamsWithHTTPClient(client *http.Client) *GetOrganizationConfigTemplatesParams {
	return &GetOrganizationConfigTemplatesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationConfigTemplatesParams contains all the parameters to send to the API endpoint

	for the get organization config templates operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationConfigTemplatesParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization config templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationConfigTemplatesParams) WithDefaults() *GetOrganizationConfigTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization config templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationConfigTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) WithTimeout(timeout time.Duration) *GetOrganizationConfigTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) WithContext(ctx context.Context) *GetOrganizationConfigTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) WithHTTPClient(client *http.Client) *GetOrganizationConfigTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) WithOrganizationID(organizationID string) *GetOrganizationConfigTemplatesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization config templates params
func (o *GetOrganizationConfigTemplatesParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationConfigTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

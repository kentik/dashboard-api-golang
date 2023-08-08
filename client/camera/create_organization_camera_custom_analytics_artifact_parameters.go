// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// NewCreateOrganizationCameraCustomAnalyticsArtifactParams creates a new CreateOrganizationCameraCustomAnalyticsArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationCameraCustomAnalyticsArtifactParams() *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	return &CreateOrganizationCameraCustomAnalyticsArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithTimeout creates a new CreateOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithTimeout(timeout time.Duration) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	return &CreateOrganizationCameraCustomAnalyticsArtifactParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithContext creates a new CreateOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a context for a request.
func NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithContext(ctx context.Context) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	return &CreateOrganizationCameraCustomAnalyticsArtifactParams{
		Context: ctx,
	}
}

// NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithHTTPClient creates a new CreateOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationCameraCustomAnalyticsArtifactParamsWithHTTPClient(client *http.Client) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	return &CreateOrganizationCameraCustomAnalyticsArtifactParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationCameraCustomAnalyticsArtifactParams contains all the parameters to send to the API endpoint

	for the create organization camera custom analytics artifact operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationCameraCustomAnalyticsArtifactParams struct {

	// CreateOrganizationCameraCustomAnalyticsArtifact.
	CreateOrganizationCameraCustomAnalyticsArtifact CreateOrganizationCameraCustomAnalyticsArtifactBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization camera custom analytics artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithDefaults() *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization camera custom analytics artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithTimeout(timeout time.Duration) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithContext(ctx context.Context) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithHTTPClient(client *http.Client) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationCameraCustomAnalyticsArtifact adds the createOrganizationCameraCustomAnalyticsArtifact to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithCreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact CreateOrganizationCameraCustomAnalyticsArtifactBody) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetCreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact)
	return o
}

// SetCreateOrganizationCameraCustomAnalyticsArtifact adds the createOrganizationCameraCustomAnalyticsArtifact to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetCreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact CreateOrganizationCameraCustomAnalyticsArtifactBody) {
	o.CreateOrganizationCameraCustomAnalyticsArtifact = createOrganizationCameraCustomAnalyticsArtifact
}

// WithOrganizationID adds the organizationID to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WithOrganizationID(organizationID string) *CreateOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization camera custom analytics artifact params
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationCameraCustomAnalyticsArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationCameraCustomAnalyticsArtifact); err != nil {
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

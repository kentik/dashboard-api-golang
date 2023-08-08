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

// NewGetOrganizationCameraCustomAnalyticsArtifactParams creates a new GetOrganizationCameraCustomAnalyticsArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationCameraCustomAnalyticsArtifactParams() *GetOrganizationCameraCustomAnalyticsArtifactParams {
	return &GetOrganizationCameraCustomAnalyticsArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithTimeout creates a new GetOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithTimeout(timeout time.Duration) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	return &GetOrganizationCameraCustomAnalyticsArtifactParams{
		timeout: timeout,
	}
}

// NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithContext creates a new GetOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a context for a request.
func NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithContext(ctx context.Context) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	return &GetOrganizationCameraCustomAnalyticsArtifactParams{
		Context: ctx,
	}
}

// NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithHTTPClient creates a new GetOrganizationCameraCustomAnalyticsArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationCameraCustomAnalyticsArtifactParamsWithHTTPClient(client *http.Client) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	return &GetOrganizationCameraCustomAnalyticsArtifactParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationCameraCustomAnalyticsArtifactParams contains all the parameters to send to the API endpoint

	for the get organization camera custom analytics artifact operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationCameraCustomAnalyticsArtifactParams struct {

	/* ArtifactID.

	   Artifact ID
	*/
	ArtifactID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization camera custom analytics artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithDefaults() *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization camera custom analytics artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithTimeout(timeout time.Duration) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithContext(ctx context.Context) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithHTTPClient(client *http.Client) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithArtifactID(artifactID string) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WithOrganizationID adds the organizationID to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WithOrganizationID(organizationID string) *GetOrganizationCameraCustomAnalyticsArtifactParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization camera custom analytics artifact params
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationCameraCustomAnalyticsArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifactId
	if err := r.SetPathParam("artifactId", o.ArtifactID); err != nil {
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

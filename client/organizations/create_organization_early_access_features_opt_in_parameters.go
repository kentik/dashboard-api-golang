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

// NewCreateOrganizationEarlyAccessFeaturesOptInParams creates a new CreateOrganizationEarlyAccessFeaturesOptInParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationEarlyAccessFeaturesOptInParams() *CreateOrganizationEarlyAccessFeaturesOptInParams {
	return &CreateOrganizationEarlyAccessFeaturesOptInParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithTimeout creates a new CreateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithTimeout(timeout time.Duration) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	return &CreateOrganizationEarlyAccessFeaturesOptInParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithContext creates a new CreateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a context for a request.
func NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithContext(ctx context.Context) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	return &CreateOrganizationEarlyAccessFeaturesOptInParams{
		Context: ctx,
	}
}

// NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient creates a new CreateOrganizationEarlyAccessFeaturesOptInParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationEarlyAccessFeaturesOptInParamsWithHTTPClient(client *http.Client) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	return &CreateOrganizationEarlyAccessFeaturesOptInParams{
		HTTPClient: client,
	}
}

/*
CreateOrganizationEarlyAccessFeaturesOptInParams contains all the parameters to send to the API endpoint

	for the create organization early access features opt in operation.

	Typically these are written to a http.Request.
*/
type CreateOrganizationEarlyAccessFeaturesOptInParams struct {

	// CreateOrganizationEarlyAccessFeaturesOptIn.
	CreateOrganizationEarlyAccessFeaturesOptIn CreateOrganizationEarlyAccessFeaturesOptInBody

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithDefaults() *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization early access features opt in params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithTimeout(timeout time.Duration) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithContext(ctx context.Context) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithHTTPClient(client *http.Client) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationEarlyAccessFeaturesOptIn adds the createOrganizationEarlyAccessFeaturesOptIn to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithCreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn CreateOrganizationEarlyAccessFeaturesOptInBody) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetCreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn)
	return o
}

// SetCreateOrganizationEarlyAccessFeaturesOptIn adds the createOrganizationEarlyAccessFeaturesOptIn to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetCreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn CreateOrganizationEarlyAccessFeaturesOptInBody) {
	o.CreateOrganizationEarlyAccessFeaturesOptIn = createOrganizationEarlyAccessFeaturesOptIn
}

// WithOrganizationID adds the organizationID to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WithOrganizationID(organizationID string) *CreateOrganizationEarlyAccessFeaturesOptInParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization early access features opt in params
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationEarlyAccessFeaturesOptInParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationEarlyAccessFeaturesOptIn); err != nil {
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

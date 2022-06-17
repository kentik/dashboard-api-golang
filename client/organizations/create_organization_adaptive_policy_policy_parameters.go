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

// NewCreateOrganizationAdaptivePolicyPolicyParams creates a new CreateOrganizationAdaptivePolicyPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrganizationAdaptivePolicyPolicyParams() *CreateOrganizationAdaptivePolicyPolicyParams {
	return &CreateOrganizationAdaptivePolicyPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationAdaptivePolicyPolicyParamsWithTimeout creates a new CreateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateOrganizationAdaptivePolicyPolicyParamsWithTimeout(timeout time.Duration) *CreateOrganizationAdaptivePolicyPolicyParams {
	return &CreateOrganizationAdaptivePolicyPolicyParams{
		timeout: timeout,
	}
}

// NewCreateOrganizationAdaptivePolicyPolicyParamsWithContext creates a new CreateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a context for a request.
func NewCreateOrganizationAdaptivePolicyPolicyParamsWithContext(ctx context.Context) *CreateOrganizationAdaptivePolicyPolicyParams {
	return &CreateOrganizationAdaptivePolicyPolicyParams{
		Context: ctx,
	}
}

// NewCreateOrganizationAdaptivePolicyPolicyParamsWithHTTPClient creates a new CreateOrganizationAdaptivePolicyPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrganizationAdaptivePolicyPolicyParamsWithHTTPClient(client *http.Client) *CreateOrganizationAdaptivePolicyPolicyParams {
	return &CreateOrganizationAdaptivePolicyPolicyParams{
		HTTPClient: client,
	}
}

/* CreateOrganizationAdaptivePolicyPolicyParams contains all the parameters to send to the API endpoint
   for the create organization adaptive policy policy operation.

   Typically these are written to a http.Request.
*/
type CreateOrganizationAdaptivePolicyPolicyParams struct {

	// CreateOrganizationAdaptivePolicyPolicy.
	CreateOrganizationAdaptivePolicyPolicy CreateOrganizationAdaptivePolicyPolicyBody

	// OrganizationID.
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithDefaults() *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create organization adaptive policy policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithTimeout(timeout time.Duration) *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithContext(ctx context.Context) *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithHTTPClient(client *http.Client) *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateOrganizationAdaptivePolicyPolicy adds the createOrganizationAdaptivePolicyPolicy to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithCreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy CreateOrganizationAdaptivePolicyPolicyBody) *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetCreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy)
	return o
}

// SetCreateOrganizationAdaptivePolicyPolicy adds the createOrganizationAdaptivePolicyPolicy to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetCreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy CreateOrganizationAdaptivePolicyPolicyBody) {
	o.CreateOrganizationAdaptivePolicyPolicy = createOrganizationAdaptivePolicyPolicy
}

// WithOrganizationID adds the organizationID to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WithOrganizationID(organizationID string) *CreateOrganizationAdaptivePolicyPolicyParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization adaptive policy policy params
func (o *CreateOrganizationAdaptivePolicyPolicyParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationAdaptivePolicyPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.CreateOrganizationAdaptivePolicyPolicy); err != nil {
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
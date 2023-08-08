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

// NewUpdateOrganizationPolicyObjectParams creates a new UpdateOrganizationPolicyObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationPolicyObjectParams() *UpdateOrganizationPolicyObjectParams {
	return &UpdateOrganizationPolicyObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationPolicyObjectParamsWithTimeout creates a new UpdateOrganizationPolicyObjectParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationPolicyObjectParamsWithTimeout(timeout time.Duration) *UpdateOrganizationPolicyObjectParams {
	return &UpdateOrganizationPolicyObjectParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationPolicyObjectParamsWithContext creates a new UpdateOrganizationPolicyObjectParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationPolicyObjectParamsWithContext(ctx context.Context) *UpdateOrganizationPolicyObjectParams {
	return &UpdateOrganizationPolicyObjectParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationPolicyObjectParamsWithHTTPClient creates a new UpdateOrganizationPolicyObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationPolicyObjectParamsWithHTTPClient(client *http.Client) *UpdateOrganizationPolicyObjectParams {
	return &UpdateOrganizationPolicyObjectParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationPolicyObjectParams contains all the parameters to send to the API endpoint

	for the update organization policy object operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationPolicyObjectParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PolicyObjectID.

	   Policy object ID
	*/
	PolicyObjectID string

	// UpdateOrganizationPolicyObject.
	UpdateOrganizationPolicyObject UpdateOrganizationPolicyObjectBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationPolicyObjectParams) WithDefaults() *UpdateOrganizationPolicyObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization policy object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationPolicyObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithTimeout(timeout time.Duration) *UpdateOrganizationPolicyObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithContext(ctx context.Context) *UpdateOrganizationPolicyObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithHTTPClient(client *http.Client) *UpdateOrganizationPolicyObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithOrganizationID(organizationID string) *UpdateOrganizationPolicyObjectParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPolicyObjectID adds the policyObjectID to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithPolicyObjectID(policyObjectID string) *UpdateOrganizationPolicyObjectParams {
	o.SetPolicyObjectID(policyObjectID)
	return o
}

// SetPolicyObjectID adds the policyObjectId to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetPolicyObjectID(policyObjectID string) {
	o.PolicyObjectID = policyObjectID
}

// WithUpdateOrganizationPolicyObject adds the updateOrganizationPolicyObject to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) WithUpdateOrganizationPolicyObject(updateOrganizationPolicyObject UpdateOrganizationPolicyObjectBody) *UpdateOrganizationPolicyObjectParams {
	o.SetUpdateOrganizationPolicyObject(updateOrganizationPolicyObject)
	return o
}

// SetUpdateOrganizationPolicyObject adds the updateOrganizationPolicyObject to the update organization policy object params
func (o *UpdateOrganizationPolicyObjectParams) SetUpdateOrganizationPolicyObject(updateOrganizationPolicyObject UpdateOrganizationPolicyObjectBody) {
	o.UpdateOrganizationPolicyObject = updateOrganizationPolicyObject
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationPolicyObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param policyObjectId
	if err := r.SetPathParam("policyObjectId", o.PolicyObjectID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationPolicyObject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

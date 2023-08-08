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

// NewUpdateOrganizationAdaptivePolicyGroupParams creates a new UpdateOrganizationAdaptivePolicyGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrganizationAdaptivePolicyGroupParams() *UpdateOrganizationAdaptivePolicyGroupParams {
	return &UpdateOrganizationAdaptivePolicyGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrganizationAdaptivePolicyGroupParamsWithTimeout creates a new UpdateOrganizationAdaptivePolicyGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateOrganizationAdaptivePolicyGroupParamsWithTimeout(timeout time.Duration) *UpdateOrganizationAdaptivePolicyGroupParams {
	return &UpdateOrganizationAdaptivePolicyGroupParams{
		timeout: timeout,
	}
}

// NewUpdateOrganizationAdaptivePolicyGroupParamsWithContext creates a new UpdateOrganizationAdaptivePolicyGroupParams object
// with the ability to set a context for a request.
func NewUpdateOrganizationAdaptivePolicyGroupParamsWithContext(ctx context.Context) *UpdateOrganizationAdaptivePolicyGroupParams {
	return &UpdateOrganizationAdaptivePolicyGroupParams{
		Context: ctx,
	}
}

// NewUpdateOrganizationAdaptivePolicyGroupParamsWithHTTPClient creates a new UpdateOrganizationAdaptivePolicyGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrganizationAdaptivePolicyGroupParamsWithHTTPClient(client *http.Client) *UpdateOrganizationAdaptivePolicyGroupParams {
	return &UpdateOrganizationAdaptivePolicyGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateOrganizationAdaptivePolicyGroupParams contains all the parameters to send to the API endpoint

	for the update organization adaptive policy group operation.

	Typically these are written to a http.Request.
*/
type UpdateOrganizationAdaptivePolicyGroupParams struct {

	/* ID.

	   ID
	*/
	ID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	// UpdateOrganizationAdaptivePolicyGroup.
	UpdateOrganizationAdaptivePolicyGroup UpdateOrganizationAdaptivePolicyGroupBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithDefaults() *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithTimeout(timeout time.Duration) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithContext(ctx context.Context) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithHTTPClient(client *http.Client) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithID(id string) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithOrganizationID(organizationID string) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUpdateOrganizationAdaptivePolicyGroup adds the updateOrganizationAdaptivePolicyGroup to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WithUpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup UpdateOrganizationAdaptivePolicyGroupBody) *UpdateOrganizationAdaptivePolicyGroupParams {
	o.SetUpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup)
	return o
}

// SetUpdateOrganizationAdaptivePolicyGroup adds the updateOrganizationAdaptivePolicyGroup to the update organization adaptive policy group params
func (o *UpdateOrganizationAdaptivePolicyGroupParams) SetUpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup UpdateOrganizationAdaptivePolicyGroupBody) {
	o.UpdateOrganizationAdaptivePolicyGroup = updateOrganizationAdaptivePolicyGroup
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrganizationAdaptivePolicyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateOrganizationAdaptivePolicyGroup); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

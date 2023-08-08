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

// NewDeleteOrganizationAdaptivePolicyGroupParams creates a new DeleteOrganizationAdaptivePolicyGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationAdaptivePolicyGroupParams() *DeleteOrganizationAdaptivePolicyGroupParams {
	return &DeleteOrganizationAdaptivePolicyGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationAdaptivePolicyGroupParamsWithTimeout creates a new DeleteOrganizationAdaptivePolicyGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationAdaptivePolicyGroupParamsWithTimeout(timeout time.Duration) *DeleteOrganizationAdaptivePolicyGroupParams {
	return &DeleteOrganizationAdaptivePolicyGroupParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationAdaptivePolicyGroupParamsWithContext creates a new DeleteOrganizationAdaptivePolicyGroupParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationAdaptivePolicyGroupParamsWithContext(ctx context.Context) *DeleteOrganizationAdaptivePolicyGroupParams {
	return &DeleteOrganizationAdaptivePolicyGroupParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationAdaptivePolicyGroupParamsWithHTTPClient creates a new DeleteOrganizationAdaptivePolicyGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationAdaptivePolicyGroupParamsWithHTTPClient(client *http.Client) *DeleteOrganizationAdaptivePolicyGroupParams {
	return &DeleteOrganizationAdaptivePolicyGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationAdaptivePolicyGroupParams contains all the parameters to send to the API endpoint

	for the delete organization adaptive policy group operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationAdaptivePolicyGroupParams struct {

	/* ID.

	   ID
	*/
	ID string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithDefaults() *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization adaptive policy group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithTimeout(timeout time.Duration) *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithContext(ctx context.Context) *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithHTTPClient(client *http.Client) *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithID(id string) *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetID(id string) {
	o.ID = id
}

// WithOrganizationID adds the organizationID to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WithOrganizationID(organizationID string) *DeleteOrganizationAdaptivePolicyGroupParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization adaptive policy group params
func (o *DeleteOrganizationAdaptivePolicyGroupParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationAdaptivePolicyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

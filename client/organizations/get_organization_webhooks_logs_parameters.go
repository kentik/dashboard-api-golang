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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationWebhooksLogsParams creates a new GetOrganizationWebhooksLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationWebhooksLogsParams() *GetOrganizationWebhooksLogsParams {
	return &GetOrganizationWebhooksLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationWebhooksLogsParamsWithTimeout creates a new GetOrganizationWebhooksLogsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationWebhooksLogsParamsWithTimeout(timeout time.Duration) *GetOrganizationWebhooksLogsParams {
	return &GetOrganizationWebhooksLogsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationWebhooksLogsParamsWithContext creates a new GetOrganizationWebhooksLogsParams object
// with the ability to set a context for a request.
func NewGetOrganizationWebhooksLogsParamsWithContext(ctx context.Context) *GetOrganizationWebhooksLogsParams {
	return &GetOrganizationWebhooksLogsParams{
		Context: ctx,
	}
}

// NewGetOrganizationWebhooksLogsParamsWithHTTPClient creates a new GetOrganizationWebhooksLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationWebhooksLogsParamsWithHTTPClient(client *http.Client) *GetOrganizationWebhooksLogsParams {
	return &GetOrganizationWebhooksLogsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationWebhooksLogsParams contains all the parameters to send to the API endpoint

	for the get organization webhooks logs operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationWebhooksLogsParams struct {

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	*/
	PerPage *int64

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

	   Format: float
	*/
	Timespan *float32

	/* URL.

	   The URL the webhook was sent to
	*/
	URL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization webhooks logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationWebhooksLogsParams) WithDefaults() *GetOrganizationWebhooksLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization webhooks logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationWebhooksLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithTimeout(timeout time.Duration) *GetOrganizationWebhooksLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithContext(ctx context.Context) *GetOrganizationWebhooksLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithHTTPClient(client *http.Client) *GetOrganizationWebhooksLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithEndingBefore(endingBefore *string) *GetOrganizationWebhooksLogsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithOrganizationID adds the organizationID to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithOrganizationID(organizationID string) *GetOrganizationWebhooksLogsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPerPage adds the perPage to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithPerPage(perPage *int64) *GetOrganizationWebhooksLogsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithStartingAfter adds the startingAfter to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithStartingAfter(startingAfter *string) *GetOrganizationWebhooksLogsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithT0(t0 *string) *GetOrganizationWebhooksLogsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithT1(t1 *string) *GetOrganizationWebhooksLogsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithTimespan(timespan *float32) *GetOrganizationWebhooksLogsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WithURL adds the url to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) WithURL(url *string) *GetOrganizationWebhooksLogsParams {
	o.SetURL(url)
	return o
}

// SetURL adds the url to the get organization webhooks logs params
func (o *GetOrganizationWebhooksLogsParams) SetURL(url *string) {
	o.URL = url
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationWebhooksLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param endingBefore
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("endingBefore", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param startingAfter
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("startingAfter", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if o.T0 != nil {

		// query param t0
		var qrT0 string

		if o.T0 != nil {
			qrT0 = *o.T0
		}
		qT0 := qrT0
		if qT0 != "" {

			if err := r.SetQueryParam("t0", qT0); err != nil {
				return err
			}
		}
	}

	if o.T1 != nil {

		// query param t1
		var qrT1 string

		if o.T1 != nil {
			qrT1 = *o.T1
		}
		qT1 := qrT1
		if qT1 != "" {

			if err := r.SetQueryParam("t1", qT1); err != nil {
				return err
			}
		}
	}

	if o.Timespan != nil {

		// query param timespan
		var qrTimespan float32

		if o.Timespan != nil {
			qrTimespan = *o.Timespan
		}
		qTimespan := swag.FormatFloat32(qrTimespan)
		if qTimespan != "" {

			if err := r.SetQueryParam("timespan", qTimespan); err != nil {
				return err
			}
		}
	}

	if o.URL != nil {

		// query param url
		var qrURL string

		if o.URL != nil {
			qrURL = *o.URL
		}
		qURL := qrURL
		if qURL != "" {

			if err := r.SetQueryParam("url", qURL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

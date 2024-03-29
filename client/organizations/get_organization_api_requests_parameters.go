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

// NewGetOrganizationAPIRequestsParams creates a new GetOrganizationAPIRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationAPIRequestsParams() *GetOrganizationAPIRequestsParams {
	return &GetOrganizationAPIRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationAPIRequestsParamsWithTimeout creates a new GetOrganizationAPIRequestsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationAPIRequestsParamsWithTimeout(timeout time.Duration) *GetOrganizationAPIRequestsParams {
	return &GetOrganizationAPIRequestsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationAPIRequestsParamsWithContext creates a new GetOrganizationAPIRequestsParams object
// with the ability to set a context for a request.
func NewGetOrganizationAPIRequestsParamsWithContext(ctx context.Context) *GetOrganizationAPIRequestsParams {
	return &GetOrganizationAPIRequestsParams{
		Context: ctx,
	}
}

// NewGetOrganizationAPIRequestsParamsWithHTTPClient creates a new GetOrganizationAPIRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationAPIRequestsParamsWithHTTPClient(client *http.Client) *GetOrganizationAPIRequestsParams {
	return &GetOrganizationAPIRequestsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationAPIRequestsParams contains all the parameters to send to the API endpoint

	for the get organization Api requests operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationAPIRequestsParams struct {

	/* AdminID.

	   Filter the results by the ID of the admin who made the API requests
	*/
	AdminID *string

	/* EndingBefore.

	   A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	EndingBefore *string

	/* Method.

	   Filter the results by the method of the API requests (must be 'GET', 'PUT', 'POST' or 'DELETE')
	*/
	Method *string

	/* OperationIds.

	   Filter the results by one or more operation IDs for the API request
	*/
	OperationIds []string

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID string

	/* Path.

	   Filter the results by the path of the API requests
	*/
	Path *string

	/* PerPage.

	   The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
	*/
	PerPage *int64

	/* ResponseCode.

	   Filter the results by the response code of the API requests
	*/
	ResponseCode *int64

	/* SourceIP.

	   Filter the results by the IP address of the originating API request
	*/
	SourceIP *string

	/* StartingAfter.

	   A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
	*/
	StartingAfter *string

	/* T0.

	   The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
	*/
	T0 *string

	/* T1.

	   The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
	*/
	T1 *string

	/* Timespan.

	   The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.

	   Format: float
	*/
	Timespan *float32

	/* UserAgent.

	   Filter the results by the user agent string of the API request
	*/
	UserAgent *string

	/* Version.

	   Filter the results by the API version of the API request
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization Api requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAPIRequestsParams) WithDefaults() *GetOrganizationAPIRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization Api requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationAPIRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithTimeout(timeout time.Duration) *GetOrganizationAPIRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithContext(ctx context.Context) *GetOrganizationAPIRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithHTTPClient(client *http.Client) *GetOrganizationAPIRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminID adds the adminID to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithAdminID(adminID *string) *GetOrganizationAPIRequestsParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetAdminID(adminID *string) {
	o.AdminID = adminID
}

// WithEndingBefore adds the endingBefore to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithEndingBefore(endingBefore *string) *GetOrganizationAPIRequestsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithMethod adds the method to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithMethod(method *string) *GetOrganizationAPIRequestsParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetMethod(method *string) {
	o.Method = method
}

// WithOperationIds adds the operationIds to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithOperationIds(operationIds []string) *GetOrganizationAPIRequestsParams {
	o.SetOperationIds(operationIds)
	return o
}

// SetOperationIds adds the operationIds to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetOperationIds(operationIds []string) {
	o.OperationIds = operationIds
}

// WithOrganizationID adds the organizationID to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithOrganizationID(organizationID string) *GetOrganizationAPIRequestsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPath adds the path to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithPath(path *string) *GetOrganizationAPIRequestsParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetPath(path *string) {
	o.Path = path
}

// WithPerPage adds the perPage to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithPerPage(perPage *int64) *GetOrganizationAPIRequestsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithResponseCode adds the responseCode to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithResponseCode(responseCode *int64) *GetOrganizationAPIRequestsParams {
	o.SetResponseCode(responseCode)
	return o
}

// SetResponseCode adds the responseCode to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetResponseCode(responseCode *int64) {
	o.ResponseCode = responseCode
}

// WithSourceIP adds the sourceIP to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithSourceIP(sourceIP *string) *GetOrganizationAPIRequestsParams {
	o.SetSourceIP(sourceIP)
	return o
}

// SetSourceIP adds the sourceIp to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetSourceIP(sourceIP *string) {
	o.SourceIP = sourceIP
}

// WithStartingAfter adds the startingAfter to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithStartingAfter(startingAfter *string) *GetOrganizationAPIRequestsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WithT0 adds the t0 to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithT0(t0 *string) *GetOrganizationAPIRequestsParams {
	o.SetT0(t0)
	return o
}

// SetT0 adds the t0 to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetT0(t0 *string) {
	o.T0 = t0
}

// WithT1 adds the t1 to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithT1(t1 *string) *GetOrganizationAPIRequestsParams {
	o.SetT1(t1)
	return o
}

// SetT1 adds the t1 to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetT1(t1 *string) {
	o.T1 = t1
}

// WithTimespan adds the timespan to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithTimespan(timespan *float32) *GetOrganizationAPIRequestsParams {
	o.SetTimespan(timespan)
	return o
}

// SetTimespan adds the timespan to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetTimespan(timespan *float32) {
	o.Timespan = timespan
}

// WithUserAgent adds the userAgent to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithUserAgent(userAgent *string) *GetOrganizationAPIRequestsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithVersion adds the version to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) WithVersion(version *int64) *GetOrganizationAPIRequestsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get organization Api requests params
func (o *GetOrganizationAPIRequestsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationAPIRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdminID != nil {

		// query param adminId
		var qrAdminID string

		if o.AdminID != nil {
			qrAdminID = *o.AdminID
		}
		qAdminID := qrAdminID
		if qAdminID != "" {

			if err := r.SetQueryParam("adminId", qAdminID); err != nil {
				return err
			}
		}
	}

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

	if o.Method != nil {

		// query param method
		var qrMethod string

		if o.Method != nil {
			qrMethod = *o.Method
		}
		qMethod := qrMethod
		if qMethod != "" {

			if err := r.SetQueryParam("method", qMethod); err != nil {
				return err
			}
		}
	}

	if o.OperationIds != nil {

		// binding items for operationIds
		joinedOperationIds := o.bindParamOperationIds(reg)

		// query array param operationIds
		if err := r.SetQueryParam("operationIds", joinedOperationIds...); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
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

	if o.ResponseCode != nil {

		// query param responseCode
		var qrResponseCode int64

		if o.ResponseCode != nil {
			qrResponseCode = *o.ResponseCode
		}
		qResponseCode := swag.FormatInt64(qrResponseCode)
		if qResponseCode != "" {

			if err := r.SetQueryParam("responseCode", qResponseCode); err != nil {
				return err
			}
		}
	}

	if o.SourceIP != nil {

		// query param sourceIp
		var qrSourceIP string

		if o.SourceIP != nil {
			qrSourceIP = *o.SourceIP
		}
		qSourceIP := qrSourceIP
		if qSourceIP != "" {

			if err := r.SetQueryParam("sourceIp", qSourceIP); err != nil {
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

	if o.UserAgent != nil {

		// query param userAgent
		var qrUserAgent string

		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {

			if err := r.SetQueryParam("userAgent", qUserAgent); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationAPIRequests binds the parameter operationIds
func (o *GetOrganizationAPIRequestsParams) bindParamOperationIds(formats strfmt.Registry) []string {
	operationIdsIR := o.OperationIds

	var operationIdsIC []string
	for _, operationIdsIIR := range operationIdsIR { // explode []string

		operationIdsIIV := operationIdsIIR // string as string
		operationIdsIC = append(operationIdsIC, operationIdsIIV)
	}

	// items.CollectionFormat: ""
	operationIdsIS := swag.JoinByFormat(operationIdsIC, "")

	return operationIdsIS
}

// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetOrganizationAdaptivePolicyOverviewReader is a Reader for the GetOrganizationAdaptivePolicyOverview structure.
type GetOrganizationAdaptivePolicyOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdaptivePolicyOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdaptivePolicyOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/adaptivePolicy/overview] getOrganizationAdaptivePolicyOverview", response, response.Code())
	}
}

// NewGetOrganizationAdaptivePolicyOverviewOK creates a GetOrganizationAdaptivePolicyOverviewOK with default headers values
func NewGetOrganizationAdaptivePolicyOverviewOK() *GetOrganizationAdaptivePolicyOverviewOK {
	return &GetOrganizationAdaptivePolicyOverviewOK{}
}

/*
GetOrganizationAdaptivePolicyOverviewOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdaptivePolicyOverviewOK struct {
	Payload *GetOrganizationAdaptivePolicyOverviewOKBody
}

// IsSuccess returns true when this get organization adaptive policy overview o k response has a 2xx status code
func (o *GetOrganizationAdaptivePolicyOverviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization adaptive policy overview o k response has a 3xx status code
func (o *GetOrganizationAdaptivePolicyOverviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization adaptive policy overview o k response has a 4xx status code
func (o *GetOrganizationAdaptivePolicyOverviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization adaptive policy overview o k response has a 5xx status code
func (o *GetOrganizationAdaptivePolicyOverviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization adaptive policy overview o k response a status code equal to that given
func (o *GetOrganizationAdaptivePolicyOverviewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization adaptive policy overview o k response
func (o *GetOrganizationAdaptivePolicyOverviewOK) Code() int {
	return 200
}

func (o *GetOrganizationAdaptivePolicyOverviewOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/overview][%d] getOrganizationAdaptivePolicyOverviewOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyOverviewOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/overview][%d] getOrganizationAdaptivePolicyOverviewOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyOverviewOK) GetPayload() *GetOrganizationAdaptivePolicyOverviewOKBody {
	return o.Payload
}

func (o *GetOrganizationAdaptivePolicyOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationAdaptivePolicyOverviewOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationAdaptivePolicyOverviewOKBody get organization adaptive policy overview o k body
swagger:model GetOrganizationAdaptivePolicyOverviewOKBody
*/
type GetOrganizationAdaptivePolicyOverviewOKBody struct {

	// counts
	Counts *GetOrganizationAdaptivePolicyOverviewOKBodyCounts `json:"counts,omitempty"`

	// limits
	Limits *GetOrganizationAdaptivePolicyOverviewOKBodyLimits `json:"limits,omitempty"`
}

// Validate validates this get organization adaptive policy overview o k body
func (o *GetOrganizationAdaptivePolicyOverviewOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationAdaptivePolicyOverviewOKBody) validateCounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Counts) { // not required
		return nil
	}

	if o.Counts != nil {
		if err := o.Counts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationAdaptivePolicyOverviewOKBody) validateLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.Limits) { // not required
		return nil
	}

	if o.Limits != nil {
		if err := o.Limits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "limits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization adaptive policy overview o k body based on the context it is used
func (o *GetOrganizationAdaptivePolicyOverviewOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationAdaptivePolicyOverviewOKBody) contextValidateCounts(ctx context.Context, formats strfmt.Registry) error {

	if o.Counts != nil {

		if swag.IsZero(o.Counts) { // not required
			return nil
		}

		if err := o.Counts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationAdaptivePolicyOverviewOKBody) contextValidateLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.Limits != nil {

		if swag.IsZero(o.Limits) { // not required
			return nil
		}

		if err := o.Limits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationAdaptivePolicyOverviewOK" + "." + "limits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdaptivePolicyOverviewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationAdaptivePolicyOverviewOKBodyCounts The current amount of various adaptive policy objects.
swagger:model GetOrganizationAdaptivePolicyOverviewOKBodyCounts
*/
type GetOrganizationAdaptivePolicyOverviewOKBodyCounts struct {

	// Number of adaptive policies currently in the organization that allow all traffic.
	AllowPolicies int64 `json:"allowPolicies,omitempty"`

	// Number of user-created adaptive policy ACLs currently in the organization.
	CustomAcls int64 `json:"customAcls,omitempty"`

	// Number of user-created adaptive policy groups currently in the organization.
	CustomGroups int64 `json:"customGroups,omitempty"`

	// Number of adaptive policies currently in the organization that deny all traffic.
	DenyPolicies int64 `json:"denyPolicies,omitempty"`

	// Number of adaptive policy groups currently in the organization.
	Groups int64 `json:"groups,omitempty"`

	// Number of adaptive policies currently in the organization.
	Policies int64 `json:"policies,omitempty"`

	// Number of policy objects (with the adaptive policy type) currently in the organization.
	PolicyObjects int64 `json:"policyObjects,omitempty"`
}

// Validate validates this get organization adaptive policy overview o k body counts
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyCounts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization adaptive policy overview o k body counts based on context it is used
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyCounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyCounts) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyCounts) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdaptivePolicyOverviewOKBodyCounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationAdaptivePolicyOverviewOKBodyLimits The current limits of various adaptive policy objects.
swagger:model GetOrganizationAdaptivePolicyOverviewOKBodyLimits
*/
type GetOrganizationAdaptivePolicyOverviewOKBodyLimits struct {

	// Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.
	AclsInAPolicy int64 `json:"aclsInAPolicy,omitempty"`

	// Maximum number of user-created adaptive policy groups allowed in the organization.
	CustomGroups int64 `json:"customGroups,omitempty"`

	// Maximum number of policy objects (with the adaptive policy type) allowed in the organization.
	PolicyObjects int64 `json:"policyObjects,omitempty"`

	// Maximum number of rules allowed in an adaptive policy ACL in the organization.
	RulesInAnACL int64 `json:"rulesInAnAcl,omitempty"`
}

// Validate validates this get organization adaptive policy overview o k body limits
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization adaptive policy overview o k body limits based on context it is used
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdaptivePolicyOverviewOKBodyLimits) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdaptivePolicyOverviewOKBodyLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

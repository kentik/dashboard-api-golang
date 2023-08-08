// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationAdaptivePolicyGroupReader is a Reader for the UpdateOrganizationAdaptivePolicyGroup structure.
type UpdateOrganizationAdaptivePolicyGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationAdaptivePolicyGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationAdaptivePolicyGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /organizations/{organizationId}/adaptivePolicy/groups/{id}] updateOrganizationAdaptivePolicyGroup", response, response.Code())
	}
}

// NewUpdateOrganizationAdaptivePolicyGroupOK creates a UpdateOrganizationAdaptivePolicyGroupOK with default headers values
func NewUpdateOrganizationAdaptivePolicyGroupOK() *UpdateOrganizationAdaptivePolicyGroupOK {
	return &UpdateOrganizationAdaptivePolicyGroupOK{}
}

/*
UpdateOrganizationAdaptivePolicyGroupOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationAdaptivePolicyGroupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization adaptive policy group o k response has a 2xx status code
func (o *UpdateOrganizationAdaptivePolicyGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization adaptive policy group o k response has a 3xx status code
func (o *UpdateOrganizationAdaptivePolicyGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization adaptive policy group o k response has a 4xx status code
func (o *UpdateOrganizationAdaptivePolicyGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization adaptive policy group o k response has a 5xx status code
func (o *UpdateOrganizationAdaptivePolicyGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization adaptive policy group o k response a status code equal to that given
func (o *UpdateOrganizationAdaptivePolicyGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization adaptive policy group o k response
func (o *UpdateOrganizationAdaptivePolicyGroupOK) Code() int {
	return 200
}

func (o *UpdateOrganizationAdaptivePolicyGroupOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/adaptivePolicy/groups/{id}][%d] updateOrganizationAdaptivePolicyGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationAdaptivePolicyGroupOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/adaptivePolicy/groups/{id}][%d] updateOrganizationAdaptivePolicyGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationAdaptivePolicyGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationAdaptivePolicyGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationAdaptivePolicyGroupBody update organization adaptive policy group body
// Example: {"description":"Group of XYZ Corp Employees","isDefaultGroup":false,"name":"Employee Group","policyObjects":[{"id":"2345","name":"Example Policy Object"}],"requiredIpMappings":[],"sgt":1000}
swagger:model UpdateOrganizationAdaptivePolicyGroupBody
*/
type UpdateOrganizationAdaptivePolicyGroupBody struct {

	// Description of the group
	Description string `json:"description,omitempty"`

	// Name of the group
	Name string `json:"name,omitempty"`

	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute)
	PolicyObjects []*UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0 `json:"policyObjects"`

	// SGT value of the group
	Sgt int64 `json:"sgt,omitempty"`
}

// Validate validates this update organization adaptive policy group body
func (o *UpdateOrganizationAdaptivePolicyGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicyObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationAdaptivePolicyGroupBody) validatePolicyObjects(formats strfmt.Registry) error {
	if swag.IsZero(o.PolicyObjects) { // not required
		return nil
	}

	for i := 0; i < len(o.PolicyObjects); i++ {
		if swag.IsZero(o.PolicyObjects[i]) { // not required
			continue
		}

		if o.PolicyObjects[i] != nil {
			if err := o.PolicyObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateOrganizationAdaptivePolicyGroup" + "." + "policyObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateOrganizationAdaptivePolicyGroup" + "." + "policyObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update organization adaptive policy group body based on the context it is used
func (o *UpdateOrganizationAdaptivePolicyGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePolicyObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationAdaptivePolicyGroupBody) contextValidatePolicyObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PolicyObjects); i++ {

		if o.PolicyObjects[i] != nil {

			if swag.IsZero(o.PolicyObjects[i]) { // not required
				return nil
			}

			if err := o.PolicyObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateOrganizationAdaptivePolicyGroup" + "." + "policyObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateOrganizationAdaptivePolicyGroup" + "." + "policyObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyGroupBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0 update organization adaptive policy group params body policy objects items0
swagger:model UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0
*/
type UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0 struct {

	// The ID of the policy object
	ID string `json:"id,omitempty"`

	// The name of the policy object
	Name string `json:"name,omitempty"`
}

// Validate validates this update organization adaptive policy group params body policy objects items0
func (o *UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization adaptive policy group params body policy objects items0 based on context it is used
func (o *UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyGroupParamsBodyPolicyObjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

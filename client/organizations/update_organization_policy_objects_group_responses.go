// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationPolicyObjectsGroupReader is a Reader for the UpdateOrganizationPolicyObjectsGroup structure.
type UpdateOrganizationPolicyObjectsGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationPolicyObjectsGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationPolicyObjectsGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}] updateOrganizationPolicyObjectsGroup", response, response.Code())
	}
}

// NewUpdateOrganizationPolicyObjectsGroupOK creates a UpdateOrganizationPolicyObjectsGroupOK with default headers values
func NewUpdateOrganizationPolicyObjectsGroupOK() *UpdateOrganizationPolicyObjectsGroupOK {
	return &UpdateOrganizationPolicyObjectsGroupOK{}
}

/*
UpdateOrganizationPolicyObjectsGroupOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationPolicyObjectsGroupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization policy objects group o k response has a 2xx status code
func (o *UpdateOrganizationPolicyObjectsGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization policy objects group o k response has a 3xx status code
func (o *UpdateOrganizationPolicyObjectsGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization policy objects group o k response has a 4xx status code
func (o *UpdateOrganizationPolicyObjectsGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization policy objects group o k response has a 5xx status code
func (o *UpdateOrganizationPolicyObjectsGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization policy objects group o k response a status code equal to that given
func (o *UpdateOrganizationPolicyObjectsGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization policy objects group o k response
func (o *UpdateOrganizationPolicyObjectsGroupOK) Code() int {
	return 200
}

func (o *UpdateOrganizationPolicyObjectsGroupOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}][%d] updateOrganizationPolicyObjectsGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationPolicyObjectsGroupOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId}][%d] updateOrganizationPolicyObjectsGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationPolicyObjectsGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationPolicyObjectsGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationPolicyObjectsGroupBody update organization policy objects group body
// Example: {"name":"Web Servers - Datacenter 10","objectIds":[]}
swagger:model UpdateOrganizationPolicyObjectsGroupBody
*/
type UpdateOrganizationPolicyObjectsGroupBody struct {

	// A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name,omitempty"`

	// A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
	ObjectIds []int64 `json:"objectIds"`
}

// Validate validates this update organization policy objects group body
func (o *UpdateOrganizationPolicyObjectsGroupBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization policy objects group body based on context it is used
func (o *UpdateOrganizationPolicyObjectsGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationPolicyObjectsGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationPolicyObjectsGroupBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationPolicyObjectsGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

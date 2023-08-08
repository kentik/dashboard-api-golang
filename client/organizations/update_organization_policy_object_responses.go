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

// UpdateOrganizationPolicyObjectReader is a Reader for the UpdateOrganizationPolicyObject structure.
type UpdateOrganizationPolicyObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationPolicyObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationPolicyObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /organizations/{organizationId}/policyObjects/{policyObjectId}] updateOrganizationPolicyObject", response, response.Code())
	}
}

// NewUpdateOrganizationPolicyObjectOK creates a UpdateOrganizationPolicyObjectOK with default headers values
func NewUpdateOrganizationPolicyObjectOK() *UpdateOrganizationPolicyObjectOK {
	return &UpdateOrganizationPolicyObjectOK{}
}

/*
UpdateOrganizationPolicyObjectOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationPolicyObjectOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization policy object o k response has a 2xx status code
func (o *UpdateOrganizationPolicyObjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization policy object o k response has a 3xx status code
func (o *UpdateOrganizationPolicyObjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization policy object o k response has a 4xx status code
func (o *UpdateOrganizationPolicyObjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization policy object o k response has a 5xx status code
func (o *UpdateOrganizationPolicyObjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization policy object o k response a status code equal to that given
func (o *UpdateOrganizationPolicyObjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization policy object o k response
func (o *UpdateOrganizationPolicyObjectOK) Code() int {
	return 200
}

func (o *UpdateOrganizationPolicyObjectOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] updateOrganizationPolicyObjectOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationPolicyObjectOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] updateOrganizationPolicyObjectOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationPolicyObjectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationPolicyObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationPolicyObjectBody update organization policy object body
// Example: {"groupIds":[],"name":"Web Servers - Datacenter 10","type":"cidr"}
swagger:model UpdateOrganizationPolicyObjectBody
*/
type UpdateOrganizationPolicyObjectBody struct {

	// CIDR Value of a policy object (e.g. 10.11.12.1/24")
	Cidr string `json:"cidr,omitempty"`

	// Fully qualified domain name of policy object (e.g. "example.com")
	Fqdn string `json:"fqdn,omitempty"`

	// The IDs of policy object groups the policy object belongs to
	GroupIds []int64 `json:"groupIds"`

	// IP Address of a policy object (e.g. "1.2.3.4")
	IP string `json:"ip,omitempty"`

	// Mask of a policy object (e.g. "255.255.0.0")
	Mask string `json:"mask,omitempty"`

	// Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name string `json:"name,omitempty"`
}

// Validate validates this update organization policy object body
func (o *UpdateOrganizationPolicyObjectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization policy object body based on context it is used
func (o *UpdateOrganizationPolicyObjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationPolicyObjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationPolicyObjectBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationPolicyObjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
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
	"github.com/go-openapi/validate"
)

// CreateOrganizationPolicyObjectsGroupReader is a Reader for the CreateOrganizationPolicyObjectsGroup structure.
type CreateOrganizationPolicyObjectsGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationPolicyObjectsGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationPolicyObjectsGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/policyObjects/groups] createOrganizationPolicyObjectsGroup", response, response.Code())
	}
}

// NewCreateOrganizationPolicyObjectsGroupCreated creates a CreateOrganizationPolicyObjectsGroupCreated with default headers values
func NewCreateOrganizationPolicyObjectsGroupCreated() *CreateOrganizationPolicyObjectsGroupCreated {
	return &CreateOrganizationPolicyObjectsGroupCreated{}
}

/*
CreateOrganizationPolicyObjectsGroupCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationPolicyObjectsGroupCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization policy objects group created response has a 2xx status code
func (o *CreateOrganizationPolicyObjectsGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization policy objects group created response has a 3xx status code
func (o *CreateOrganizationPolicyObjectsGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization policy objects group created response has a 4xx status code
func (o *CreateOrganizationPolicyObjectsGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization policy objects group created response has a 5xx status code
func (o *CreateOrganizationPolicyObjectsGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization policy objects group created response a status code equal to that given
func (o *CreateOrganizationPolicyObjectsGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization policy objects group created response
func (o *CreateOrganizationPolicyObjectsGroupCreated) Code() int {
	return 201
}

func (o *CreateOrganizationPolicyObjectsGroupCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/policyObjects/groups][%d] createOrganizationPolicyObjectsGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationPolicyObjectsGroupCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/policyObjects/groups][%d] createOrganizationPolicyObjectsGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationPolicyObjectsGroupCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationPolicyObjectsGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationPolicyObjectsGroupBody create organization policy objects group body
// Example: {"name":"Web Servers - Datacenter 10","objectIds":[]}
swagger:model CreateOrganizationPolicyObjectsGroupBody
*/
type CreateOrganizationPolicyObjectsGroupBody struct {

	// Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)
	Category string `json:"category,omitempty"`

	// A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	// Required: true
	Name *string `json:"name"`

	// A list of Policy Object ID's that this NetworkObjectGroup should be associated to (note: these ID's will replace the existing associated Policy Objects)
	ObjectIds []int64 `json:"objectIds"`
}

// Validate validates this create organization policy objects group body
func (o *CreateOrganizationPolicyObjectsGroupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationPolicyObjectsGroupBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationPolicyObjectsGroup"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization policy objects group body based on context it is used
func (o *CreateOrganizationPolicyObjectsGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationPolicyObjectsGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationPolicyObjectsGroupBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationPolicyObjectsGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

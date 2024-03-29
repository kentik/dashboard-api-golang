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

// UpdateOrganizationBrandingPoliciesPrioritiesReader is a Reader for the UpdateOrganizationBrandingPoliciesPriorities structure.
type UpdateOrganizationBrandingPoliciesPrioritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationBrandingPoliciesPrioritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationBrandingPoliciesPrioritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /organizations/{organizationId}/brandingPolicies/priorities] updateOrganizationBrandingPoliciesPriorities", response, response.Code())
	}
}

// NewUpdateOrganizationBrandingPoliciesPrioritiesOK creates a UpdateOrganizationBrandingPoliciesPrioritiesOK with default headers values
func NewUpdateOrganizationBrandingPoliciesPrioritiesOK() *UpdateOrganizationBrandingPoliciesPrioritiesOK {
	return &UpdateOrganizationBrandingPoliciesPrioritiesOK{}
}

/*
UpdateOrganizationBrandingPoliciesPrioritiesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationBrandingPoliciesPrioritiesOK struct {
	Payload *UpdateOrganizationBrandingPoliciesPrioritiesOKBody
}

// IsSuccess returns true when this update organization branding policies priorities o k response has a 2xx status code
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization branding policies priorities o k response has a 3xx status code
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization branding policies priorities o k response has a 4xx status code
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization branding policies priorities o k response has a 5xx status code
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization branding policies priorities o k response a status code equal to that given
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization branding policies priorities o k response
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) Code() int {
	return 200
}

func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/brandingPolicies/priorities][%d] updateOrganizationBrandingPoliciesPrioritiesOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/brandingPolicies/priorities][%d] updateOrganizationBrandingPoliciesPrioritiesOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) GetPayload() *UpdateOrganizationBrandingPoliciesPrioritiesOKBody {
	return o.Payload
}

func (o *UpdateOrganizationBrandingPoliciesPrioritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateOrganizationBrandingPoliciesPrioritiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationBrandingPoliciesPrioritiesBody update organization branding policies priorities body
// Example: {"brandingPolicyIds":["123","456","789"]}
swagger:model UpdateOrganizationBrandingPoliciesPrioritiesBody
*/
type UpdateOrganizationBrandingPoliciesPrioritiesBody struct {

	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies
	//
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

// Validate validates this update organization branding policies priorities body
func (o *UpdateOrganizationBrandingPoliciesPrioritiesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization branding policies priorities body based on context it is used
func (o *UpdateOrganizationBrandingPoliciesPrioritiesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationBrandingPoliciesPrioritiesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationBrandingPoliciesPrioritiesBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationBrandingPoliciesPrioritiesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateOrganizationBrandingPoliciesPrioritiesOKBody update organization branding policies priorities o k body
swagger:model UpdateOrganizationBrandingPoliciesPrioritiesOKBody
*/
type UpdateOrganizationBrandingPoliciesPrioritiesOKBody struct {

	//       An ordered list of branding policy IDs that determines the priority order of how to apply the policies
	//
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

// Validate validates this update organization branding policies priorities o k body
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization branding policies priorities o k body based on context it is used
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationBrandingPoliciesPrioritiesOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationBrandingPoliciesPrioritiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

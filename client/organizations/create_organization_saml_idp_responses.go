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

// CreateOrganizationSamlIdpReader is a Reader for the CreateOrganizationSamlIdp structure.
type CreateOrganizationSamlIdpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationSamlIdpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationSamlIdpCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/saml/idps] createOrganizationSamlIdp", response, response.Code())
	}
}

// NewCreateOrganizationSamlIdpCreated creates a CreateOrganizationSamlIdpCreated with default headers values
func NewCreateOrganizationSamlIdpCreated() *CreateOrganizationSamlIdpCreated {
	return &CreateOrganizationSamlIdpCreated{}
}

/*
CreateOrganizationSamlIdpCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationSamlIdpCreated struct {
	Payload []*CreateOrganizationSamlIdpCreatedBodyItems0
}

// IsSuccess returns true when this create organization saml idp created response has a 2xx status code
func (o *CreateOrganizationSamlIdpCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization saml idp created response has a 3xx status code
func (o *CreateOrganizationSamlIdpCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization saml idp created response has a 4xx status code
func (o *CreateOrganizationSamlIdpCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization saml idp created response has a 5xx status code
func (o *CreateOrganizationSamlIdpCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization saml idp created response a status code equal to that given
func (o *CreateOrganizationSamlIdpCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization saml idp created response
func (o *CreateOrganizationSamlIdpCreated) Code() int {
	return 201
}

func (o *CreateOrganizationSamlIdpCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/saml/idps][%d] createOrganizationSamlIdpCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationSamlIdpCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/saml/idps][%d] createOrganizationSamlIdpCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationSamlIdpCreated) GetPayload() []*CreateOrganizationSamlIdpCreatedBodyItems0 {
	return o.Payload
}

func (o *CreateOrganizationSamlIdpCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationSamlIdpBody create organization saml idp body
// Example: {"sloLogoutUrl":"https://somewhere.com","x509certSha1Fingerprint":"00:11:22:33:44:55:66:77:88:99:00:11:22:33:44:55:66:77:88:99"}
swagger:model CreateOrganizationSamlIdpBody
*/
type CreateOrganizationSamlIdpBody struct {

	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutURL string `json:"sloLogoutUrl,omitempty"`

	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	// Required: true
	X509certSha1Fingerprint *string `json:"x509certSha1Fingerprint"`
}

// Validate validates this create organization saml idp body
func (o *CreateOrganizationSamlIdpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX509certSha1Fingerprint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationSamlIdpBody) validateX509certSha1Fingerprint(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationSamlIdp"+"."+"x509certSha1Fingerprint", "body", o.X509certSha1Fingerprint); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization saml idp body based on context it is used
func (o *CreateOrganizationSamlIdpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationSamlIdpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationSamlIdpBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationSamlIdpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationSamlIdpCreatedBodyItems0 create organization saml idp created body items0
swagger:model CreateOrganizationSamlIdpCreatedBodyItems0
*/
type CreateOrganizationSamlIdpCreatedBodyItems0 struct {

	// URL that is consuming SAML Identity Provider (IdP)
	ConsumerURL string `json:"consumerUrl,omitempty"`

	// ID associated with the SAML Identity Provider (IdP)
	IdpID string `json:"idpId,omitempty"`

	// Dashboard will redirect users to this URL when they sign out.
	SloLogoutURL string `json:"sloLogoutUrl,omitempty"`

	// Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.
	X509certSha1Fingerprint string `json:"x509certSha1Fingerprint,omitempty"`
}

// Validate validates this create organization saml idp created body items0
func (o *CreateOrganizationSamlIdpCreatedBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization saml idp created body items0 based on context it is used
func (o *CreateOrganizationSamlIdpCreatedBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationSamlIdpCreatedBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationSamlIdpCreatedBodyItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationSamlIdpCreatedBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

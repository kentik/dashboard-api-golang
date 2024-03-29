// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrganizationAdminReader is a Reader for the CreateOrganizationAdmin structure.
type CreateOrganizationAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationAdminCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/admins] createOrganizationAdmin", response, response.Code())
	}
}

// NewCreateOrganizationAdminCreated creates a CreateOrganizationAdminCreated with default headers values
func NewCreateOrganizationAdminCreated() *CreateOrganizationAdminCreated {
	return &CreateOrganizationAdminCreated{}
}

/*
CreateOrganizationAdminCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationAdminCreated struct {
	Payload *CreateOrganizationAdminCreatedBody
}

// IsSuccess returns true when this create organization admin created response has a 2xx status code
func (o *CreateOrganizationAdminCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization admin created response has a 3xx status code
func (o *CreateOrganizationAdminCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization admin created response has a 4xx status code
func (o *CreateOrganizationAdminCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization admin created response has a 5xx status code
func (o *CreateOrganizationAdminCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization admin created response a status code equal to that given
func (o *CreateOrganizationAdminCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization admin created response
func (o *CreateOrganizationAdminCreated) Code() int {
	return 201
}

func (o *CreateOrganizationAdminCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/admins][%d] createOrganizationAdminCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationAdminCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/admins][%d] createOrganizationAdminCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationAdminCreated) GetPayload() *CreateOrganizationAdminCreatedBody {
	return o.Payload
}

func (o *CreateOrganizationAdminCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateOrganizationAdminCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationAdminBody create organization admin body
// Example: {"authenticationMethod":"Email","email":"miles@meraki.com","name":"Miles Meraki","networks":[{"access":"full","id":"N_24329156"}],"orgAccess":"none","tags":[{"access":"read-only","tag":"west"}]}
swagger:model CreateOrganizationAdminBody
*/
type CreateOrganizationAdminBody struct {

	// The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of 'Email' or 'Cisco SecureX Sign-On'. The default is Email authentication
	// Enum: [Cisco SecureX Sign-On Email]
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`

	// The email of the dashboard administrator. This attribute can not be updated.
	// Required: true
	Email *string `json:"email"`

	// The name of the dashboard administrator
	// Required: true
	Name *string `json:"name"`

	// The list of networks that the dashboard administrator has privileges on
	Networks []*CreateOrganizationAdminParamsBodyNetworksItems0 `json:"networks"`

	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	// Required: true
	// Enum: [enterprise full none read-only]
	OrgAccess *string `json:"orgAccess"`

	// The list of tags that the dashboard administrator has privileges on
	Tags []*CreateOrganizationAdminParamsBodyTagsItems0 `json:"tags"`
}

// Validate validates this create organization admin body
func (o *CreateOrganizationAdminBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrgAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createOrganizationAdminBodyTypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cisco SecureX Sign-On","Email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminBodyTypeAuthenticationMethodPropEnum = append(createOrganizationAdminBodyTypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminBodyAuthenticationMethodCiscoSecureXSignDashOn captures enum value "Cisco SecureX Sign-On"
	CreateOrganizationAdminBodyAuthenticationMethodCiscoSecureXSignDashOn string = "Cisco SecureX Sign-On"

	// CreateOrganizationAdminBodyAuthenticationMethodEmail captures enum value "Email"
	CreateOrganizationAdminBodyAuthenticationMethodEmail string = "Email"
)

// prop value enum
func (o *CreateOrganizationAdminBody) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminBodyTypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminBody) validateAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthenticationMethodEnum("createOrganizationAdmin"+"."+"authenticationMethod", "body", o.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAdmin"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAdmin"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminBody) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(o.Networks) { // not required
		return nil
	}

	for i := 0; i < len(o.Networks); i++ {
		if swag.IsZero(o.Networks[i]) { // not required
			continue
		}

		if o.Networks[i] != nil {
			if err := o.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdmin" + "." + "networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdmin" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createOrganizationAdminBodyTypeOrgAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enterprise","full","none","read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminBodyTypeOrgAccessPropEnum = append(createOrganizationAdminBodyTypeOrgAccessPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminBodyOrgAccessEnterprise captures enum value "enterprise"
	CreateOrganizationAdminBodyOrgAccessEnterprise string = "enterprise"

	// CreateOrganizationAdminBodyOrgAccessFull captures enum value "full"
	CreateOrganizationAdminBodyOrgAccessFull string = "full"

	// CreateOrganizationAdminBodyOrgAccessNone captures enum value "none"
	CreateOrganizationAdminBodyOrgAccessNone string = "none"

	// CreateOrganizationAdminBodyOrgAccessReadDashOnly captures enum value "read-only"
	CreateOrganizationAdminBodyOrgAccessReadDashOnly string = "read-only"
)

// prop value enum
func (o *CreateOrganizationAdminBody) validateOrgAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminBodyTypeOrgAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminBody) validateOrgAccess(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationAdmin"+"."+"orgAccess", "body", o.OrgAccess); err != nil {
		return err
	}

	// value enum
	if err := o.validateOrgAccessEnum("createOrganizationAdmin"+"."+"orgAccess", "body", *o.OrgAccess); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminBody) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdmin" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdmin" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization admin body based on the context it is used
func (o *CreateOrganizationAdminBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationAdminBody) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Networks); i++ {

		if o.Networks[i] != nil {

			if swag.IsZero(o.Networks[i]) { // not required
				return nil
			}

			if err := o.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdmin" + "." + "networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdmin" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateOrganizationAdminBody) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Tags); i++ {

		if o.Tags[i] != nil {

			if swag.IsZero(o.Tags[i]) { // not required
				return nil
			}

			if err := o.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdmin" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdmin" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationAdminCreatedBody create organization admin created body
swagger:model CreateOrganizationAdminCreatedBody
*/
type CreateOrganizationAdminCreatedBody struct {

	// Status of the admin's account
	// Enum: [locked ok pending unverified]
	AccountStatus string `json:"accountStatus,omitempty"`

	// Admin's authentication method
	// Enum: [Cisco SecureX Sign-On Email]
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`

	// Admin's email address
	Email string `json:"email,omitempty"`

	// Indicates whether the admin has an API key
	HasAPIKey bool `json:"hasApiKey,omitempty"`

	// Admin's ID
	ID string `json:"id,omitempty"`

	// Time when the admin was last active
	// Format: date-time
	LastActive strfmt.DateTime `json:"lastActive,omitempty"`

	// Admin's username
	Name string `json:"name,omitempty"`

	// Admin network access information
	Networks []*CreateOrganizationAdminCreatedBodyNetworksItems0 `json:"networks"`

	// Admin's level of access to the organization
	// Enum: [enterprise full none read-only]
	OrgAccess string `json:"orgAccess,omitempty"`

	// Admin tag information
	Tags []*CreateOrganizationAdminCreatedBodyTagsItems0 `json:"tags"`

	// Indicates whether two-factor authentication is enabled
	TwoFactorAuthEnabled bool `json:"twoFactorAuthEnabled,omitempty"`
}

// Validate validates this create organization admin created body
func (o *CreateOrganizationAdminCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastActive(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrgAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createOrganizationAdminCreatedBodyTypeAccountStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["locked","ok","pending","unverified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminCreatedBodyTypeAccountStatusPropEnum = append(createOrganizationAdminCreatedBodyTypeAccountStatusPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminCreatedBodyAccountStatusLocked captures enum value "locked"
	CreateOrganizationAdminCreatedBodyAccountStatusLocked string = "locked"

	// CreateOrganizationAdminCreatedBodyAccountStatusOk captures enum value "ok"
	CreateOrganizationAdminCreatedBodyAccountStatusOk string = "ok"

	// CreateOrganizationAdminCreatedBodyAccountStatusPending captures enum value "pending"
	CreateOrganizationAdminCreatedBodyAccountStatusPending string = "pending"

	// CreateOrganizationAdminCreatedBodyAccountStatusUnverified captures enum value "unverified"
	CreateOrganizationAdminCreatedBodyAccountStatusUnverified string = "unverified"
)

// prop value enum
func (o *CreateOrganizationAdminCreatedBody) validateAccountStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminCreatedBodyTypeAccountStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateAccountStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.AccountStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccountStatusEnum("createOrganizationAdminCreated"+"."+"accountStatus", "body", o.AccountStatus); err != nil {
		return err
	}

	return nil
}

var createOrganizationAdminCreatedBodyTypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cisco SecureX Sign-On","Email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminCreatedBodyTypeAuthenticationMethodPropEnum = append(createOrganizationAdminCreatedBodyTypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminCreatedBodyAuthenticationMethodCiscoSecureXSignDashOn captures enum value "Cisco SecureX Sign-On"
	CreateOrganizationAdminCreatedBodyAuthenticationMethodCiscoSecureXSignDashOn string = "Cisco SecureX Sign-On"

	// CreateOrganizationAdminCreatedBodyAuthenticationMethodEmail captures enum value "Email"
	CreateOrganizationAdminCreatedBodyAuthenticationMethodEmail string = "Email"
)

// prop value enum
func (o *CreateOrganizationAdminCreatedBody) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminCreatedBodyTypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthenticationMethodEnum("createOrganizationAdminCreated"+"."+"authenticationMethod", "body", o.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateLastActive(formats strfmt.Registry) error {
	if swag.IsZero(o.LastActive) { // not required
		return nil
	}

	if err := validate.FormatOf("createOrganizationAdminCreated"+"."+"lastActive", "body", "date-time", o.LastActive.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(o.Networks) { // not required
		return nil
	}

	for i := 0; i < len(o.Networks); i++ {
		if swag.IsZero(o.Networks[i]) { // not required
			continue
		}

		if o.Networks[i] != nil {
			if err := o.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdminCreated" + "." + "networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdminCreated" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createOrganizationAdminCreatedBodyTypeOrgAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enterprise","full","none","read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminCreatedBodyTypeOrgAccessPropEnum = append(createOrganizationAdminCreatedBodyTypeOrgAccessPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminCreatedBodyOrgAccessEnterprise captures enum value "enterprise"
	CreateOrganizationAdminCreatedBodyOrgAccessEnterprise string = "enterprise"

	// CreateOrganizationAdminCreatedBodyOrgAccessFull captures enum value "full"
	CreateOrganizationAdminCreatedBodyOrgAccessFull string = "full"

	// CreateOrganizationAdminCreatedBodyOrgAccessNone captures enum value "none"
	CreateOrganizationAdminCreatedBodyOrgAccessNone string = "none"

	// CreateOrganizationAdminCreatedBodyOrgAccessReadDashOnly captures enum value "read-only"
	CreateOrganizationAdminCreatedBodyOrgAccessReadDashOnly string = "read-only"
)

// prop value enum
func (o *CreateOrganizationAdminCreatedBody) validateOrgAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminCreatedBodyTypeOrgAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateOrgAccess(formats strfmt.Registry) error {
	if swag.IsZero(o.OrgAccess) { // not required
		return nil
	}

	// value enum
	if err := o.validateOrgAccessEnum("createOrganizationAdminCreated"+"."+"orgAccess", "body", o.OrgAccess); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminCreatedBody) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdminCreated" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdminCreated" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization admin created body based on the context it is used
func (o *CreateOrganizationAdminCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationAdminCreatedBody) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Networks); i++ {

		if o.Networks[i] != nil {

			if swag.IsZero(o.Networks[i]) { // not required
				return nil
			}

			if err := o.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdminCreated" + "." + "networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdminCreated" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateOrganizationAdminCreatedBody) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Tags); i++ {

		if o.Tags[i] != nil {

			if swag.IsZero(o.Tags[i]) { // not required
				return nil
			}

			if err := o.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationAdminCreated" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationAdminCreated" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationAdminCreatedBodyNetworksItems0 create organization admin created body networks items0
swagger:model CreateOrganizationAdminCreatedBodyNetworksItems0
*/
type CreateOrganizationAdminCreatedBodyNetworksItems0 struct {

	// Admin's level of access to the network
	Access string `json:"access,omitempty"`

	// Network ID
	ID string `json:"id,omitempty"`
}

// Validate validates this create organization admin created body networks items0
func (o *CreateOrganizationAdminCreatedBodyNetworksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization admin created body networks items0 based on context it is used
func (o *CreateOrganizationAdminCreatedBodyNetworksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBodyNetworksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBodyNetworksItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminCreatedBodyNetworksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationAdminCreatedBodyTagsItems0 create organization admin created body tags items0
swagger:model CreateOrganizationAdminCreatedBodyTagsItems0
*/
type CreateOrganizationAdminCreatedBodyTagsItems0 struct {

	// Access level for the tag
	Access string `json:"access,omitempty"`

	// Tag value
	Tag string `json:"tag,omitempty"`
}

// Validate validates this create organization admin created body tags items0
func (o *CreateOrganizationAdminCreatedBodyTagsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization admin created body tags items0 based on context it is used
func (o *CreateOrganizationAdminCreatedBodyTagsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBodyTagsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminCreatedBodyTagsItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminCreatedBodyTagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationAdminParamsBodyNetworksItems0 create organization admin params body networks items0
swagger:model CreateOrganizationAdminParamsBodyNetworksItems0
*/
type CreateOrganizationAdminParamsBodyNetworksItems0 struct {

	// The privilege of the dashboard administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	// Required: true
	Access *string `json:"access"`

	// The network ID
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this create organization admin params body networks items0
func (o *CreateOrganizationAdminParamsBodyNetworksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationAdminParamsBodyNetworksItems0) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", o.Access); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminParamsBodyNetworksItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization admin params body networks items0 based on context it is used
func (o *CreateOrganizationAdminParamsBodyNetworksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminParamsBodyNetworksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminParamsBodyNetworksItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminParamsBodyNetworksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationAdminParamsBodyTagsItems0 create organization admin params body tags items0
swagger:model CreateOrganizationAdminParamsBodyTagsItems0
*/
type CreateOrganizationAdminParamsBodyTagsItems0 struct {

	// The privilege of the dashboard administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	// Required: true
	// Enum: [full guest-ambassador monitor-only read-only]
	Access *string `json:"access"`

	// The name of the tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this create organization admin params body tags items0
func (o *CreateOrganizationAdminParamsBodyTagsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createOrganizationAdminParamsBodyTagsItems0TypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","guest-ambassador","monitor-only","read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationAdminParamsBodyTagsItems0TypeAccessPropEnum = append(createOrganizationAdminParamsBodyTagsItems0TypeAccessPropEnum, v)
	}
}

const (

	// CreateOrganizationAdminParamsBodyTagsItems0AccessFull captures enum value "full"
	CreateOrganizationAdminParamsBodyTagsItems0AccessFull string = "full"

	// CreateOrganizationAdminParamsBodyTagsItems0AccessGuestDashAmbassador captures enum value "guest-ambassador"
	CreateOrganizationAdminParamsBodyTagsItems0AccessGuestDashAmbassador string = "guest-ambassador"

	// CreateOrganizationAdminParamsBodyTagsItems0AccessMonitorDashOnly captures enum value "monitor-only"
	CreateOrganizationAdminParamsBodyTagsItems0AccessMonitorDashOnly string = "monitor-only"

	// CreateOrganizationAdminParamsBodyTagsItems0AccessReadDashOnly captures enum value "read-only"
	CreateOrganizationAdminParamsBodyTagsItems0AccessReadDashOnly string = "read-only"
)

// prop value enum
func (o *CreateOrganizationAdminParamsBodyTagsItems0) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationAdminParamsBodyTagsItems0TypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationAdminParamsBodyTagsItems0) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", o.Access); err != nil {
		return err
	}

	// value enum
	if err := o.validateAccessEnum("access", "body", *o.Access); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationAdminParamsBodyTagsItems0) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", o.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization admin params body tags items0 based on context it is used
func (o *CreateOrganizationAdminParamsBodyTagsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationAdminParamsBodyTagsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationAdminParamsBodyTagsItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationAdminParamsBodyTagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
